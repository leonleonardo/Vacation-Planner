package database

import (
	"fmt"
	"log"
	"vacation-planner/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	//default code found in GormPG docs
	dsn := "host=24.199.69.57 user=dev password=K@Gf$+$n?EBV^K3% dbname=dev port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//test migrating User struct
	db.AutoMigrate(&models.User{}, &models.SavedLocation{})

	fmt.Println("successful connection")

	return db
}
