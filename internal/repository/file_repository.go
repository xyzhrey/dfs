package repository

import (
	"context"

	"dfs/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FileRepository struct {
	DB *pgxpool.Pool
}

func NewFileRepository(db *pgxpool.Pool) *FileRepository {
	return &FileRepository{
		DB: db,
	}
}

func (r *FileRepository) CreateFile(
	file models.File,
) error {

	_, err := r.DB.Exec(
		context.Background(),
		`
		INSERT INTO files(
			id,
			filename,
			size
		)
		VALUES($1,$2,$3)
		`,
		file.ID,
		file.Filename,
		file.Size,
	)

	return err
}

func (r *FileRepository) CreateChunk(
	chunk models.Chunk,
) error {

	_, err := r.DB.Exec(
		context.Background(),
		`
		INSERT INTO chunks(
			id,
			file_id,
			chunk_index,
			path,
			size
		)
		VALUES($1,$2,$3,$4,$5)
		`,
		chunk.ID,
		chunk.FileID,
		chunk.ChunkIndex,
		chunk.Path,
		chunk.Size,
	)

	return err
}
