package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/snarad/golang-examples/grpc-go-course/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error in connecting to server: % v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("erro while calling Sum RPC %v ", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")
	req := &calculatorpb.PrimeDecompositionRequest{
		Number: 210,
	}

	res, err := c.PrimeDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeDecomposition RPC %v ", err)
	}

	for {
		decomposition, err := res.Recv()
		if err == io.EOF {
			// We have reached end of stream
			break
		}
		if err != nil {
			log.Fatalf("Something happened %v", err)
		}
		log.Printf("Response from PrimeDecomposition: %v \n", decomposition.GetPrimeFactor())
	}
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")
	stream, err := c.ComputeAverage(context.Background())
	numbers := []float64{1.0, 2.0, 3.0, 4.0}
	for _, number := range numbers {
		fmt.Printf("Sending numbers: %v\n", number)
		if err := stream.Send(&calculatorpb.ComputeAvergaeRequest{
			Number: number,
		}); err != nil {
			log.Fatalf("error in sending request: %v", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while ComputeAverage RPC")
	}
	log.Printf("Compute Average response: %v", res.GetComputedAvergae())
}
