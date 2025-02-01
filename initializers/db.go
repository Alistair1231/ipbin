package initializers

import (
	"os"

	"github.com/Alistair1231/ipbin/models"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

// for postgres, see https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL

var DB *gorm.DB

func ConnectToDatabase() {
	// github.com/mattn/go-sqlite3
	var err error
	DB, err = gorm.Open(sqlite.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Paste{})
	DB.AutoMigrate(&models.User{})
}
