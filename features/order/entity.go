package order

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID             uint
	BorrowerID     uint
	BorrowerName   string
	OwnerID        uint
	OwnerName      string
	TotalRentPrice float64
	OrderedAt      time.Time
	OrderStatus    string
	TransactionID  string
}

type OrderHandler interface {
	Add() echo.HandlerFunc
	GetOrderHistory() echo.HandlerFunc
	GetRentHistory() echo.HandlerFunc
	NotificationTransactionStatus() echo.HandlerFunc
	UpdateStatus() echo.HandlerFunc
}

type OrderService interface {
	Add(token interface{}, totalPrice float64) (Core, string, error)
	GetOrderHistory(token interface{}) ([]Core, error)
	GetRentHistory(token interface{}) ([]Core, error)
	NotificationTransactionStatus(transactionID string) error
	UpdateStatus(orderID uint, status string) error
}

type OrderData interface {
	Add(userId uint, totalRentPrice float64) (Core, string, error)
	GetOrderHistory(userID uint) ([]Core, error)
	GetRentHistory(userID uint) ([]Core, error)
	NotificationTransactionStatus(transactionID, transStatus string) error
	UpdateStatus(orderID uint, status string) error
}
