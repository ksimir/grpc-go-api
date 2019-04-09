package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	ic "github.com/ksimir/grpc-go-api/pkg/api/v1/inventory"
	pc "github.com/ksimir/grpc-go-api/pkg/api/v1/player"
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

// createPlayer calls the RPC method CreatePlayer of PlayerServer
func createPlayer(ctx context.Context, client pc.PlayerClient, player *pc.PlayerRequest) {
	resp, err := client.CreatePlayer(ctx, player)
	if err != nil {
		log.Fatalf("Could not create player: %v", err)
	}
	if resp.Success {
		log.Printf("A new player has been added with id: %s", resp.Id)
	} else {
		log.Printf("Failed to add the new player with id: %s", resp.Id)
	}
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

// getCustomer calls the RPC method GetCustomers of CustomerServer
func getPlayer(ctx context.Context, client pc.PlayerClient, filter *pc.PlayerId) {
	player, err := client.GetPlayer(ctx, filter)
	if err != nil {
		log.Fatalf("Could not get Player: %v", err)
	}
	if player != nil {
		if player.Id == "" {
			log.Printf("No player found with ID %s", filter.Id)
		} else {
			log.Printf("Player: %v", player)
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

	// Creates a new client for each gRPC API
	playclient := pc.NewPlayerClient(conn)
	invClient := ic.NewInventoryClient(conn)

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

	u1, err := uuid.NewRandom()
	fmt.Printf("Player #1 ID is %s", u1.String())
	player := &pc.PlayerRequest{
		Api:      apiVersion,
		Id:       u1.String(),
		Username: "Samir Hammoudi",
		Email:    "sh@xyz.com",
		Phone:    "732-757-2923",
	}

	// Create a new player
	createPlayer(ctx, playclient, player)

	u2, err := uuid.NewRandom()
	fmt.Printf("Player #2 ID is : %s", u2.String())
	player = &pc.PlayerRequest{
		Api:      apiVersion,
		Id:       u2.String(),
		Username: "Zinedine Zidane",
		Email:    "zz@xyz.com",
		Phone:    "732-757-2924",
	}

	// Create a new player
	createPlayer(ctx, playclient, player)

	// Get player with a specific ID
	id := &pc.PlayerId{
		Api: apiVersion,
		Id:  "fb19b3e3-ee5f-4d68-b554-663b60032d3f",
	}
	getPlayer(ctx, playclient, id)

	fmt.Printf("Adding Item 1 to Player fb19b3e3-ee5f-4d68-b554-663b60032d3f")
	item := &ic.ItemRequest{
		Api:      apiVersion,
		Id:       2,
		Pid:      "fb19b3e3-ee5f-4d68-b554-663b60032d3f",
		Quantity: 1,
	}

	// Create a new player
	addItem(ctx, invClient, item)
}
