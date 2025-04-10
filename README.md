# URL Shortener in Go 🐹

A simple URL shortening service built using **Go**, **MySQL**, and **Gin**.

## Features

- Generate short links for long URLs
- Redirect using short codes
- Persistent storage with MySQL
- Environment configuration via `.env`
- Ready to deploy or extend

## Technologies Used

- Golang (Gin framework)
- MySQL
- MySQL Workbench (local DB management)
- RESTful API structure

## 📦 Project Structure
|-------url-shortener
|       |------shoreten-service
               |----handler
                    |--handler.go
               |----model
                    |--url.go
               |----repository
                    |--db.go
                    |--url_repo.go
               |----service
                    |--service.go
                main.go
        .gitignore
        go.mod
        go.sum

