package services

import (
	"advancerentbook-api/features/book"
	"advancerentbook-api/helper"
	"errors"
	"log"
	"mime/multipart"
)

type bookUseCase struct {
	qry book.BookData
}

// BookDetail implements book.BookService
func (*bookUseCase) BookDetail(bookID uint) (book.Core, error) {
	panic("unimplemented")
}

// Delete implements book.BookService
func (*bookUseCase) Delete(token interface{}, bookID uint) error {
	panic("unimplemented")
}

// GetAllBook implements book.BookService
func (*bookUseCase) GetAllBook() ([]book.Core, error) {
	panic("unimplemented")
}

// Update implements book.BookService
func (*bookUseCase) Update(token interface{}, fileData multipart.FileHeader, updatedBook book.Core) (book.Core, error) {
	panic("unimplemented")
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
