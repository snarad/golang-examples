package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/snarad/golang-examples/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Add function was invoked with numbers: %v \n", req)
	num1 := req.GetFirstNumber()
	num2 := req.GetSecondNumber()
	result := num1 + num2
	res := &calculatorpb.SumResponse{
		SumResult: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Calculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
