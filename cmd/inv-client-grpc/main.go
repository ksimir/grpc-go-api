package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	ic "github.com/ksimir/grpc-go-api/pkg/api/v1/inventory"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// Config is configuration for remote server
type Config struct {
	// gRPC server address/port connection parameters section
	GRPCAddress string
	GRPCPort    string
	// gRPC authentication parameters section
	key       string
	token     string
	keyfile   string
	audience  string // name of the service configuration in api_auth.yaml
	verifytls bool
}

// addItem calls the RPC method AddItem of InventoryServer
func addItem(ctx context.Context, client ic.InventoryClient, item *ic.ItemRequest) {
	resp, err := client.AddItem(ctx, item)
	if err != nil {
		log.Fatalf("Could not add item: %v", err)
	}
	if resp.Success {
		log.Printf("A new item has been added")
	} else {
		log.Printf("Failed to add the new item")
	}
}

// getCustomers calls the RPC method GetCustomers of CustomerServer
func getItems(ctx context.Context, client ic.InventoryClient, item *ic.ItemRequest) {
	// calling the streaming API
	stream, err := client.GetItems(ctx, item)
	if err != nil {
		log.Fatalf("Error on get items: %v", err)
	}
	for {
		item, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetItems(_) = _, %v", client, err)
		}
		log.Printf("Item: %v", item)
	}
}

// getCustomer calls the RPC method GetCustomers of CustomerServer
func getItem(ctx context.Context, client ic.InventoryClient, item *ic.ItemRequest) {
	item, err := client.GetItem(ctx, item)
	if err != nil {
		log.Fatalf("Could not get Item: %v", err)
	}
	if item != nil {
		if item.Pid == "" {
			log.Printf("No item found for player ID %s", item.Pid)
		} else {
			log.Printf("Item: %v", item)
		}
	}
}

func main() {
	var cfg Config
	flag.StringVar(&cfg.GRPCAddress, "grpc-address", "", "gRPC address to connect to")
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to connect to")
	flag.StringVar(&cfg.key, "api-key", "", "API key.")
	flag.StringVar(&cfg.token, "token", "", "Authentication token.")
	flag.StringVar(&cfg.keyfile, "keyfile", "", "Path to a Google service account key file.")
	flag.StringVar(&cfg.audience, "audience", "", "Audience.")
	flag.BoolVar(&cfg.verifytls, "verifytls", true, "FLag to disable cert authentication.")
	flag.Parse()

	address := cfg.GRPCAddress + ":" + cfg.GRPCPort
	var conn *grpc.ClientConn
	var err error

	if cfg.verifytls {
		creds, err := credentials.NewClientTLSFromFile("root.pem", "")
		if err != nil {
			log.Fatalf("failed to generate credentials: %v", err)
		}

		// Set up a connection to the gRPC server.
		conn, err = grpc.Dial(address, grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
	} else {
		conn, err = grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
	}

	// Creates a new CustomerClient
	client := ic.NewInventoryClient(conn)

	// API authentication section
	if cfg.keyfile != "" {
		log.Printf("Authenticating using Google service account key in %s", cfg.keyfile)
		keyBytes, err := ioutil.ReadFile(cfg.keyfile)
		if err != nil {
			log.Fatalf("Unable to read service account key file %s: %v", cfg.keyfile, err)
		}

		tokenSource, err := google.JWTAccessTokenSourceFromJSON(keyBytes, cfg.audience)
		if err != nil {
			log.Fatalf("Error building JWT access token source: %v", err)
		}
		jwt, err := tokenSource.Token()
		if err != nil {
			log.Fatalf("Unable to generate JWT token: %v", err)
		}
		cfg.token = jwt.AccessToken
		// NOTE: the generated JWT token has a 1h TTL.
		// Make sure to refresh the token before it expires by calling TokenSource.Token() for each outgoing requests.
		// Calls to this particular implementation of TokenSource.Token() are cheap.
	}

	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	// defer cancel()

	// If API key is provided, then add x-api-key in the metadata
	if cfg.key != "" {
		log.Printf("Using API key: %s", cfg.key)
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("x-api-key", cfg.key))
	}
	// If service account is provided with a JSON key file
	if cfg.token != "" {
		log.Printf("Using authentication token: %s", cfg.token)
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("Authorization", fmt.Sprintf("Bearer %s", cfg.token)))
	}

	fmt.Printf("Adding Item 1 to Player fb19b3e3-ee5f-4d68-b554-663b60032d3f")
	item := &ic.ItemRequest{
		Api:      apiVersion,
		Id:       1,
		Pid:      "fb19b3e3-ee5f-4d68-b554-663b60032d3f",
		Quantity: 1,
	}
	// Add a new item
	addItem(ctx, client, item)

	// Filter with an empty Keyword
	filter := &ic.ItemRequest{
		Api: apiVersion,
		Pid: "fb19b3e3-ee5f-4d68-b554-663b60032d3f",
	}
	getItems(ctx, client, filter)

	// Get item with a specific ID
	filter = &ic.ItemRequest{
		Api: apiVersion,
		Id:  1,
		Pid: "fb19b3e3-ee5f-4d68-b554-663b60032d3f",
	}
	getItem(ctx, client, filter)
}
