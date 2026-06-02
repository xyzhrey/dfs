package handlers

import (
	"net/http"
	"os"

	"dfs/internal/repository"

	"github.com/labstack/echo/v4"
)

type DeleteHandler struct {
	Repo *repository.FileRepository
}

func NewDeleteHandler(
	repo *repository.FileRepository,
) *DeleteHandler {

	return &DeleteHandler{
		Repo: repo,
	}
}

func (h *DeleteHandler) Delete(
	c echo.Context,
) error {

	fileID := c.Param("id")

	chunks, err := h.Repo.GetChunksByFileID(fileID)

	if err != nil {
		return err
	}

	for _, chunk := range chunks {
		_ = os.Remove(chunk.Path)
	}

	err = h.Repo.DeleteChunks(fileID)

	if err != nil {
		return err
	}

	err = h.Repo.DeleteFile(fileID)

	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"message": "deleted",
		},
	)
}
