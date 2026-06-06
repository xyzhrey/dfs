package grpc

import (
	"dfs/storagepb"

	gogrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewStorageClient(
	address string,
) (
	storagepb.StorageServiceClient,
	*gogrpc.ClientConn,
	error,
) {

	conn, err := gogrpc.NewClient(
		address,
		gogrpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		return nil, nil, err
	}

	client := storagepb.NewStorageServiceClient(conn)

	return client, conn, nil
}
