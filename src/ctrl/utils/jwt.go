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

func (j *MJwt) Get(is bool, name string) (string, error) {
	claims := Claims{
		name,
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

func (j *MJwt) GetClaims() (*Claims, error) {
	tokenString := j.r.PostForm("token")

	if tokenString == "" {
		return nil, fmt.Errorf("no token provided")
	}
	//解析和验证 JWT 的
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.Cfg.Str.Key), nil

	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
