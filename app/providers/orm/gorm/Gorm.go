package gorm

import (
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	// gorm dialects

	"gmmshops.go/app/providers/logs"
	"gmmshops.go/app/providers/orm/gorm/mysql"
)

var (
	// db : gorm Database Connections
	db *gorm.DB
)

// Initial :
func Initial() *gorm.DB {

	if strings.ToLower(os.Getenv("DATABASE")) == "mysql" {
		db, err := mysql.Open(os.Getenv("DATABASE"), os.Getenv("DATABASE_CONN"))
		if err != nil {
			logs.Logger().Fatal("Can't connect database.")
			return nil
		}
		defer db.Close()

	} else if strings.ToLower(os.Getenv("DATABASE")) == "mssql" {
		db, err := mysql.Open(os.Getenv("DATABASE"), os.Getenv("DATABASE_CONN"))
		if err != nil {
			logs.Logger().Fatal("Can't connect database.")
			return nil
		}
		defer db.Close()

	} else if strings.ToLower(os.Getenv("DATABASE")) == "postgres" {
		db, err := mysql.Open(os.Getenv("DATABASE"), os.Getenv("DATABASE_CONN"))
		if err != nil {
			logs.Logger().Fatal("Can't connect database.")
			return nil
		}
		defer db.Close()

	} else {
		logs.Logger().Info("Database: ", strings.ToLower(os.Getenv("DATABASE")), ", ", os.Getenv("DATABASE_CONN"))
		logs.Logger().Info("Database: disabled")
		return nil
	}

	logs.Logger().Info("GORM: OK")

	return db

}

// Gorm :
func Gorm() *gorm.DB {
	if db == nil {
		db = Initial()
	}

	return db

}
