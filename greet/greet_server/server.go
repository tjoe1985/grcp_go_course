package main

import (
	"context"
	pb "github.com/tjoe1985/grcp_go_course/greet/greetpb"
	"google.golang.org/grpc"
	"strconv"
	"time"

	"log"
	"net"
)

const (
	port = "0.0.0.0:50051"
)

type GreetServiceServer struct {
	pb.UnimplementedGreetServiceServer
}
type server struct {
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
func (*GreetServiceServer) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("Greet function was invoked with: ", req)
	firstName := req.GetGreeting().GetFirstName()
	preparedString := "Hello there ...." + firstName + " is your name? I think?"
	return &pb.GreetResponse{Result: preparedString}, nil
}
func (*server) GreetManyTimes(req *pb.GreetManytimesRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Println("GreetManyTimes function was invoked with: ", req)
	fName := req.GetGreeting().GetFirstName()
	for i := 1; i < 10; i++ {
		result := "This is a streaming Request " + fName + " and you will get it every " +
			"5 seconds cause i refuse to be fast so here is number: " + strconv.Itoa(i)
		time.Sleep(5 * time.Second)
		res := &pb.GreetManytimesResponse{Result: result}
		stream.Send(res)
	}
	return nil
}
