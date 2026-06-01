package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"dfs/internal/chunker"
	"dfs/internal/models"
	"dfs/internal/repository"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UploadHandler struct {
	Repo *repository.FileRepository
}

func NewUploadHandler(
	repo *repository.FileRepository,
) *UploadHandler {
	return &UploadHandler{
		Repo: repo,
	}
}

func (h *UploadHandler) Upload(
	c echo.Context,
) error {

	fileHeader, err := c.FormFile("file")

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "file required",
			},
		)
	}

	src, err := fileHeader.Open()

	if err != nil {
		return err
	}

	defer src.Close()

	fileID := uuid.New().String()

	uploadPath := filepath.Join(
		"data",
		"uploads",
		fileID+"_"+fileHeader.Filename,
	)

	dst, err := os.Create(uploadPath)

	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = dst.ReadFrom(src)

	if err != nil {
		return err
	}

	chunks, err := chunker.Split(
		uploadPath,
		4*1024*1024,
	)

	if err != nil {
		return err
	}

	err = h.Repo.CreateFile(models.File{
		ID:       fileID,
		Filename: fileHeader.Filename,
		Size:     fileHeader.Size,
	})

	if err != nil {
		return err
	}

	for _, chunk := range chunks {

		err = h.Repo.CreateChunk(
			models.Chunk{
				ID:         chunk.ID,
				FileID:     fileID,
				ChunkIndex: chunk.Index,
				Path:       chunk.Path,
				Size:       chunk.Size,
			},
		)

		if err != nil {
			return err
		}
	}

	return c.JSON(
		http.StatusCreated,
		map[string]string{
			"file_id": fileID,
		},
	)
}
