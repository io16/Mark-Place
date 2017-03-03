package models

import (
	"github.com/labstack/echo"
	"encoding/json"
	"net/http"
	"log"
	"regexp"
	"strconv"
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"../conf/"
	"fmt"
)

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"size:255"`
	Login string `json:"login" gorm:"size:255"`
	Email string `json:"email"gorm:"size:255"`
	Pass  string `json:"pass"gorm:"size:255"`
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
func saveUserToDB(user User) bool {

	db := conf.InitDb()
	defer db.Close()
	//log.Print(user)

	test := User{}
	db.Where("login = ?", user.Login).First(&test)
	if (test.ID == 0) {
		fmt.Println("user Created")
		password := []byte (user.Pass)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		user.Pass = string(hashedPassword)
		db.Create(&user)
		return true

	} else {
		fmt.Print("user is alerady exist")

		return false

	}
}
func AddUser(c echo.Context) error {
	user := User{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	userStatus := isUserValid(user)
	if userStatus {
		userStatus = saveUserToDB(user)

	}
	if err != nil {
		log.Printf("Failed processing addUser request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.String(http.StatusOK, strconv.FormatBool(userStatus)) // if user created -- return true
}