package config

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(sqlite.Open("server-app-bet3.0.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Erreur de connexion à la base de données :", err)
		panic("Échec de la connexion à la base de données.")
	}

	fmt.Println("Connexion à la base de données établie.")
	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
