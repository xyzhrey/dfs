package main

import (
	"net/http"
	"fmt"
	"dfs/api/handlers"
	"dfs/internal/config"
	"dfs/internal/postgres"
	"dfs/internal/repository"
	"log"

  "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()

if err != nil {
    log.Fatal(".env not found")
}

	pool, err := postgres.NewPool(config.DBURL())
	if err != nil {
		panic(err)
	}
	defer pool.Close()
	repo := repository.NewFileRepository(pool)
	
	uploadHandler := handlers.NewUploadHandler(repo)
	downloadHandler := handlers.NewDownloadHandler(repo)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "DFS Gateway Running")
	})
	fmt.Println("DB_URL =", config.DBURL())
	e.POST("/upload", uploadHandler.Upload)
	e.GET(
		"/download/:id",
		downloadHandler.Download,
	)
	listHandler := handlers.NewListHandler(repo)

e.GET("/files", listHandler.List)
deleteHandler := handlers.NewDeleteHandler(repo)

e.DELETE(
	"/file/:id",
	deleteHandler.Delete,
)

for _, r := range e.Routes() {
    println(r.Method, r.Path)
}

e.Logger.Fatal(e.Start(":8080"))
}
