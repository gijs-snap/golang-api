package controllers

import (
  "github.com/gin-gonic/gin"
  "github.com/dgrijalva/jwt-go"  
  "fmt"
  "gymgo.com/m/models"
  "net/http"
  "time"
  "os"
)

var user = models.User{
  ID: 1,
  Name: "name",
  Email: "email@email.com",
  Password: "password",
}

func CreateToken(userid uint) (string, error) {
	jwt_secret, exists := os.LookupEnv("JWT_SECRET")

    if !exists {
	  fmt.Println(".ENV not found")
	}
	
	var err error

	tokenClaims := jwt.MapClaims{}
	tokenClaims["authorized"] = true
	tokenClaims["user_id"] = userid

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte(jwt_secret))

	if err != nil {
	   return "", err
	}

	return tokenString, nil
}

func VerifyToken(c *gin.Context) (bool){
	cookie, err := c.Cookie("token")

	if err != nil {
	  c.JSON(404, "Validation cookie not found")
	  return false
	}
	
	claims := jwt.MapClaims{}
	jwt_secret, exists := os.LookupEnv("JWT_SECRET")
  
	if !exists {
	   c.JSON(404, ".ENV not found")
	}
	tkn, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwt_secret), nil
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

func LoginUser(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
	   c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
	   return
	}

	if user.Email != u.Email|| user.Password != u.Password {
	   c.JSON(http.StatusUnauthorized, "Please provide valid login details")
	   return
	}
	
	token, err := CreateToken(user.ID)
	if err != nil {
	   c.JSON(http.StatusUnprocessableEntity, err.Error())
	   return
	}
	expireCookie := time.Now().Add(time.Minute * 15).Unix()
	c.SetCookie("token", token, int(expireCookie), "/", "localhost", false, true)
	c.JSON(http.StatusOK, "Logged in")
}
 
func registerUser(c *gin.Context) {

}