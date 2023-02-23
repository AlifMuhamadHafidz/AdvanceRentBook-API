package cart

import "github.com/labstack/echo/v4"

// type Core struct {
// 	ID        uint    `json:"id"`
// 	UserID    uint    `json:"id_user"`
// 	OwnerID   uint    `json:"owner_id"`
// 	BookID    uint    `json:"book_id"`
// 	Image     string  `json:"image"`
// 	RentPrice float64 `json:"rent_price"`
// 	BookName  string  `json:"book_name"`
// 	OwnerName string  `json:"owner_name"`
// }

type Core struct {
	ID        uint   `json:"id"`
	BookName  string `json:"book_name"`
	Image     string `json:"image"`
	Owner     string `json:"owner"`
	RentPrice int    `json:"rent_price"`
	User      User   `json:"user"`
}

type User struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type CartHandler interface {
	AddCart() echo.HandlerFunc
	ShowCart() echo.HandlerFunc
	DeleteCart() echo.HandlerFunc
}

type CartService interface {
	AddCart(token interface{}, bookID uint, newCart Core) (Core, error)
	ShowCart(token interface{}) ([]Core, error)
	DeleteCart(token interface{}) error
}

type CartData interface {
	AddCart(userID uint, bookID uint, newCart Core) (Core, error)
	ShowCart(userID uint) ([]Core, error)
	DeleteCart(userID uint) error
}
