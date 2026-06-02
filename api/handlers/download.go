package handlers

import (
	"net/http"
	"path/filepath"

	"dfs/internal/chunker"
	"dfs/internal/repository"

	"github.com/labstack/echo/v4"
)

type DownloadHandler struct {
	Repo *repository.FileRepository
}

func NewDownloadHandler(
	repo *repository.FileRepository,
) *DownloadHandler {

	return &DownloadHandler{
		Repo: repo,
	}
}

func (h *DownloadHandler) Download(
	c echo.Context,
) error {

	fileID := c.Param("id")

	file, err := h.Repo.GetFile(fileID)

	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			map[string]string{
				"error": "file not found",
			},
		)
	}

	chunks, err := h.Repo.GetChunksByFileID(fileID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
    "error": err.Error(),
})
	}

	var chunkPaths []string

	for _, chunk := range chunks {
		chunkPaths = append(
			chunkPaths,
			chunk.Path,
		)
	}

	outputPath := filepath.Join(
		"data",
		"tmp",
		fileID+"_"+file.Filename,
	)

	err = chunker.Reassemble(
		chunkPaths,
		outputPath,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
    "error": err.Error(),
})
	}

	return c.Attachment(
		outputPath,
		file.Filename,
	)
}
