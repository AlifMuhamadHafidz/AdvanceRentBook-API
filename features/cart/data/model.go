package data

import (
	"advancerentbook-api/features/cart"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint
	BookID    uint
	RentPrice int
	User      User
}

type Book struct {
	gorm.Model
	Title     string
	Image     string
	RentPrice int
	UserID    uint
}

type User struct {
	gorm.Model
	Name      string
	Email     string
	Phone     string
	Address   string
	UserImage string
}

func DataToCore(data Cart) cart.Core {
	return cart.Core{
		ID:        data.ID,
		RentPrice: data.RentPrice,
		User: cart.User{
			ID:      data.User.ID,
			Name:    data.User.Name,
			Email:   data.User.Email,
			Phone:   data.User.Phone,
			Address: data.User.Address,
		},
	}
}

func CoreToData(core cart.Core) Cart {
	return Cart{
		Model:     gorm.Model{ID: core.ID},
		RentPrice: core.RentPrice,
	}
}
