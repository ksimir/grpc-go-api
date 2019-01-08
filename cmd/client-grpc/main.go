package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	v1 "github.com/ksimir/grpc-go-api/pkg/api/v1"
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
}

// createPlayer calls the RPC method CreatePlayer of PlayerServer
func createPlayer(client v1.PlayerClient, player *v1.PlayerRequest) {
	resp, err := client.CreatePlayer(context.Background(), player)
	if err != nil {
		log.Fatalf("Could not create player: %v", err)
	}
	if resp.Success {
		log.Printf("A new player has been added with id: %s", resp.Id)
	} else {
		log.Printf("Failed to add the new player with id: %s", resp.Id)
	}
}

// getCustomers calls the RPC method GetCustomers of CustomerServer
func getPlayers(client v1.PlayerClient, filter *v1.PlayerFilter) {
	// calling the streaming API
	stream, err := client.GetPlayers(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v", err)
	}
	for {
		player, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetPlayers(_) = _, %v", client, err)
		}
		log.Printf("Player: %v", player)
	}
}

// getCustomer calls the RPC method GetCustomers of CustomerServer
func getPlayer(client v1.PlayerClient, filter *v1.PlayerId) {
	player, err := client.GetPlayer(context.Background(), filter)
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
	flag.Parse()

	address := cfg.GRPCAddress + ":" + cfg.GRPCPort

	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
	client := v1.NewPlayerClient(conn)

	u1, err := uuid.NewRandom()
	fmt.Printf("Player #1 ID is %s", u1.String())
	player := &v1.PlayerRequest{
		Api:      apiVersion,
		Id:       u1.String(),
		Username: "Samir Hammoudi",
		Email:    "sh@xyz.com",
		Phone:    "732-757-2923",
	}

	// Create a new player
	createPlayer(client, player)

	u2, err := uuid.NewRandom()
	fmt.Printf("Player #2 ID is : %s", u2.String())
	player = &v1.PlayerRequest{
		Api:      apiVersion,
		Id:       u2.String(),
		Username: "Zinedine Zidane",
		Email:    "zz@xyz.com",
		Phone:    "732-757-2924",
	}

	// Create a new player
	createPlayer(client, player)

	// Filter with an empty Keyword
	filter := &v1.PlayerFilter{
		Api:     apiVersion,
		Keyword: "Samir Hammoudi",
	}
	getPlayers(client, filter)

	// Get player with a specific ID
	id := &v1.PlayerId{
		Api: apiVersion,
		Id:  "fb19b3e3-ee5f-4d68-b554-663b60032d3f",
	}
	getPlayer(client, id)
}
