package handler

import (
	"advancerentbook-api/features/book"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type bookControll struct {
	srv book.BookService
}

// BookDetail implements book.BookHandler
func (*bookControll) BookDetail() echo.HandlerFunc {
	panic("unimplemented")
}

// Delete implements book.BookHandler
func (*bookControll) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// GetAllBook implements book.BookHandler
func (*bookControll) GetAllBook() echo.HandlerFunc {
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
