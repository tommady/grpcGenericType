package main

import (
	"log"

	"github.com/golang/protobuf/ptypes/struct"
	pb "github.com/tommady/grpcGenericType/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGenericDoClient(conn)
	req := new(pb.Request)
	req.Arg = &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: "string"}}
	r, err := c.DoSomething(context.Background(), req)
	if err != nil {
		log.Fatalf("str generic failed: %v", err)
	}
	log.Printf("Str: %s", r.Ret.GetStringValue())

	req.Arg = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: 3.14159}}
	r, err = c.DoSomething(context.Background(), req)
	if err != nil {
		log.Fatalf("number generic failed: %v", err)
	}
	log.Printf("Number: %f", r.Ret.GetNumberValue())
}
