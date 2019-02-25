package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"github.com/snarad/golang-examples/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("ComputeAverage function was invoked with req\n")
	i := float64(0)
	sum := float64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			computedAverage := sum / i
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				ComputedAvergae: computedAverage,
			})
		}
		if err != nil {
			log.Fatalf("error while recieving client stream: %v", err)
		}
		sum += req.GetNumber()
		i++
	}
}

func (*server) FindMax(stream calculatorpb.CalculatorService_FindMaxServer) error {
	fmt.Printf("FindMax function was invoked\n")
	maxNumber := int64(math.MinInt64)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while recieving client request stream: %v", err)
		}
		fmt.Printf("Recieved new number: %v\n", req.GetNumber())
		number := req.GetNumber()
		if number > maxNumber {
			maxNumber = number
			if err = stream.Send(&calculatorpb.FindMaxResponse{
				MaxNumber: int64(maxNumber),
			}); err != nil {
				log.Fatalf("Error in sending the stream to client: %v", err)
			}
		}
	}
}

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	fmt.Printf("SquareRoot function was invoked with req: %v\n", req)
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Recieved a negative number: %v", number))
	}
	return &calculatorpb.SquareRootResponse{SquareRoot: math.Sqrt(float64(number))}, nil
}

func main() {
	fmt.Println("Calculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	reflection.Register(s)

	s.Serve(lis)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
