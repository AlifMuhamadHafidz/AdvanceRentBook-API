package data

import (
	book "advancerentbook-api/features/book/data"
	"advancerentbook-api/features/order"
	user "advancerentbook-api/features/user/data"
	"time"
)

type Order struct {
	ID             uint `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	BorrowerID     uint
	OwnerID        uint
	TotalRentPrice float64
	OrderedAt      time.Time
	OrderStatus    string
	TransactionID  string
	PaymentURL     string

	Owner    user.User `gorm:"foreignkey:OwnerID;association_foreignkey:ID"`
	Borrower user.User `gorm:"foreignkey:BorrowerID;association_foreignkey:ID"`
}

type OrderBook struct {
	ID        uint      `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	OrderID   uint      `json:"order_id"`
	BookID    uint      `json:"book_id"`
	RentPrice float64   `json:"rent_price"`
	Order     Order     `gorm:"foreignkey:OrderID;association_foreignkey:ID"`
	Book      book.Book `gorm:"foreignkey:BookID;association_foreignkey:ID"`
}

func DataToCore(data Order) order.Core {
	return order.Core{
		ID:             data.ID,
		BorrowerID:     data.BorrowerID,
		OwnerID:        data.OwnerID,
		TotalRentPrice: data.TotalRentPrice,
		OrderedAt:      data.OrderedAt,
		OrderStatus:    data.OrderStatus,
		TransactionID:  data.TransactionID,
	}
}

func CoreToData(data order.Core) Order {
	return Order{
		ID:             data.ID,
		BorrowerID:     data.BorrowerID,
		OwnerID:        data.OwnerID,
		TotalRentPrice: data.TotalRentPrice,
		OrderedAt:      data.OrderedAt,
		OrderStatus:    data.OrderStatus,
		TransactionID:  data.TransactionID,
	}
}
