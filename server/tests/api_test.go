package tests

import (
    "testing"
    "gorm.io/driver/postgres"
	"gorm.io/gorm"
    "vacation-planner/models"
)

// Test function for DB connection
func TestDBConnection(t *testing.T) {
    dsn := "host=24.199.69.57 user=dev password=K@Gf$+$n?EBV^K3% dbname=dev port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    // If opening the database amounts to an error
	if err != nil {
		t.Fatalf("Not successfully connected to database.")
	} else {
        // Else, use db so Go doesn't cry about it not being used
        db.AutoMigrate(&models.User{}, &models.SavedLocation{})
    }
}
