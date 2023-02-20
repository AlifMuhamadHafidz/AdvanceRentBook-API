package services

import (
	"advancerentbook-api/features/user"
	"advancerentbook-api/helper"
	"errors"
	"mime/multipart"
	"strings"
)

type userUseCase struct {
	qry user.UserData
}

// Deactivate implements user.UserService
func (*userUseCase) Deactivate(token interface{}) error {
	panic("unimplemented")
}

// Login implements user.UserService
func (*userUseCase) Login(username string, password string) (user.Core, error) {
	panic("unimplemented")
}

// Profile implements user.UserService
func (*userUseCase) Profile(token interface{}) (user.Core, error) {
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
