package data

import (
	book "advancerentbook-api/features/book/data"
	user "advancerentbook-api/features/user/data"

	"advancerentbook-api/features/cart"
)

type Cart struct {
	ID        uint
	UserID    uint
	BookID    uint
	RentPrice float64
	Owner     user.User `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Book      book.Book `gorm:"foreignkey:BookID;association_foreignkey:ID"`
}

func DataToCore(data Cart) cart.Core {
	return cart.Core{
		ID:        data.ID,
		UserID:    data.UserID,
		BookID:    data.BookID,
		RentPrice: data.RentPrice,
	}
}

func CoreToData(data cart.Core) Cart {
	return Cart{
		ID:        data.ID,
		UserID:    data.UserID,
		BookID:    data.BookID,
		RentPrice: data.RentPrice,
	}
}
