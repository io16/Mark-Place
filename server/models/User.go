package models

import (
	"github.com/labstack/echo"
	"encoding/json"
	"net/http"
	"log"
	"regexp"
	"strconv"
)

type User struct {
	Name  string `json:"name"`
	Login string `json:"login"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

func isEmailValid(email string) bool {
	r, _ := regexp.Compile("^[a-zA-z_]([A-Za-z0-9_.-]{0,100})@([a-z]{2,8}[.][a-z]{2,8})$")

	return r.MatchString(email)
}
func isUserValid(user User) bool {
	validationStatus := false
	r, _ := regexp.Compile("^[A-Za-z0-9_-]{3,50}$")

	if r.MatchString(user.Login) &&  r.MatchString(user.Name) && r.MatchString(user.Pass) && isEmailValid(user.Email) {
		validationStatus = true
	}

	return validationStatus

}
func AddUser(c echo.Context) error {
	user := User{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	userStatus := isUserValid(user)
	log.Printf("this is your user: %#v /n \n", user)
	if err != nil {
		log.Printf("Failed processing addUser request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.String(http.StatusOK, strconv.FormatBool(userStatus))
}