package data

import (
	"advancerentbook-api/features/book"

	"gorm.io/gorm"
)

type bookQuery struct {
	db *gorm.DB
}

// BookDetail implements book.BookData
func (*bookQuery) BookDetail(bookID uint) (book.Core, error) {
	panic("unimplemented")
}

// Delete implements book.BookData
func (*bookQuery) Delete(userID uint, bookID uint) error {
	panic("unimplemented")
}

// GetAllBook implements book.BookData
func (*bookQuery) GetAllBook() ([]book.Core, error) {
	panic("unimplemented")
}

// Update implements book.BookData
func (*bookQuery) Update(userID uint, updatedBook book.Core) (book.Core, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) book.BookData {
	return &bookQuery{
		db: db,
	}
}

func (bq *bookQuery) Add(userID uint, newBook book.Core) (book.Core, error) {
	cnv := CoreToData(newBook)
	cnv.UserID = uint(userID)
	err := bq.db.Create(&cnv).Error

	if err != nil {
		return book.Core{}, err
	}

	newBook.ID = cnv.ID
	newBook.UserID = cnv.UserID

	return newBook, nil
}
