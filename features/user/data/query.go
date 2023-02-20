package data

import (
	"advancerentbook-api/features/user"
	"errors"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// Deactivate implements user.UserData
func (*userQuery) Deactivate(userID uint) error {
	panic("unimplemented")
}

// Profile implements user.UserData
func (*userQuery) Profile(userID uint) (user.Core, error) {
	panic("unimplemented")
}

// Update implements user.UserData
func (*userQuery) Update(userID uint, updateData user.Core) (user.Core, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {
	dupEmail := CoreToData(newUser)
	err := uq.db.Where("email = ?", newUser.Email).First(&dupEmail).Error
	if err == nil {
		log.Println("duplicated")
		return user.Core{}, errors.New("email duplicated")
	}

	newUser.ProfilePicture = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png"

	cnv := CoreToData(newUser)
	err = uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return user.Core{}, errors.New("server error")
	}

	newUser.ID = cnv.ID
	return newUser, nil
}

func (uq *userQuery) Login(username string) (user.Core, error) {
	if username == "" {
		log.Println("data empty, query error")
		return user.Core{}, errors.New("username not allowed empty")
	}
	res := User{}
	if err := uq.db.Where("username = ?", username).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return user.Core{}, errors.New("data not found")
	}

	return ToCore(res), nil
}
