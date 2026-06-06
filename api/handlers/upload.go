package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"dfs/internal/chunker"
	"dfs/internal/models"
	"dfs/internal/repository"
	"dfs/storagepb"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UploadHandler struct {
Repo *repository.FileRepository,
StorageClient storagepb.StorageServiceClient
}



func NewUploadHandler(
	repo *repository.FileRepository,
	storageClient storagepb.StorageServiceClient,
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
		return c.JSON(http.StatusInternalServerError, map[string]string{
    "error": err.Error(),
})
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
		return c.JSON(http.StatusInternalServerError, map[string]string{
    "error": err.Error(),
})
	}

	defer dst.Close()

	_, err = dst.ReadFrom(src)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
    "error": err.Error(),
})
	}

	chunks, err := chunker.SplitToMemory(
		uploadPath,
		4*1024*1024,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
    "error": err.Error(),
})
	}

	err = h.Repo.CreateFile(models.File{
		ID:       fileID,
		Filename: fileHeader.Filename,
		Size:     fileHeader.Size,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
    "error": err.Error(),
})
	}



for _, chunk := range chunks {

	resp, err := h.StorageClient.StoreChunk(
		c.Request().Context(),
		&storagepb.StoreChunkRequest{
			ChunkId: chunk.ID,
			Data:    chunk.Data,
		},
	)

	if err != nil {
		return err
	}

	err = h.Repo.CreateChunk(
		models.Chunk{
			ID:         chunk.ID,
			FileID:     fileID,
			ChunkIndex: chunk.Index,
			Path:       resp.Path,
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
