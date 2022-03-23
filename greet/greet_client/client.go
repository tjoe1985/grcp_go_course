package main

import (
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
	log.Println("created client", client)
}
