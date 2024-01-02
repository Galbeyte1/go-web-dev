package database

import (
	"github.com/Galbeyte1/go-web-dev/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



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
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database.")
	}

	connection.AutoMigrate(&models.User{})
}