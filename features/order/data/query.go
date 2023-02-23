package data

import (
	cart "advancerentbook-api/features/cart/data"
	"advancerentbook-api/features/order"
	"log"
	"time"

	"gorm.io/gorm"
)

type orderQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) order.OrderData {
	return &orderQuery{
		db: db,
	}
}

func (oq *orderQuery) Add(userID uint, totalRentPrice float64) (order.Core, string, error) {
	tx := oq.db.Begin()

	// mengambil cart user
	userCart := []cart.Cart{}
	if err := tx.Where("user_id = ?", userID).Find(&userCart).Error; err != nil {
		tx.Rollback()
		log.Println("error retrieve user cart: ", err.Error())
		return order.Core{}, "", err
	}

	// membuat order
	orderinput := Order{
		BorrowerID:     userID,
		OrderStatus:    "waiting for payment",
		OrderedAt:      time.Now(),
		TotalRentPrice: totalRentPrice,
	}
}
