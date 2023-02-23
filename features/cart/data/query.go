package data

import (
	"advancerentbook-api/features/cart"
	"errors"
	"log"

	"gorm.io/gorm"
)

type cartQuery struct {
	db *gorm.DB
}

// DeleteCart implements cart.CartData
func (*cartQuery) DeleteCart(userID uint) error {
	panic("unimplemented")
}

func New(db *gorm.DB) cart.CartData {
	return &cartQuery{
		db: db,
	}
}

func (cq *cartQuery) AddCart(userID uint, bookID uint, newCart cart.Core) (cart.Core, error) {
	book := Book{}
	err := cq.db.Where("id=?", bookID).First(&book).Error
	if err != nil {
		log.Println("query error", err.Error())
		return cart.Core{}, errors.New("server error")
	}
	cnv := CoreToData(newCart)
	cnv.BookID = book.ID
	cnv.UserID = userID
	cnv.RentPrice = book.RentPrice
	err = cq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return cart.Core{}, errors.New("server error")
	}
	result := DataToCore(cnv)
	return result, nil
}

func (cq *cartQuery) ShowCart(userID uint) ([]cart.Core, error) {
	res := []Cart{}
	err := cq.db.Where("user_id = ?", userID).Find(&res).Error

	if err != nil {
		log.Println("query error", err.Error())
		return []cart.Core{}, errors.New("server error")
	}

	result := []cart.Core{}
	for i := 0; i < len(res); i++ {
		result = append(result, DataToCore(res[i]))
		// cari data user berdasarkan cart user_id

		book := Book{}
		err = cq.db.Where("id = ?", res[i].BookID).First(&book).Error
		if err != nil {
			log.Println("query error", err.Error())
			return []cart.Core{}, errors.New("server error")
		}
		user := User{}
		err = cq.db.Where("id = ?", book.UserID).First(&user).Error
		if err != nil {
			log.Println("query error", err.Error())
			return []cart.Core{}, errors.New("server error")
		}
		// cari data product berdasarkan cart product_id
		result[i].Owner = user.Name
		result[i].BookName = book.Title
		result[i].Image = book.Image
	}
	return result, nil

}
