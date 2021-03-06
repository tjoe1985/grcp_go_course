package main

import (
	"context"
	pb "github.com/tjoe1985/grcp_go_course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const address = "localhost:50051"

func main() {
	log.Println("hello i am running")
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Println("error connecting : ", err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)
	doUnary(client)
}

func doUnary(client pb.GreetServiceClient) {
	log.Println("Running Unary RPC Request.")
	//Created request
	req := &pb.GreetRequest{Greeting: &pb.Greeting{FirstName: "Joel", LastName: "Ewy"}}
	response, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Println("Error during Greet request: ", err)
	}
	log.Println("Incoming response:", response.Result)
}
