package handler

import (
	"log"
	"net/http"

	"github.com/BhagavanReddyA/url-shortener/shortener-service/model"
	"github.com/BhagavanReddyA/url-shortener/shortener-service/repository"
	"github.com/BhagavanReddyA/url-shortener/shortener-service/service"
	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	Repo *repository.URLRepository
}

func NewURLHandler(repo *repository.URLRepository) *URLHandler {
	return &URLHandler{Repo: repo}
}

func (h *URLHandler) ShortenURL(c *gin.Context) {
	var req model.URLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("The request is invalid")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	existingShortCode, err := h.Repo.GetShortCodeByOriginalURL(req.OriginalURL)
	if err == nil {
		existingShortCodeRes := model.URLResponse{
			ShortURL: "http://localhost:8080/" + existingShortCode,
		}
		c.JSON(http.StatusOK, existingShortCodeRes)
		return
	}

	shortCode := service.GenerateShortCode(6)
	generateerr := h.Repo.SaveURL(shortCode, req.OriginalURL)
	if generateerr != nil {
		log.Println("Failed to save url")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL"})
		return
	}

	res := model.URLResponse{
		ShortURL: "http://localhost:8080/" + shortCode,
	}
	c.JSON(http.StatusOK, res)
}

func (h *URLHandler) Redirect(c *gin.Context) {
	shortCode := c.Param("code")
	OriginalURL, err := h.Repo.GetOriginalURL(shortCode)
	if err != nil {
		log.Println("Failed to fetch the original url")
		return
	}
	c.Redirect(http.StatusMovedPermanently, OriginalURL)
}
