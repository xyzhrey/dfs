package chunker

import (
	"io"
	"os"

	"github.com/google/uuid"
)

type MemoryChunk struct {
	ID    string
	Data  []byte
	Size  int64
	Index int
}

func SplitToMemory(
	filePath string,
	chunkSize int64,
) ([]MemoryChunk, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var chunks []MemoryChunk

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

		data := make([]byte, n)
		copy(data, buffer[:n])

		chunks = append(
			chunks,
			MemoryChunk{
				ID:    uuid.New().String(),
				Data:  data,
				Size:  int64(n),
				Index: index,
			},
		)

		index++
	}

	return chunks, nil
}
