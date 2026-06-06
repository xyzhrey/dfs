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


func (r *FileRepository) DeleteChunks(
	fileID string,
) error {

	_, err := r.DB.Exec(
		context.Background(),
		`
		DELETE FROM chunks
		WHERE file_id=$1
		`,
		fileID,
	)

	return err
}

func (r *FileRepository) DeleteFile(
	fileID string,
) error {

	_, err := r.DB.Exec(
		context.Background(),
		`
		DELETE FROM files
		WHERE id=$1
		`,
		fileID,
	)

	return err
}

func (r *FileRepository) GetFile(
	fileID string,
) (models.File, error) {

	var file models.File

	err := r.DB.QueryRow(
		context.Background(),
		`
		SELECT
			id,
			filename,
			size
		FROM files
		WHERE id=$1
		`,
		fileID,
	).Scan(
		&file.ID,
		&file.Filename,
		&file.Size,
	)

	return file, err
}
func (r *FileRepository) ListFiles() ([]models.File, error) {

	rows, err := r.DB.Query(
		context.Background(),
		`
		SELECT
			id,
			filename,
			size
		FROM files
		ORDER BY created_at DESC
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var files []models.File

	for rows.Next() {

		var file models.File

		err := rows.Scan(
			&file.ID,
			&file.Filename,
			&file.Size,
		)

		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}
func (r *FileRepository) GetChunksByFileID(
	fileID string,
) ([]models.Chunk, error) {

	rows, err := r.DB.Query(
		context.Background(),
		`
		SELECT
			id,
			file_id,
			chunk_index,
			path,
			size
		FROM chunks
		WHERE file_id=$1
		ORDER BY chunk_index
		`,
		fileID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var chunks []models.Chunk

	for rows.Next() {

		var chunk models.Chunk

		err := rows.Scan(
			&chunk.ID,
			&chunk.FileID,
			&chunk.ChunkIndex,
			&chunk.Path,
			&chunk.Size,
		)

		if err != nil {
			return nil, err
		}

		chunks = append(chunks, chunk)
	}

	return chunks, nil
}

