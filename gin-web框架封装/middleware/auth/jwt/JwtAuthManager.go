package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type jwtAuthManager struct {
	secret string
	exp    time.Duration
	alg    string
}

func NewJwtAuth() *jwtAuthManager {
	return &jwtAuthManager{
		secret: "123",
		exp:    time.Hour * 24,
		alg:    "HS256",
	}
}

func (jwtAuth *jwtAuthManager) Check(c *gin.Context) bool {
	token := c.Request.Header.Get("Authorization")
	token = strings.Replace(token, "Beater ", "", -1)
	if token == "" {
		return false
	}
	var keyFun jwt.Keyfunc
	keyFun = func(token *jwt.Token) (interface{}, error) {
		b := []byte(jwtAuth.secret)
		return b, nil
	}
	authJwtToken, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, keyFun)
	if err != nil {
		fmt.Println(err)
		return false
	}

	c.Set("User", map[string]interface{}{
		"token": authJwtToken,
	})
	return authJwtToken.Valid
}

func (jwtAuth *jwtAuthManager) User(c *gin.Context) interface{} {
	var jwtToken *jwt.Token
	tokenStr := c.Request.Header.Get("Authorization")
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", -1)
	if tokenStr == "" {
		return map[interface{}]interface{}{}
	}

	var err error
	jwtToken, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		b := []byte(jwtAuth.secret)
		return b, nil
	})

	if err != nil {
		fmt.Println(err)
		return map[interface{}]interface{}{}
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		var user map[string]interface{}
		if err := json.Unmarshal([]byte(claims["user"].(string)), &user); err != nil {
			fmt.Println(err)
			return map[interface{}]interface{}{}
		}
		c.Set("User", map[interface{}]interface{}{
			"token": jwtToken,
			"user":  user,
		})
		return user
	} else {
		fmt.Println(ok)
		return map[interface{}]interface{}{}
	}
}

func (jwtAuth *jwtAuthManager) Login(req *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{} {
	token := jwt.New(jwt.GetSigningMethod(jwtAuth.alg))
	userStr, err := json.Marshal(user)
	if err != nil {
		return nil
	}
	token.Claims = jwt.MapClaims{
		"user": string(userStr),
		"exp":  time.Now().Add(jwtAuth.exp).Unix(),
	}

	tokenString, err := token.SignedString([]byte(jwtAuth.secret))
	if err != nil {
		return nil
	}
	return tokenString
}

func (jwtAuth *jwtAuthManager) Logout(req *http.Request, w http.ResponseWriter) bool {
	return true
}
