package config

import (
	"fmt"
	"github.com/felixlambertv/nextalent/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	connectionUrl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		"postgres_db", "user", "admin", "nextalent", "5432", "disable",
	)

	db, err := gorm.Open(postgres.Open(connectionUrl))
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(
		&model.Person{},
	)

	return db
}
