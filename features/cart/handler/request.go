package handler

import "advancerentbook-api/features/cart"

type AddCartReq struct {
	UserID    uint    `json:"id_user" form:"id_user"`
	BookID    uint    `json:"book_id" form:"book_id"`
	RentPrice float64 `json:"rent_price" form:"rent_price"`
}

func ReqToCore(data interface{}) *cart.Core {
	res := cart.Core{}

	switch data.(type) {
	case AddCartReq:
		cnv := data.(AddCartReq)
		res.UserID = cnv.UserID
		res.BookID = cnv.BookID
		res.RentPrice = cnv.RentPrice
	default:
		return nil
	}
	return &res
}
