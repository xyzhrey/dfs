package storage

import (
	"os"
	"path/filepath"
)

func StoreChunk(chunkID string, data []byte) (string, error) {

	err := os.MkdirAll(
		"data/chunks",
		0755,
	)

	if err != nil {
		return "", err
	}

	path := filepath.Join(
		"data",
		"chunks",
		chunkID+".chunk",
	)

	err = os.WriteFile(
		path,
		data,
		0644,
	)

	if err != nil {
		return "", err
	}

	return path, nil
}

func GetChunk(chunkID string) ([]byte, error) {

	path := filepath.Join(
		"data",
		"chunks",
		chunkID+".chunk",
	)

	return os.ReadFile(path)
}

func DeleteChunk(chunkID string) error {

	path := filepath.Join(
		"data",
		"chunks",
		chunkID+".chunk",
	)

	return os.Remove(path)
}
