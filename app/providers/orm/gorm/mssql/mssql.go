package mssql

import (
	"os"

	"github.com/jinzhu/gorm"
	// gorm dialects
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// Open :
func Open(dbType string, conn string) (*gorm.DB, error) {
	return gorm.Open(os.Getenv("DATABASE"), os.Getenv("DATABASE_CONN"))
}
