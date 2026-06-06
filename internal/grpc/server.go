package grpc

import (
	"context"

	"dfs/internal/storage"
	"dfs/storagepb"
)

type StorageServer struct {
	storagepb.UnimplementedStorageServiceServer
}

func NewStorageServer() *StorageServer {
	return &StorageServer{}
}

func (s *StorageServer) StoreChunk(
	ctx context.Context,
	req *storagepb.StoreChunkRequest,
) (*storagepb.StoreChunkResponse, error) {

	path, err := storage.StoreChunk(
		req.ChunkId,
		req.Data,
	)

	if err != nil {
		return nil, err
	}

	return &storagepb.StoreChunkResponse{
		Path: path,
	}, nil
}

func (s *StorageServer) GetChunk(
	ctx context.Context,
	req *storagepb.GetChunkRequest,
) (*storagepb.GetChunkResponse, error) {

	data, err := storage.GetChunk(
		req.ChunkId,
	)

	if err != nil {
		return nil, err
	}

	return &storagepb.GetChunkResponse{
		Data: data,
	}, nil
}

func (s *StorageServer) DeleteChunk(
	ctx context.Context,
	req *storagepb.DeleteChunkRequest,
) (*storagepb.DeleteChunkResponse, error) {

	err := storage.DeleteChunk(
		req.ChunkId,
	)

	if err != nil {
		return nil, err
	}

	return &storagepb.DeleteChunkResponse{
		Success: true,
	}, nil
}
