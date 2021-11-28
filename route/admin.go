package route

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 使用全局变量
var jwtKey = []byte("jwt-tocken")

// 使用 mysql 存储
var users = map[string]string{
	"admin":  "admin",
	"tyrone": "tyrone",
}

// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type Claims struct {
	Username string `json:"username"`
	Password string `json:"passwd"`
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(c.Request.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		c.JSON(http.StatusBadRequest, gin.H{"error": "json 解析失败"})
		return
	}

	// Get the expected password from our in memory map
	expectedPassword, ok := users[creds.Username]

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if !ok || expectedPassword != creds.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误"})
		return
	}

	tokenString, err := generateJwtToken(creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	//http.SetCookie(c, &http.Cookie{
	//	Name:    "token",
	//	Value:   tokenString,
	//	Expires: expirationTime,
	//})
	c.SetCookie("jwt-token", tokenString, 60, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "登陆成功"})
	return
}

func generateJwtToken(creds Credentials) (token string, err error) {

	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
		Password: creds.Password,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	jwtTokenString, err := jwtToken.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error

		return "", fmt.Errorf("token 生成失败")
	}

	return jwtTokenString, nil
}

func Welcome(c *gin.Context) (string, error) {
	// 从请求中获取 Cookie jwt 令牌
	jwtToken, err := c.Cookie("jwt-token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "token 认证失败",
		})
		return "", err
	}

	// init claims
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token 未认证",
			})
			return "", err
		}
	}

	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "token 未认证",
		})
		return "", err
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Welcome %s", claims.Username),
	})
	return claims.Username, nil
}
