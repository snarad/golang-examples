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
	fmt.Printf("Sum function was invoked with numbers: %v \n", req)
	num1 := req.GetFirstNumber()
	num2 := req.GetSecondNumber()
	result := num1 + num2
	res := &calculatorpb.SumResponse{
		SumResult: result,
	}
	return res, nil
}

func (*server) PrimeDecomposition(req *calculatorpb.PrimeDecompositionRequest, stream calculatorpb.CalculatorService_PrimeDecompositionServer) error {
	fmt.Printf("PrimeDecomposition function was invoked with req: %v \n", req)
	divisor := int64(2)
	number := req.GetNumber()
	for number > 1 {
		if number%divisor == 0 {
			primeFactor := divisor
			stream.Send(&calculatorpb.PrimeDecompositionResponse{
				PrimeFactor: primeFactor,
			})
			number = number / divisor
		} else {
			divisor = divisor + 1
			fmt.Printf("Divisior has increased to %v \n", divisor)
		}
	}
	return nil
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
