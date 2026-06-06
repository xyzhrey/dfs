package main

import (
	"log"
	"net"

	dfsgrpc "dfs/internal/grpc"
	"dfs/storagepb"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen(
		"tcp",
		":50051",
	)

	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	storagepb.RegisterStorageServiceServer(
		server,
		dfsgrpc.NewStorageServer(),
	)

	log.Println(
		"storage node listening on :50051",
	)

	log.Fatal(
		server.Serve(lis),
	)
}
