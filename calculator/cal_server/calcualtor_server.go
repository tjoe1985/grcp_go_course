package cal_server

import (
	"context"
	cpb "github.com/tjoe1985/grcp_go_course/calculator/calpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = "0.0.0.0:50051"

type CalServiceServer struct {
	cpb.UnimplementedSumServiceServer
}

func main() {
	log.Println("Hello main from calculator services is running. ")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("error failed to listen: ", err)
	}
	s := grpc.NewServer()
	cpb.RegisterSumServiceServer(s, &CalServiceServer{})
	err = s.Serve(lis)
	if err != nil {
		log.Println("Error Serving: ", err)
	}
}

func (*CalServiceServer) Sum(ctx context.Context, req *cpb.SumRequest) (*cpb.SumResponse, error) {
	log.Println("Sum functions was invoked with: ", req)
	a := req.GetParams().GetA()
	b := req.GetParams().GetB()
	c := a + b
	return &cpb.SumResponse{Result: c}, nil
}
