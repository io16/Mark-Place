package models

import (
	"github.com/labstack/echo"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"golang.org/x/crypto/bcrypt"
	//"github.com/jinzhu/gorm"
	"../conf/"
	"fmt"
	"github.com/labstack/gommon/log"

)

type Place struct {
	ID          int
	Name        string
	Description string
	TownID      int
	Town        Town
}

type Town struct {
	ID   int
	Name string
}

func AddUser(c echo.Context) error {

	//u := new(conf.User)
	//if err := c.Bind(u); err != nil {
	//	return err
	//}
	//return c.JSON(http.StatusOK, u)
	user := conf.User{}

	json.Unmarshal([]byte(c.FormValue("data")), &user)

	userStatus := isUserValid(user)
	if userStatus {
		userStatus = saveUserToDB(user)

	}

	return c.String(http.StatusOK, strconv.FormatBool(userStatus)) // if user created -- return true
}
func isUserValid(user conf.User) bool {
	validationStatus := false
	r, _ := regexp.Compile("^[A-Za-z0-9_-]{3,50}$")

	if r.MatchString(user.Login) &&  r.MatchString(user.Name) && r.MatchString(user.Pass) && isEmailValid(user.Email) {
		validationStatus = true
	} else {
		log.Print("user invalid")
	}

	return validationStatus

}

func isEmailValid(email string) bool {
	r, _ := regexp.Compile("^[a-zA-z_]([A-Za-z0-9_.-]{0,100})@([a-z]{2,8}[.][a-z]{2,8})$")

	return r.MatchString(email)
}

func saveUserToDB(user conf.User) bool {

	db := conf.InitDb()
	defer db.Close()

	temp := conf.User{}

	db.Where("login = ?", user.Login).First(&temp)

	if (temp.ID == 0) {
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
		fmt.Println("user is alerady exist")

		return false

	}
	return true
}

func UserAuthentication(login, password string) (conf.User, error) {
	user := conf.User{}

	db := conf.InitDb()
	defer db.Close()

	db.Where("login = ?", login).First(&user)
	db.Model(user).Related(&user.Role)

	err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(password))
	fmt.Println(user)
	if err != nil {
		fmt.Print(err)
		return user, err
	}

	return user, nil
}