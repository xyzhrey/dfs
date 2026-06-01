package models

type File struct {
	ID       string
	Filename string
	Size     int64
}

type Chunk struct {
	ID         string
	FileID     string
	ChunkIndex int
	Path       string
	Size       int64
}
