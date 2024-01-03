package database

import (
	"github.com/Galbeyte1/go-web-dev/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	/*
	  DB_USERNAME: "galbeyte"
      DB_PASSWORD: "devenv"
      DB_DB: "goServerDatabase"
      DB_HOST: "db"
      DB_PORT: "5432"
      DB_SSL_MODE: "disable"
	*/

	dsn := "host=db user=galbeyte password=devenv dbname=goServerDatabase"
	// dsn := "host=localhost user=galbeyte password=devenv dbname=goserverdatabase"

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database.")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}