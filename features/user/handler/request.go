package handler

import (
	"advancerentbook-api/features/user"
	"mime/multipart"
)

type RegisterRequest struct {
	Name        string `json:"name" form:"name"`
	Username    string `json:"username" form:"username"`
	BirthOfDate string `json:"birth_of_date" form:"birth_of_date"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Phone       string `json:"phone" form:"phone"`
	Address     string `json:"address" form:"address"`
	FileHeader  multipart.FileHeader
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func ReqToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Name = cnv.Name
		res.Username = cnv.Username
		res.BirthOfDate = cnv.BirthOfDate
		res.Email = cnv.Email
		res.Password = cnv.Password
		res.Phone = cnv.Phone
		res.Address = cnv.Address
	case LoginRequest:
		cnv := data.(LoginRequest)
		res.Username = cnv.Username
		res.Password = cnv.Password
	default:
		return nil
	}

	return &res
}
