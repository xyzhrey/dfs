package main

import (
	"net/http"

	"dfs/api/handlers"
	"dfs/internal/config"
	"dfs/internal/postgres"
	"dfs/internal/repository"

	"github.com/labstack/echo/v4"
)

func main() {

	pool, err := postgres.NewPool(config.DBURL())
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	repo := repository.NewFileRepository(pool)
	uploadHandler := handlers.NewUploadHandler(repo)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "DFS Gateway Running")
	})

	e.POST("/upload", uploadHandler.Upload)

	e.Logger.Fatal(e.Start(":8080"))
}
