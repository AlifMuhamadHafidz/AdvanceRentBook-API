package handler

import (
	"advancerentbook-api/features/cart"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type cartControll struct {
	srv cart.CartService
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

func (cc *cartControll) ShowCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		res, err := cc.srv.ShowCart(token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    GetCartResp(res),
			"message": "success show all book in cart",
		})
	}
}

func (cc *cartControll) DeleteCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		paramID := c.Param("id")
		cartID, _ := strconv.Atoi(paramID)
		err := cc.srv.DeleteCart(token, uint(cartID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete cart",
		})
	}
}
