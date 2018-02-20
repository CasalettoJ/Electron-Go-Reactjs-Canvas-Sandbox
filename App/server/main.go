package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/casalettoj/exploration/App/proto"
	"github.com/casalettoj/exploration/App/server/pkg"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Opening server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50050))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := hello.Server{}
	grpcServer := grpc.NewServer()

	pb.RegisterTalkServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
