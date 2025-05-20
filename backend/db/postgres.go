package db

import (
	"log"

	"github.com/wellminozzo/LoginGo-React-teste/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func DbConnection() {
	dsn := "host=localhost user=postgres password=project dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Erro ao migrar o banco:", err)
	}
	log.Println("Banco de dados conectado com sucesso!")

}
