package handler

import "advancerentbook-api/features/cart"

type AddCartReq struct {
	BookID    uint `json:"book_id" form:"book_id"`
	RentPrice int  `json:"rent_price" form:"rent_price"`
}

func ReqToCore(data interface{}) *cart.Core {
	res := cart.Core{}

	switch data.(type) {
	case AddCartReq:
		cnv := data.(AddCartReq)
		res.RentPrice = cnv.RentPrice
	default:
		return nil
	}
	return &res
}
