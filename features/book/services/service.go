package services

import (
	"advancerentbook-api/features/book"
	"advancerentbook-api/helper"
	"errors"
	"log"
	"mime/multipart"
	"strings"
)

type bookUseCase struct {
	qry book.BookData
}

func New(bd book.BookData) book.BookService {
	return &bookUseCase{
		qry: bd,
	}
}

func (buc *bookUseCase) Add(token interface{}, fileData multipart.FileHeader, newBook book.Core) (book.Core, error) {
	userID := helper.ExtractToken(token)

	if userID <= 0 {
		return book.Core{}, errors.New("user not found")
	}

	url, err := helper.GetUrlImagesFromAWS(fileData)
	if err != nil {
		return book.Core{}, errors.New("validate: " + err.Error())
	}
	newBook.Image = url

	res, err := buc.qry.Add(uint(userID), newBook)

	if err != nil {
		log.Println("cannot post book", err.Error())
		return book.Core{}, errors.New("server error")
	}

	return res, nil
}

func (buc *bookUseCase) GetAllBook(quote string) ([]book.Core, error) {
	res, err := buc.qry.GetAllBook(quote)

	if err != nil {
		if strings.Contains(err.Error(), "book") {
			return []book.Core{}, errors.New("book not found")
		} else {
			return []book.Core{}, errors.New("data not found")
		}
	}
	return res, nil
}

func (buc *bookUseCase) BookDetail(bookID uint) (book.Core, error) {
	res, err := buc.qry.BookDetail(bookID)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return book.Core{}, errors.New("data not found")
		} else {
			return book.Core{}, errors.New("internal server error")
		}
	}

	return res, nil
}

func (buc *bookUseCase) Update(token interface{}, bookID uint, fileData multipart.FileHeader, updatedBook book.Core) (book.Core, error) {
	userID := helper.ExtractToken(token)

	if userID <= 0 {
		return book.Core{}, errors.New("user not found")
	}

	url, err := helper.GetUrlImagesFromAWS(fileData)
	if err != nil {
		return book.Core{}, errors.New("validate: " + err.Error())
	}
	updatedBook.Image = url
	res, err := buc.qry.Update(uint(userID), bookID, updatedBook)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return book.Core{}, errors.New("data not found")
		} else {
			return book.Core{}, errors.New("internal server error")
		}
	}

	return res, nil
}

func (buc *bookUseCase) Delete(token interface{}, bookID uint) error {
	id := helper.ExtractToken(token)

	err := buc.qry.Delete(uint(id), bookID)

	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("data not found")
	}

	return nil
}
