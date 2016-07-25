package main

import (
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/struct"
	pb "github.com/tommady/grpcGenericType/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) DoSomething(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	ret := new(pb.Reply)
	if x, ok := in.Arg.GetKind().(*structpb.Value_StringValue); ok {
		s := x.StringValue + " Generic!!!"
		ret.Ret = &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: s}}
	} else if x, ok := in.Arg.GetKind().(*structpb.Value_NumberValue); ok {
		i := x.NumberValue + 3345678
		ret.Ret = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: i}}
	} else {
		return nil, fmt.Errorf("Type is not allowed")
	}
	return ret, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGenericDoServer(s, &server{})
	s.Serve(lis)
}
