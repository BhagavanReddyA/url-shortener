package main

import (
	"log"

	"github.com/BhagavanReddyA/url-shortener/shortener-service/handler"
	"github.com/BhagavanReddyA/url-shortener/shortener-service/repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	db, err := repository.InitDB()
	if err != nil {
		log.Fatal("DB Connection Failed: ", err)
	}
	defer db.Close()

	urlRepo := repository.NewURLRepository(db)
	urlHandler := handler.NewURLHandler(urlRepo)

	r := gin.Default()
	r.POST("/shorten", urlHandler.ShortenURL)
	r.GET("/:code", urlHandler.Redirect)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
