package main

import (
	"fmt"
	"./models/"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=test sslmode=disable password=299792458")
	defer db.Close()
	if err != nil {

		fmt.Print(err)
	}
	db.AutoMigrate(&models.User{})
	e := echo.New()
	e.POST("/user", models.AddUser)

	e.Logger.Fatal(e.Start(":1324"))

}