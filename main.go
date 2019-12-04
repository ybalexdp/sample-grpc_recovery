package main

import (
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/ybalexdp/sample-grpc_recovery/pb"
	"github.com/ybalexdp/sample-grpc_recovery/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func main() {

	//var recoveryFunc grpc_recovery.RecoveryHandlerFunc
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoveryFunc),
	}

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
	)

	listenPort, err := net.Listen("tcp", ":2007")
	if err != nil {
		log.Fatalln(err)
	}
	sampleService := &service.SampleService{}
	pb.RegisterSampleServer(server, sampleService)
	server.Serve(listenPort)
}

func recoveryFunc(p interface{}) error {
	fmt.Printf("p: %+v\n", p)
	return grpc.Errorf(codes.Internal, "Unexpected error")
}
