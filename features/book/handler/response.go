package handler

import (
	"advancerentbook-api/features/book"
	"errors"
)

type Book struct {
	ID        uint    `json:"id"`
	Image     string  `json:"image"`
	Title     string  `json:"title"`
	Published int     `json:"published"`
	Author    string  `json:"author"`
	Publisher string  `json:"publisher"`
	RentPrice float64 `json:"rent_price"`
}

func BookResponse(data book.Core) Book {
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

func ConvertBookUpdateResponse(input book.Core) (interface{}, error) {
	ResponseFilter := book.Core{}
	ResponseFilter = input
	result := make(map[string]interface{})
	if ResponseFilter.ID != 0 {
		result["id"] = ResponseFilter.ID
	}
	if ResponseFilter.Image != "" {
		result["image"] = ResponseFilter.Image
	}
	if ResponseFilter.Title != "" {
		result["title"] = ResponseFilter.Title
	}
	if ResponseFilter.Published != 0 {
		result["published"] = ResponseFilter.Published
	}
	if ResponseFilter.Author != "" {
		result["author"] = ResponseFilter.Author
	}
	if ResponseFilter.Publisher != "" {
		result["publisher"] = ResponseFilter.Publisher
	}
	if ResponseFilter.RentPrice != 0 {
		result["rent_price"] = ResponseFilter.RentPrice
	}

	if len(result) <= 1 {
		return book.Core{}, errors.New("no data was change")
	}
	return result, nil
}
