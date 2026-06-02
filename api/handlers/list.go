package handlers

import (
	"net/http"

	"dfs/internal/repository"

	"github.com/labstack/echo/v4"
)

type ListHandler struct {
	Repo *repository.FileRepository
}

func NewListHandler(
	repo *repository.FileRepository,
) *ListHandler {

	return &ListHandler{
		Repo: repo,
	}
}

func (h *ListHandler) List(
	c echo.Context,
) error {

	files, err := h.Repo.ListFiles()

	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		files,
	)
}
