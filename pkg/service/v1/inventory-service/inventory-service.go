package v1

import (
	"context"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/spanner"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/ksimir/grpc-go-api/pkg/api/v1/inventory"
	"github.com/ksimir/grpc-go-api/pkg/logger"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// server is used to implement inventory.InventoryServer.
type server struct {
	dataClient *spanner.Client
}

// NewInventoryServiceServer creates Inventory service
func NewInventoryServiceServer(db *spanner.Client) v1.InventoryServer {
	return &server{dataClient: db}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *server) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// AddItem add a new Item
func (s *server) AddItem(ctx context.Context, in *v1.ItemRequest) (*v1.ItemResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(in.Api); err != nil {
		return nil, err
	}

	start := time.Now()

	if err := saveItem(ctx, os.Stdout, in, s.dataClient); err != nil {
		logger.Log.Error("Failed to save item to Spanner", zap.Int64("Item ID", in.Id))
		return &v1.ItemResponse{Api: apiVersion, Success: false}, err
	}

	elapsed := time.Since(start)
	logger.Log.Debug("saveItem duration", zap.Duration("Duration", elapsed))

	return &v1.ItemResponse{Api: apiVersion, Success: true}, nil
}

// GetItems returns all items from a given player
func (s *server) GetItems(in *v1.ItemRequest, stream v1.Inventory_GetItemsServer) error {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(in.Api); err != nil {
		return err
	}

	items, err := readItemsByPid(context.Background(), os.Stdout, in, s.dataClient)
	if err != nil {
		logger.Log.Error("Failed to read items information from Spanner")
		return err
	}

	for _, item := range items {
		if in.Pid != "" {
			if !strings.Contains(item.Pid, in.Pid) {
				continue
			}
		}
		if err := stream.Send(item); err != nil {
			return err
		}
	}
	return nil
}

// GetItem returns an item by given item id and player id
func (s *server) GetItem(ctx context.Context, in *v1.ItemRequest) (*v1.ItemRequest, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(in.Api); err != nil {
		return nil, err
	}

	// define the default empty response
	item := v1.ItemRequest{
		Api: apiVersion,
	}

	items, err := readItemByID(ctx, os.Stdout, in, s.dataClient)
	if err != nil {
		logger.Log.Error("Failed to read item information from Spanner")
		return nil, err
	}

	// only 1 result should be returneｄ because the combination of player ID and item ID is unique
	if items[0].Id == in.Id {
		item := v1.ItemRequest{
			Api:      apiVersion,
			Id:       items[0].Id,
			Pid:      items[0].Pid,
			Quantity: items[0].Quantity,
		}
		return &item, nil
	}
	return &item, nil
}

// UpdateItem an existing Item
func (s *server) UpdateItem(ctx context.Context, in *v1.ItemRequest) (*v1.ItemResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(in.Api); err != nil {
		return nil, err
	}

	//TODO

	return &v1.ItemResponse{Api: apiVersion, Success: true}, nil
}

// DeleteItem an existing Item
func (s *server) DeleteItem(ctx context.Context, in *v1.ItemRequest) (*v1.ItemResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(in.Api); err != nil {
		return nil, err
	}

	//TODO

	return &v1.ItemResponse{Api: apiVersion, Success: true}, nil
}

// saveItem saves an item to Cloud Spanner
func saveItem(ctx context.Context, w io.Writer, in *v1.ItemRequest, dataClient *spanner.Client) error {
	itemColumns := []string{"id", "item_id", "quantity"}
	m := []*spanner.Mutation{
		spanner.InsertOrUpdate("item_inventory", itemColumns, []interface{}{in.Pid, in.Id, in.Quantity}),
	}
	_, err := dataClient.Apply(ctx, m)
	return err
}

// readItemsByPid returns the item inventory for a specific player
func readItemsByPid(ctx context.Context, w io.Writer, in *v1.ItemRequest, dataClient *spanner.Client) ([]*v1.ItemRequest, error) {
	var result []*v1.ItemRequest
	start := time.Now()
	ro := dataClient.ReadOnlyTransaction()
	defer ro.Close()

	sqlStatement := "SELECT id, item_id, quantity FROM item_inventory WHERE id = '" + in.Pid + "'"
	stmt := spanner.Statement{SQL: sqlStatement}
	iter := ro.Query(ctx, stmt)
	defer iter.Stop()

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var id string
		var itemID, quantity int64
		if err := row.Columns(&id, &itemID, &quantity); err != nil {
			return nil, err
		}

		item := v1.ItemRequest{
			Api:      apiVersion,
			Id:       itemID,
			Pid:      id,
			Quantity: quantity,
		}
		result = append(result, &item)
	}
	elapsed := time.Since(start)
	logger.Log.Debug("readItemsByPid duration", zap.Duration("Duration", elapsed))
	return result, nil
}

// readItemByID returns the item with id specified in the request
func readItemByID(ctx context.Context, w io.Writer, in *v1.ItemRequest, dataClient *spanner.Client) ([]*v1.ItemRequest, error) {
	var result []*v1.ItemRequest
	start := time.Now()
	ro := dataClient.ReadOnlyTransaction()
	defer ro.Close()

	sqlStatement := "SELECT id, item_id, quantity FROM item_inventory WHERE id = '" + in.Pid + "' AND item_id = " + strconv.FormatInt(in.Id, 10)
	stmt := spanner.Statement{SQL: sqlStatement}
	iter := ro.Query(ctx, stmt)
	defer iter.Stop()

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var id string
		var itemID, quantity int64
		if err := row.Columns(&id, &itemID, &quantity); err != nil {
			return nil, err
		}

		item := v1.ItemRequest{
			Api:      apiVersion,
			Id:       itemID,
			Pid:      id,
			Quantity: quantity,
		}
		result = append(result, &item)
	}
	elapsed := time.Since(start)
	logger.Log.Debug("readItemByID duration", zap.Duration("Duration", elapsed))
	return result, nil
}
