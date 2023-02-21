package handler

import (
	"advancerentbook-api/features/book"
	"mime/multipart"
)

type PostBookRequest struct {
	Title      string  `json:"title" form:"title"`
	Published  int     `json:"published" form:"published"`
	Author     string  `json:"author" form:"author"`
	Publisher  string  `json:"publisher" form:"publisher"`
	RentPrice  float64 `json:"rent_price" form:"rent_price"`
	FileHeader multipart.FileHeader
}

func ReqToCore(data interface{}) *book.Core {
	res := book.Core{}

	switch data.(type) {
	case PostBookRequest:
		cnv := data.(PostBookRequest)
		res.Title = cnv.Title
		res.Published = cnv.Published
		res.Author = cnv.Author
		res.Publisher = cnv.Publisher
		res.RentPrice = cnv.RentPrice
	default:
		return nil
	}

	return &res
}
