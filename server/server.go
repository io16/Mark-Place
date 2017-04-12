package main

import (
	"net/http"
	"time"
	"./models/"
	"./conf/"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"log"

	"html/template"

	"io"
)

type UserData struct {
	Name string
	Role string
}
// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name string `json:"name"`
	Role string   `json:"admin"`
	jwt.StandardClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	user, err := models.UserAuthentication(username, password)
	if err != nil {
		log.Println(err)

	} else {

		claims := &jwtCustomClaims{
			user.Name,
			user.Role.UserRole,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
		//expiration := time.Now().Add(365 * 24 * time.Hour)
		//cookie    :=    http.Cookie{Name: "token",Value:t,Expires:expiration}
		//
		//c.SetCookie(&cookie)
		//userDate := UserData{}
		//
		//userDate.Name = user.Name
		//userDate.Role = "admin"
		//log.Println(userDate)
		//return c.Render(http.StatusOK, "index",userDate )

	}

	return echo.ErrUnauthorized
}


func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	if claims.Role != "admin" {
		return c.String(http.StatusOK, "Welcome " + name + "! access denied" + " your status is " + claims.Role)
	}
	return c.String(http.StatusOK, "Welcome " + name + "your status is " + claims.Role)
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {

	return c.Render(http.StatusOK, "hello", "test")
}

func Index(c echo.Context) error {

	userDate := UserData{}
	if c.Get("user") != nil{
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		name := claims.Name

		userDate.Name = name
		userDate.Role = claims.Role
	}
	qor

	return c.Render(http.StatusOK, "mainTemplate", userDate)
}
func main() {

	db := conf.InitDb()
	defer db.Close()
	//db.AutoMigrate(&conf.User{}, &conf.Order{}, &conf.Basket{}, &conf.Product{}, &conf.Category{}, &conf.Role{})
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("public/*.html")),
	}

	e.Renderer = t

	e.Static("/", "assets")
	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: " time=${time_rfc3339} method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.Recover())

	e.POST("/login", login)
	e.POST("/user", models.AddUser)
	e.POST("/product", models.AddProduct)

	e.File("/login", "public/login.html")
	e.File("/", "public/index.html")
	e.File("/registration", "public/registration.html")
	e.File("/admin", "public/adminProducts.html")
	e.File("/login", "public/login.html")

	// Restricted group
	r := e.Group("/restricted")


	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", restricted)
	r.GET("/hello", Hello)
	r.GET("/test", Index)

	e.Logger.Fatal(e.Start(":1323"))
}