package main

import (
	pb "github.com/tjoe1985/grcp_go_course/greet/greetpb"
	"google.golang.org/grpc"

	"log"
	"net"
)

const (
	port = "0.0.0.0:50051"
)

type GreetServiceServer struct {
	pb.UnimplementedGreetServiceServer
}

func main() {
	log.Println("Hello father Joel i am running. ")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("error failed to listen: ", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &GreetServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Println("Error Serving: ", err)
	}
}
