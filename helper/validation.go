package helper

import (
	"advancerentbook-api/features/user"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func TypeFile(test multipart.File) (string, error) {
	fileByte, _ := io.ReadAll(test)
	fileType := http.DetectContentType(fileByte)
	TipenamaFile := ""
	if fileType == "image/png" {
		TipenamaFile = ".png"
	} else {
		TipenamaFile = ".jpg"
	}
	if fileType == "image/png" || fileType == "image/jpeg" || fileType == "image/jpg" {
		return TipenamaFile, nil
	}
	return "", errors.New("file type not match")
}

type UserValidate struct {
	Name        string `validate:"required"`
	Username    string `validate:"required"`
	BirthOfDate string `validate:"required"`
	Email       string `validate:"required,email"`
	Phone       string `validate:"required,numeric"`
	Address     string `validate:"required"`
	Password    string `validate:"required,min=3,alphanum"`
}

func CoreToRegVal(data user.Core) UserValidate {
	return UserValidate{
		Name:        data.Name,
		BirthOfDate: data.BirthOfDate,
		Username:    data.Username,
		Email:       data.Email,
		Phone:       data.Phone,
		Address:     data.Address,
		Password:    data.Password,
	}
}
func RegistrationValidate(data user.Core) error {
	validate := validator.New()
	val := CoreToRegVal(data)
	if err := validate.Struct(val); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			vlderror := ""
			if e.Field() == "Password" && e.Value() != "" {
				vlderror = fmt.Sprintf("%s is not %s", e.Value(), e.Tag())
				return errors.New(vlderror)
			}
			if e.Value() == "" {
				vlderror = fmt.Sprintf("%s is %s", e.Field(), e.Tag())
				return errors.New(vlderror)
			} else {
				vlderror = fmt.Sprintf("%s is not %s", e.Value(), e.Tag())
				return errors.New(vlderror)
			}
		}
	}
	return nil
}
