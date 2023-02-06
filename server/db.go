package main

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"os"
)

func Connect() *gorm.DB {
    dbURL := os.Getenv("dbURL")

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&Destination{})

    return db
}