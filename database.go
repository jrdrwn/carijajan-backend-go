package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	u, err := url.Parse(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to parse connection string: %v", err)
	}

	// Extract user info
	user := u.User.Username()
	password, _ := u.User.Password()

	// Build DSN
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		u.Hostname(),
		u.Port(),
		user,
		password,
		strings.TrimPrefix(u.Path, "/"),
		u.Query().Get("sslmode"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db

	if err != nil {
		log.Fatal("failed to connect database")
	}
}
