#!/bin/bash

# run the below commands at grpc-go-course level
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.