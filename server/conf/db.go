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

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"size:255"`
	Login string `json:"login" gorm:"size:255"`
	Email string `json:"email"gorm:"size:255"`
	Pass  string `json:"pass"gorm:"size:255"`
	Role  Role
	RoleID int
}

type Role struct {
	gorm.Model
	UserRole  string
}

type Basket struct {
	gorm.Model
	User  User
	Order []Order
	Date  string
}
type Category struct {
	gorm.Model
	Name string
}
type Order struct {
	gorm.Model
	Product Product
	Count   int
	Status  string
}
type Product struct {
	gorm.Model
	Title       string        `json:"title" `
	Description string        `json:"description" `
	Count       int           `json:"count" `
	Price       float32       `json:"price" `
	Url         string        `json:"url" `
	Category    []Category	  `json:"category" `
}

func InitDb() (*gorm.DB) {
	DBCon, err := gorm.Open("postgres", "host=localhost user=postgres dbname=test sslmode=disable password=299792458")
	if err != nil {
		log.Panic(err)
	}

	return DBCon
}
