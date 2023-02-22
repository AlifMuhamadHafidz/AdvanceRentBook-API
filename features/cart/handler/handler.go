package handler

import (
	"advancerentbook-api/features/cart"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type cartControll struct {
	srv cart.CartService
}

// DeleteCart implements cart.CartHandler
func (*cartControll) DeleteCart() echo.HandlerFunc {
	panic("unimplemented")
}

// ShowCart implements cart.CartHandler
func (*cartControll) ShowCart() echo.HandlerFunc {
	panic("unimplemented")
}

// UpdateCart implements cart.CartHandler
func (*cartControll) UpdateCart() echo.HandlerFunc {
	panic("unimplemented")
}

func New(srv cart.CartService) cart.CartHandler {
	return &cartControll{
		srv: srv,
	}
}

func (cc *cartControll) AddCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := AddCartReq{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}

		res, err := cc.srv.AddCart(token, input.BookID, *ReqToCore(input))
		if err != nil {
			log.Println("error running add book service: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "server problem"})
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success add book to cart",
		})
	}
}
