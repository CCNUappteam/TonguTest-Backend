package service

import (
	MD5 "crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	_ "gopkg.in/gomail.v2"
	"math/rand"
	"time"
	User "tongue/model/user"
	"tongue/pkg/errno"
)

var storage = NewInnerStorage()

type Storage interface {
	setStorage(email, code string)
	clearStorage(email string)
	NewInnerStorage()
}

type innerStorage struct {
	storage  map[string]string
	codeChan chan string
}

func NewInnerStorage() *innerStorage {
	var one innerStorage
	one.storage = make(map[string]string)
	one.codeChan = make(chan string)

	go func() {
		for {
			select {
			case person := <-one.codeChan:
				delete(one.storage, person)
			}
		}
	}()

	return &one
}
func VerificationCode(email string) (string, error) {
	if err := User.IfExist(email, " "); err != nil {
		return "", errno.ServerErr(errno.ErrUserExisted, err.Error())
	}
	dialer := gomail.NewDialer(viper.GetString("email.host"), viper.GetInt("email.port"), viper.GetString("email.sender"), viper.GetString("email.secretKey"))
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err, code := newEmail(dialer, email)
	if err != nil {
		return "", errno.ServerErr(errno.ErrVerificationCode, err.Error())
	} else {
		storage.setStorage(email, code)
		fmt.Println(storage)
		go storage.clearStorage(email)
	}
	return code, nil
}

func newEmail(dialer *gomail.Dialer, emailOwner string) (error, string) {
	email := gomail.NewMessage()
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000000))
	message := fmt.Sprintf("欢迎使用TongueTest💗💗，您的验证码为 %s ，有效时间有2分钟。🥰", code)
	email.SetHeader("From", viper.GetString("email.sender"))
	email.SetHeader("To", emailOwner)
	email.SetHeader("Subject", "TongueTest认证")
	email.SetBody("text/html", message)

	return dialer.DialAndSend(email), code
}

func (i *innerStorage) setStorage(email, code string) {
	i.storage[email] = code
}

func (i *innerStorage) clearStorage(email string) {
	for {
		select {
		case <-time.After(time.Minute * 2):
			i.codeChan <- email
			return
		}
	}
}

func Register(email, name, password, age, gender, code string) error {
	// 是否重复注册
	if err := User.IfExist(email, name); err != nil {
		return errno.ServerErr(errno.ErrUserExisted, err.Error())
	}

	user := User.UserModel{
		Name:   name,
		Email:  email,
		Avatar: "https://cdn.jsdelivr.net/gh/Hyeonwuu/Image/user.png",
		Age:    age,
		Gender: gender,
	}
	if actual, exist := storage.storage[email]; !exist || code != actual {
		return errno.ErrValidation
	}
	md5 := MD5.New()
	md5.Write([]byte(password))
	user.HashPassword = hex.EncodeToString(md5.Sum(nil))
	if err := user.CreateUser(); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
