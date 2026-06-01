package main

import (
	"fmt"

	"dfs/internal/chunker"
)

func main() {
	chunks, err := chunker.Split("test.txt", 4)

	if err != nil {
		panic(err)
	}

	fmt.Println(chunks)
}
