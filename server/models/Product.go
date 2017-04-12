package models

import (
	"net/http"

	"github.com/labstack/echo"
	"../conf/"
	"encoding/json"

	"log"
)

func AddProduct(c echo.Context) error {
	product := conf.Product{}

	defer c.Request().Body.Close()
	//err:= json.Unmarshal([]byte(c.FormValue("data")), &product)

	err := json.NewDecoder(c.Request().Body).Decode(&product)
	if err != nil {
		log.Print(err)
	}

	db := conf.InitDb()
	defer db.Close()
	db.Create(&product)

	return c.String(http.StatusOK, "true")
}

