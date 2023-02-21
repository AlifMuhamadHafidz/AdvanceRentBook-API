package data

import (
	"advancerentbook-api/features/book"
	user "advancerentbook-api/features/user/data"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	UserID    uint
	User      user.User
	Title     string
	Image     string
	Published int
	Author    string
	Publisher string
	RentPrice float64
}

func DataToCore(data Book) book.Core {
	return book.Core{
		ID:        data.ID,
		UserID:    data.UserID,
		Title:     data.Title,
		Image:     data.Image,
		Published: data.Published,
		Author:    data.Author,
		Publisher: data.Publisher,
		RentPrice: data.RentPrice,
	}
}

func CoreToData(data book.Core) Book {
	return Book{
		Model:     gorm.Model{ID: data.ID},
		UserID:    data.UserID,
		Title:     data.Title,
		Image:     data.Image,
		Published: data.Published,
		Author:    data.Author,
		Publisher: data.Publisher,
		RentPrice: data.RentPrice,
	}
}
