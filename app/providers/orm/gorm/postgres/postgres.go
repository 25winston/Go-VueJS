package postgres

import (
	"os"

	"github.com/jinzhu/gorm"
	// gorm dialects
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Open :
func Open(dbType string, conn string) (*gorm.DB, error) {
	return gorm.Open(os.Getenv("DATABASE"), os.Getenv("DATABASE_CONN"))
}
