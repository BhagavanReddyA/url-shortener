package repository

import (
	"database/sql"
	"log"
	"time"
)

type URLRepository struct {
	DB *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{DB: db}
}

func (repo *URLRepository) SaveURL(shortCode, OriginalURL string) error {
	expirationTime := time.Now().Add(24 * time.Hour)
	FormattedExpirationTime := expirationTime.Format("2006-01-02 15:04:05")
	_, err := repo.DB.Exec("INSERT INTO urls (short_code, original_url, expiration_date, analytics) values(?,?,?,?)", shortCode, OriginalURL, FormattedExpirationTime, 0)
	if err != nil {
		log.Println("Failed to insert")
	}
	return err
}

func (repo *URLRepository) GetOriginalURL(shortCode string) (string, error) {
	var originalURL string
	err := repo.DB.QueryRow("SELECT original_url from urls WHERE short_code = ?", shortCode).Scan(&originalURL)
	if err == nil {
		_, analyitcserr := repo.DB.Exec("UPDATE urls SET analytics = analytics +1 WHERE short_code = ?", shortCode)
		if analyitcserr != nil {
			log.Println("Failed to update analytics")
			return shortCode, analyitcserr
		}
	}
	return originalURL, err
}

func (repo *URLRepository) GetShortCodeByOriginalURL(OriginalURL string) (string, error) {
	var shortCode string
	err := repo.DB.QueryRow("SELECT short_code from urls WHERE original_url = ?", OriginalURL).Scan(&shortCode)
	return shortCode, err
}

func (repo *URLRepository) DeleteExpiredURLS() error {
	expirationTime := time.Now().Format("2006-01-02 15:04:05")
	_, err := repo.DB.Exec("DELETE FROM urls WHERE expiration_date <= ?", expirationTime)
	log.Println("the error is: ", err)
	return err
}

func (repo *URLRepository) URLAnalytics(shortCode string) (int, error) {
	var urlAnalytics int
	err := repo.DB.QueryRow("SELECT analytics FROM urls WHERE short_code = ?", shortCode).Scan(&urlAnalytics)
	return urlAnalytics, err
}
