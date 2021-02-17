package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gymgo.com/m/models"
)

// CreateToken creates a JWT token
func CreateToken(userid uint) (string, error) {
	jwtSecret, exists := os.LookupEnv("JWT_SECRET")

	if !exists {
		fmt.Println(".ENV not found")
	}

	var err error

	tokenClaims := jwt.MapClaims{}
	tokenClaims["authorized"] = true
	tokenClaims["user_id"] = userid

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies a users JWT token
func VerifyToken(c *gin.Context) bool {
	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(404, "Validation cookie not found")
		return false
	}

	claims := jwt.MapClaims{}
	jwtSecret, exists := os.LookupEnv("JWT_SECRET")

	if !exists {
		c.JSON(404, ".ENV not found")
	}
	tkn, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(404, "No valid secret")
			return false
		}
		fmt.Println(err)
		c.JSON(404, "Error parsing key")
		return false
	}
	if !tkn.Valid {
		c.JSON(404, "No valid token")
		return false
	}
	fmt.Println(claims)
	return true
}

// LoginUser checks passed parameters against user in db to log in
func LoginUser(c *gin.Context) {
	var jsonData models.User
	if c.BindJSON(&jsonData) != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	tempUser := models.User{}
	res := models.DB.Select("*").Where("email = ?", jsonData.Email).First(&tempUser)

	if res.RowsAffected == 0 {
		c.JSON(401, "Email not found")
		return
	}

	if tempUser.Password != jsonData.Password {
		c.JSON(http.StatusUnauthorized, "Invalid password")
		return
	}

	token, err := CreateToken(tempUser.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	expireCookie := time.Now().Add(time.Minute * 15).Unix()
	c.SetCookie("token", token, int(expireCookie), "/", "localhost", false, true)
	c.JSON(200, "Logged in")
}

// RegisterUser checks if user exists, if not creates new user
func RegisterUser(c *gin.Context) {
	var jsonData models.User
	if c.BindJSON(&jsonData) != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	fmt.Println(jsonData)

	if jsonData.Email == "" || jsonData.Name == "" || jsonData.Password == "" {
		c.JSON(400, "Missing parameters")
		return
	}

	tempUser := models.User{}
	res := models.DB.Select("*").Where("email = ?", jsonData.Email).First(&tempUser)

	if res.RowsAffected != 0 {
		c.JSON(http.StatusUnauthorized, "Email already in use")
		return
	}

	newUser := models.User{Name: jsonData.Name, Email: jsonData.Email, Password: jsonData.Password}
	models.DB.Create(&newUser)
	c.JSON(200, "Account created")
}
