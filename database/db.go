package database

import (
	"fmt"
	"log"
	"sesi_12/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "password"
	port     = "5432"
	dbname   = "db_go_sql"
	db       *gorm.DB
	err      error
)

func StartDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "inventory.",
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatal("error connecting to db", err)
	}

	log.Println("successfully connected to")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
