package v1

import (
	"context"
	"io"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/spanner"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/ksimir/grpc-go-api/pkg/api/v1"
	"github.com/ksimir/grpc-go-api/pkg/logger"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// server is used to implement player.PlayerServer.
type server struct {
	dataClient *spanner.Client
}

// NewPlayerServiceServer creates Player service
func NewPlayerServiceServer(db *spanner.Client) v1.PlayerServer {
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

// CreatePlayer creates a new Player
func (s *server) CreatePlayer(ctx context.Context, in *v1.PlayerRequest) (*v1.PlayerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(in.Api); err != nil {
		return nil, err
	}

	start := time.Now()

	if err := savePlayer(ctx, os.Stdout, in, s.dataClient); err != nil {
		logger.Log.Fatal("Failed to save player to Spanner", zap.String("Player ID", in.Id))
		return &v1.PlayerResponse{Api: apiVersion, Id: in.Id, Success: false}, err
	}

	elapsed := time.Since(start)
	logger.Log.Debug("savePlayer duration", zap.Duration("Duration", elapsed))

	return &v1.PlayerResponse{Api: apiVersion, Id: in.Id, Success: true}, nil
}

// GetPlayers returns all players by given filter
func (s *server) GetPlayers(filter *v1.PlayerFilter, stream v1.Player_GetPlayersServer) error {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(filter.Api); err != nil {
		return err
	}

	players, err := readPlayersByUsername(context.Background(), os.Stdout, filter, s.dataClient)
	if err != nil {
		logger.Log.Fatal("Failed to read players information from Spanner")
		return err
	}

	for _, player := range players {
		if filter.Keyword != "" {
			if !strings.Contains(player.Username, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(player); err != nil {
			return err
		}
	}
	return nil
}

// GetPlayer returns player by given id
func (s *server) GetPlayer(ctx context.Context, filter *v1.PlayerId) (*v1.PlayerRequest, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(filter.Api); err != nil {
		return nil, err
	}

	// define the default empty response
	player := v1.PlayerRequest{
		Api: apiVersion,
	}

	players, err := readPlayerByID(ctx, os.Stdout, filter, s.dataClient)
	if err != nil {
		logger.Log.Fatal("Failed to read player information from Spanner")
		return nil, err
	}

	// only 1 result should be returneｄ because Player.ID is unique
	if filter.Id != "" {
		if strings.Contains(players[0].Id, filter.Id) {
			player := v1.PlayerRequest{
				Api:      apiVersion,
				Id:       players[0].Id,
				Username: players[0].Username,
				Email:    players[0].Email,
				Phone:    players[0].Phone,
			}
			return &player, nil
		}
	}
	return &player, nil
}

// savePlayer saves a player to Cloud Spanner
func savePlayer(ctx context.Context, w io.Writer, in *v1.PlayerRequest, dataClient *spanner.Client) error {
	playerColumns := []string{"id", "username", "email", "phone"}
	m := []*spanner.Mutation{
		spanner.InsertOrUpdate("Players", playerColumns, []interface{}{in.Id, in.Username, in.Email, in.Phone}),
	}
	_, err := dataClient.Apply(ctx, m)
	return err
}

// readPlayersByUsername returns the player list with username specified in the filter
func readPlayersByUsername(ctx context.Context, w io.Writer, filter *v1.PlayerFilter, dataClient *spanner.Client) ([]*v1.PlayerRequest, error) {
	var result []*v1.PlayerRequest
	start := time.Now()
	ro := dataClient.ReadOnlyTransaction()
	defer ro.Close()

	sqlStatement := "SELECT id, username, email, phone FROM Players WHERE username = '" + filter.Keyword + "'"
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
		var id, username, email, phone string
		if err := row.Columns(&id, &username, &email, &phone); err != nil {
			return nil, err
		}

		player := v1.PlayerRequest{
			Api:      apiVersion,
			Id:       id,
			Username: username,
			Email:    email,
			Phone:    phone,
		}
		result = append(result, &player)
	}
	elapsed := time.Since(start)
	logger.Log.Debug("readPlayersByUsername duration", zap.Duration("Duration", elapsed))
	return result, nil
}

// readPlayerByID returns the player with id specified in the request
func readPlayerByID(ctx context.Context, w io.Writer, filter *v1.PlayerId, dataClient *spanner.Client) ([]*v1.PlayerRequest, error) {
	var result []*v1.PlayerRequest
	start := time.Now()
	ro := dataClient.ReadOnlyTransaction()
	defer ro.Close()

	sqlStatement := "SELECT id, username, email, phone FROM Players WHERE id = '" + filter.Id + "'"
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

		var id, username, email, phone string
		if err := row.Columns(&id, &username, &email, &phone); err != nil {
			return nil, err
		}

		player := v1.PlayerRequest{
			Api:      apiVersion,
			Id:       id,
			Username: username,
			Email:    email,
			Phone:    phone,
		}
		result = append(result, &player)
	}
	elapsed := time.Since(start)
	logger.Log.Debug("readPlayerByID duration", zap.Duration("Duration", elapsed))
	return result, nil
}
