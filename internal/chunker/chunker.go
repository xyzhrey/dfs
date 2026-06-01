package chunker

import (
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type Chunk struct {
	ID    string
	Path  string
	Size  int64
	Index int
}

func Split(filePath string, chunkSize int64) ([]Chunk, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var chunks []Chunk
	buffer := make([]byte, chunkSize)
	index := 0

	for {
		n, err := file.Read(buffer)

		if err != nil && err != io.EOF {
			return nil, err
		}

		if n == 0 {
			break
		}

		chunkID := uuid.New().String()

		chunkPath := filepath.Join(
			"data",
			"chunks",
			chunkID+".chunk",
		)

		err = os.WriteFile(chunkPath, buffer[:n], 0644)
		if err != nil {
			return nil, err
		}

		chunks = append(chunks, Chunk{
			ID:    chunkID,
			Path:  chunkPath,
			Size:  int64(n),
			Index: index,
		})

		index++
	}

	return chunks, nil
}
