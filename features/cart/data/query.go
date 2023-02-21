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
	cnvC := CoreToData(newCart)
	cnvC.UserID = userID
	cnvC.BookID = bookID

	err := cq.db.Create(&cnvC).Error
	if err != nil {
		log.Println("add cart query error: ", err.Error())
		return cart.Core{}, errors.New("server problem")
	}
	return DataToCore(cnvC), nil
}
