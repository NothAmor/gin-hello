package security

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/NothAmor/gin-hello/app/database"
	"github.com/NothAmor/gin-hello/app/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("ticketaware")

type (
	LoginInfo struct {
		Email    string
		Username string
		Password string
	}
	Claims struct {
		Username string
		Email    string
		jwt.StandardClaims
	}
)

func SignIn(login LoginInfo) (userInfo *models.Users, token string, ok bool) {
	// 初始化数据库连接
	db := database.InitDatabase()

	// 将密码加密为md5
	h := md5.New()
	h.Write([]byte(login.Password))
	password := hex.EncodeToString(h.Sum(nil))

	// 通过邮箱查询用户是否存在
	user := &models.Users{}
	err := db.Where("email = ?", login.Email).First(user).Error
	if err != nil {
		ok = false
	}

	// 如果查询有数据且传入的密码md5=数据库内的md5则验证成功
	if user.ID != 0 && password == user.Password {
		// 设置过期时间
		expTime := time.Now().Add(time.Hour * 24 * 7)
		claims := &Claims{
			Username: user.Username,
			Email:    user.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		stringToken, err := token.SignedString(jwtKey)
		if err != nil {
			panic(err)
		}
		ok = true

		return user, stringToken, ok
	} else {
		stringToken := ""
		ok = false
		return user, stringToken, ok
	}
}

func Validator(tokenString string) (username string, email string, ok bool) {
	claims := new(Claims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return
	}
	ok = true
	return claims.Username, claims.Email, ok
}
