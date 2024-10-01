package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func Connect(config *Config) {
	dns := fmt.Sprintf(
		"host=%s port=%s password=%s user=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Password, config.User, config.DBName, config.SSLMode,
	)
	d, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		log.Println("DB connection successful......")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
