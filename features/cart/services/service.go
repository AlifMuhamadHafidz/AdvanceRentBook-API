package services

import (
	"advancerentbook-api/features/cart"
	"advancerentbook-api/helper"
	"errors"
	"log"
	"strings"
)

type cartUseCase struct {
	qry cart.CartData
}

// DeleteCart implements cart.CartService
func (*cartUseCase) DeleteCart(token interface{}) error {
	panic("unimplemented")
}

// UpdateCart implements cart.CartService
func (*cartUseCase) UpdateCart(token interface{}, cartID uint, updatedCart cart.Core) (cart.Core, error) {
	panic("unimplemented")
}

func New(cd cart.CartData) cart.CartService {
	return &cartUseCase{
		qry: cd,
	}
}

func (cuc *cartUseCase) AddCart(token interface{}, bookID uint, newCart cart.Core) (cart.Core, error) {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		log.Println("error extract token add cart")
		return cart.Core{}, errors.New("user not found")
	}
	res, err := cuc.qry.AddCart(uint(userId), bookID, newCart)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error add query in service: ", err.Error())
		return cart.Core{}, errors.New(msg)
	}
	return res, nil
}

func (cuc *cartUseCase) ShowCart(token interface{}) ([]cart.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := cuc.qry.ShowCart(uint(userID))
	if err != nil {
		log.Println("query error", err.Error())
		return []cart.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}
