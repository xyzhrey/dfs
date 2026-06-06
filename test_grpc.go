package main

import (
	"context"
	"fmt"

	dfsgrpc "dfs/internal/grpc"

	"dfs/storagepb"
)

func main() {

	client, conn, err := dfsgrpc.NewStorageClient(
		"localhost:50051",
	)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	resp, err := client.StoreChunk(
		context.Background(),
		&storagepb.StoreChunkRequest{
			ChunkId: "chunk1",
			Data:    []byte("hello dfs"),
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Path)
}
