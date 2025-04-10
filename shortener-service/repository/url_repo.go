package repository

import (
	"database/sql"
	"log"
)

type URLRepository struct {
	DB *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{DB: db}
}

func (repo *URLRepository) SaveURL(shortCode, OriginalURL string) error {
	_, err := repo.DB.Exec("INSERT INTO urls (short_code, original_url) values(?,?)", shortCode, OriginalURL)
	if err != nil {
		log.Println("Failed to insert")
	}
	return err
}

func (repo *URLRepository) GetOriginalURL(shortCode string) (string, error) {
	var originalURL string
	err := repo.DB.QueryRow("SELECT original_url from urls WHERE short_code = ?", shortCode).Scan(&originalURL)
	return originalURL, err
}

func (repo *URLRepository) GetShortCodeByOriginalURL(OriginalURL string) (string, error) {
	var shortCode string
	err := repo.DB.QueryRow("SELECT short_code from urls WHERE original_url = ?", OriginalURL).Scan(&shortCode)
	return shortCode, err
}
