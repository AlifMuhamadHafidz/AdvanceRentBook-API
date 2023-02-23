package handler

import "advancerentbook-api/features/cart"

type CartResponse struct {
	ID        uint   `json:"id"`
	BookName  string `json:"book_name"`
	Owner     string `json:"owner"`
	RentPrice int    `json:"rent_price"`
	Image     string `json:"image"`
}

func CToResponse(data cart.Core) CartResponse {
	return CartResponse{
		ID:        data.ID,
		BookName:  data.BookName,
		Owner:     data.Owner,
		RentPrice: data.RentPrice,
		Image:     data.Image,
	}
}

func GetCartResp(data []cart.Core) []CartResponse {
	res := []CartResponse{}
	for _, v := range data {
		res = append(res, CToResponse(v))
	}
	return res
}
