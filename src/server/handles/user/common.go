package user

import (
	"io"
	"fmt"
	"crypto/md5"
	"strconv"
	"math/rand"
	"time"
	"server/net/responese"
	"github.com/gin-gonic/gin"
)

type LoginData struct{
	User string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User string `json:"username"`
	Token string `json:"token"`
	AccessLevel int `json:"accesslevel"`
	Timeout int64 		`json:"timeout"`
}

func Login(c *gin.Context) {
	loginData := LoginData{}
	err := c.BindJSON(&loginData)
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	if admin.Password != loginData.Password || admin.Username != loginData.User {
		responese.Error(c, "username or password is invalidate", nil)
		return
	}
	responese.Success(c, "login successed", &LoginResponse{admin.Username, genToken(admin), admin.AccessLevel, time.Now().Unix() + 3600})
}

func genToken(u User)string{
	nano := time.Now().UnixNano()
    rand.Seed(nano)
    rndNum := rand.Int63()
	sessionID := Md5(u.Username + Md5(strconv.FormatInt(nano, 10))+Md5(strconv.FormatInt(rndNum, 10)))
	return sessionID
}

func Md5(text string) string {
    hashMd5 := md5.New()
    io.WriteString(hashMd5, text)
    return fmt.Sprintf("%x", hashMd5.Sum(nil))
}