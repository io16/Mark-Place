package conf

import (
	"github.com/jinzhu/gorm"

	"log"
)

var (
	// DBCon is the connection handle
	// for the database
	DBCon *gorm.DB
)

func InitDb() (*gorm.DB) {
	DBCon, err := gorm.Open("postgres", "host=localhost user=postgres dbname=test sslmode=disable password=299792458")
	if err != nil {
		log.Panic(err)
	}

	return DBCon
}
