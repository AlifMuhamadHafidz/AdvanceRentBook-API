package handler

import (
	"advancerentbook-api/features/book"
	"advancerentbook-api/helper"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type bookControll struct {
	srv book.BookService
}

// Delete implements book.BookHandler
func (*bookControll) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements book.BookHandler
func (*bookControll) Update() echo.HandlerFunc {
	panic("unimplemented")
}

func New(srv book.BookService) book.BookHandler {
	return &bookControll{
		srv: srv,
	}
}

func (bc *bookControll) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := PostBookRequest{}

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}
		//proses cek apakah user input foto ?
		checkFile, _, _ := c.Request().FormFile("image")
		if checkFile != nil {
			formHeader, err := c.FormFile("image")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader = *formHeader
		}

		res, err := bc.srv.Add(token, input.FileHeader, *ReqToCore(input))

		if err != nil {
			if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "size") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "file size max 500kb"})
			} else if strings.Contains(err.Error(), "validate") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "not registered") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "unable to process data"})
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success post a book",
		})

	}
}

// GetAllBook implements book.BookHandler
func (bc *bookControll) GetAllBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		resBook := []Book{}
		quotes := c.QueryParam("q")

		res, err := bc.srv.GetAllBook(quotes)
		if err != nil {
			if len(res) == 0 {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "book not found"})
			}
		}
		for _, val := range res {
			resBook = append(resBook, BookResponse(val))
		}
		if quotes == "" {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"data":    resBook,
				"message": "success show all book",
			})
		}

		result := []Book{}
		for i := 0; i < len(res); i++ {
			result = append(result, BookResponse(res[i]))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "searching success",
		})

	}
}

func (bc *bookControll) BookDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")

		bookID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "invalid input")
		}

		res, err := bc.srv.BookDetail(uint(bookID))

		if err != nil {
			if strings.Contains(err.Error(), "data") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "data not found",
				})
			}
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "success get book details", BookResponse(res)))
	}
}
