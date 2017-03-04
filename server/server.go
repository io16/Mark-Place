package main

import (
	//"fmt"
	"./models/"
	"./conf/"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

func main() {
	db := conf.InitDb()
	defer db.Close()

	db.AutoMigrate(&models.User{})
	e := echo.New()
	e.Static("/", "assets")
	e.File("/", "public/index.html")
	e.File("/registration", "public/registration.html")

	e.POST("/user", models.AddUser)

	e.Logger.Fatal(e.Start(":1324"))

}