package cart

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"id_user"`
	OwnerID   uint    `json:"owner_id"`
	BookID    uint    `json:"book_id"`
	Image     string  `json:"image"`
	RentPrice float64 `json:"rent_price"`
	BookName  string  `json:"book_name"`
	OwnerName string  `json:"owner_name"`
}

type CartHandler interface {
	AddCart() echo.HandlerFunc
	ShowCart() echo.HandlerFunc
	UpdateCart() echo.HandlerFunc
	DeleteCart() echo.HandlerFunc
}

type CartService interface {
	AddCart(token interface{}, bookID uint, newCart Core) (Core, error)
	ShowCart(token interface{}) ([]Core, error)
	UpdateCart(token interface{}, cartID uint, updatedCart Core) (Core, error)
	DeleteCart(token interface{}) error
}

type CartData interface {
	AddCart(userID uint, bookID uint, newCart Core) (Core, error)
	ShowCart(userID uint) ([]Core, error)
	UpdateCart(userID uint, cartID uint, updatedCart Core) (Core, error)
	DeleteCart(userID uint) error
}
