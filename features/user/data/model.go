package data

import (
	"advancerentbook-api/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	ProfilePicture string
	Username       string
	BirthOfDate    string
	Email          string
	Password       string
	Phone          string
	Address        string
}

func ToCore(data User) user.Core {
	return user.Core{
		ID:             data.ID,
		ProfilePicture: data.ProfilePicture,
		Name:           data.Name,
		BirthOfDate:    data.BirthOfDate,
		Username:       data.Username,
		Email:          data.Email,
		Phone:          data.Phone,
		Address:        data.Address,
		Password:       data.Password,
	}
}

func CoreToData(data user.Core) User {
	return User{
		Model:          gorm.Model{ID: data.ID},
		ProfilePicture: data.ProfilePicture,
		Name:           data.Name,
		BirthOfDate:    data.BirthOfDate,
		Username:       data.Username,
		Email:          data.Email,
		Phone:          data.Phone,
		Address:        data.Address,
		Password:       data.Password,
	}
}
