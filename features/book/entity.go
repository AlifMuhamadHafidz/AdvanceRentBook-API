package book

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"user_id"`
	Title     string  `json:"title"`
	Image     string  `json:"image"`
	Published int     `json:"published"`
	Author    string  `json:"author"`
	Publisher string  `json:"publisher"`
	RentPrice float64 `json:"rent_price"`
}

type BookHandler interface {
	Add() echo.HandlerFunc
	GetAllBook() echo.HandlerFunc
	BookDetail() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type BookService interface {
	Add(token interface{}, fileData multipart.FileHeader, newBook Core) (Core, error)
	GetAllBook(quote string) ([]Core, error)
	BookDetail(bookID uint) (Core, error)
	Update(token interface{}, bookID uint, fileData multipart.FileHeader, updatedBook Core) (Core, error)
	Delete(token interface{}, bookID uint) error
}

type BookData interface {
	Add(userID uint, newBook Core) (Core, error)
	GetAllBook(quote string) ([]Core, error)
	BookDetail(bookID uint) (Core, error)
	Update(userID uint, bookID uint, updatedBook Core) (Core, error)
	Delete(userID uint, bookID uint) error
}
