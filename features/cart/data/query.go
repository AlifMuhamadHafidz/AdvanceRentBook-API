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

// ShowCart implements cart.CartData
func (*cartQuery) ShowCart(userID uint) ([]cart.Core, error) {
	panic("unimplemented")
}

// UpdateCart implements cart.CartData
func (*cartQuery) UpdateCart(userID uint, cartID uint, updatedCart cart.Core) (cart.Core, error) {
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
