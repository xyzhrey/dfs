package chunker

import (
	"io"
	"os"
)

func Reassemble(
	chunkPaths []string,
	outputPath string,
) error {

	out, err := os.Create(outputPath)

	if err != nil {
		return err	}

	defer out.Close()

	for _, path := range chunkPaths {

		in, err := os.Open(path)

		if err != nil {
			return err
		}

		_, err = io.Copy(out, in)

		in.Close()

		if err != nil {
			return err
		}
	}

	return nil
}
