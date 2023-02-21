package handler

import "advancerentbook-api/features/book"

type Book struct {
	ID        uint    `json:"id"`
	Image     string  `json:"image"`
	Title     string  `json:"title"`
	Published int     `json:"published"`
	Author    string  `json:"author"`
	Publisher string  `json:"publisher"`
	RentPrice float64 `json:"rent_price"`
}

func SearchResponse(data book.Core) Book {
	return Book{
		ID:        data.ID,
		Image:     data.Image,
		Title:     data.Title,
		Published: data.Published,
		Author:    data.Author,
		Publisher: data.Publisher,
		RentPrice: data.RentPrice,
	}
}

type Search struct {
	ID        uint    `json:"id"`
	Image     string  `json:"image"`
	Title     string  `json:"title"`
	Published int     `json:"published"`
	Author    string  `json:"author"`
	Publisher string  `json:"publisher"`
	RentPrice float64 `json:"rent_price"`
}

func SearchBookResponse(data book.Core) Search {
	return Search{
		ID:        data.ID,
		Image:     data.Image,
		Title:     data.Title,
		Published: data.Published,
		Author:    data.Author,
		Publisher: data.Publisher,
		RentPrice: data.RentPrice,
	}
}
