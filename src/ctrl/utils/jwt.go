package utils

import (
	"fmt"
	cfg "src/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type MJwt struct {
	r *gin.Context
}

type Claims struct {
	Username string `json:"username"`
	Auth     bool   `json:"auth"`
	jwt.RegisteredClaims
}

func GetJwt(r *gin.Context) *MJwt {
	return &MJwt{r}
}

func (j *MJwt) Get(is bool) (string, error) {
	username := j.r.PostForm("uname")
	claims := Claims{
		username,
		is,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "user",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(cfg.Cfg.Str.Key) // 替换为你的密钥
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (j *MJwt) Check() (bool, error) {
	tokenString := j.r.PostForm("token")
	if tokenString == "" {
		return false, fmt.Errorf("no token provided")
	}
	secret := []byte(cfg.Cfg.Str.Key) // 替换为你的密钥

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// 使用 claims.Username 等信息
		fmt.Printf("Username from token: %s\n", claims.Username)
		return true, nil
	} else {
		return false, fmt.Errorf("invalid token")
	}
}

func (j *MJwt) AuthCheck() {
	//中间件

}

func (j *MJwt) GetClaims() (*Claims, error) {
	tokenString := j.r.PostForm("token")
	if tokenString == "" {
		return nil, fmt.Errorf("no token provided")
	}
	secret := []byte(cfg.Cfg.Str.Key) // 替换为你的密钥

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// 使用 claims.Username 等信息
		fmt.Printf("Username from token: %s\n", claims.Username)
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
