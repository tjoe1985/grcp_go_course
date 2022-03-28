package main

import (
	"context"
	cpb "github.com/tjoe1985/grcp_go_course/calculator/calpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const address = "localhost:50051"

func main() {
	log.Println("cal client is running")
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Println("error connecting : ", err)
	}
	defer conn.Close()

	client := cpb.NewSumServiceClient(conn)
	doCalUnary(client)
}

func doCalUnary(client cpb.SumServiceClient) {
	log.Println("Running Unary RPC request from cal client service")
	// create request
	req := &cpb.SumRequest{
		Params: &cpb.Params{
			A: 10,
			B: 3,
		},
	}
	response, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Println("Error during Greet request: ", err)
	}
	log.Println("Incoming response:", response.Result)
}
