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

func (uq *userQuery) Profile(userID uint) (user.Core, error) {
	res := User{}
	err := uq.db.Where("id = ?", userID).First(&res).Error
	if err != nil {
		log.Println("query err", err.Error())
		return user.Core{}, errors.New("account not found")
	}
	return ToCore(res), nil
}

func (uq *userQuery) Update(userID uint, updateData user.Core) (user.Core, error) {
	if updateData.Email != "" {
		dupEmail := User{}
		err := uq.db.Where("email = ?", updateData.Email).First(&dupEmail).Error
		if err == nil {
			log.Println("duplicated")
			return user.Core{}, errors.New("email duplicated")
		}
	}

	cnv := CoreToData(updateData)
	qry := uq.db.Model(&User{}).Where("id = ?", userID).Updates(&cnv)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return user.Core{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return user.Core{}, errors.New("user not found")
	}
	result := ToCore(cnv)
	result.ID = userID
	return result, nil
}

func (uq *userQuery) Deactivate(userID uint) error {
	getID := User{}
	err := uq.db.Where("id = ?", userID).First(&getID).Error
	if err != nil {
		log.Println("get user error : ", err.Error())
		return errors.New("failed to get user data")
	}

	if getID.ID != userID {
		log.Println("unauthorized request")
		return errors.New("unauthorized request")
	}
	qryDelete := uq.db.Delete(&User{}, userID)
	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("failed to delete user content, data not found")
	}

	return nil
}
