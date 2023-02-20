package services

import (
	"advancerentbook-api/config"
	"advancerentbook-api/features/user"
	"advancerentbook-api/helper"
	"errors"
	"log"
	"mime/multipart"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type userUseCase struct {
	qry user.UserData
}

// Deactivate implements user.UserService
func (*userUseCase) Deactivate(token interface{}) error {
	panic("unimplemented")
}

// Update implements user.UserService
func (*userUseCase) Update(token interface{}, fileData multipart.FileHeader, updateData user.Core) (user.Core, error) {
	panic("unimplemented")
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{
		qry: ud,
	}
}

func (uuc *userUseCase) Register(newUser user.Core) (user.Core, error) {
	if len(newUser.Password) != 0 {
		//validation
		err := helper.RegistrationValidate(newUser)
		if err != nil {
			return user.Core{}, errors.New("validate: " + err.Error())
		}
	}
	hashed := helper.GeneratePassword(newUser.Password)
	newUser.Password = string(hashed)

	res, err := uuc.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "email already registered"
		} else {
			msg = "server error"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Login(username, password string) (string, user.Core, error) {
	res, err := uuc.qry.Login(username)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "there is a problem with the server"
		}
		return "", user.Core{}, errors.New(msg)
	}

	if err := helper.ComparePassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", user.Core{}, errors.New("password not matched")
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = res.ID
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	useToken, _ := token.SignedString([]byte(config.JWTKey))

	return useToken, res, nil
}

func (uuc *userUseCase) Profile(token interface{}) (user.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := uuc.qry.Profile(uint(userID))
	if err != nil {
		log.Println("data not found")
		return user.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}
