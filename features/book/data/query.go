package data

import (
	"advancerentbook-api/features/book"
	"errors"
	"log"

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

func (bq *bookQuery) GetAllBook(quote string) ([]book.Core, error) {
	res := []Book{}

	err := bq.db.Where("title LIKE ?", "%"+quote+"%").Find(&res).Error
	if err != nil {
		log.Println("no data processed", err.Error())
		return []book.Core{}, errors.New("no book found")
	}
	result := []book.Core{}
	for _, val := range res {
		result = append(result, DataToCore(val))
		if quote == "" {
			err := bq.db.Find(&res).Error
			if err != nil {
				log.Println("data not found", err.Error())
				return []book.Core{}, errors.New("data not found")
			}
		}
	}
	return result, nil
}
