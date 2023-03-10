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

func (bq *bookQuery) BookDetail(bookID uint) (book.Core, error) {
	res := Book{}

	err := bq.db.Where("id = ?", bookID).First(&res).Error
	if err != nil {
		log.Println("data not found", err.Error())
		return book.Core{}, errors.New("data not found")
	}

	result := DataToCore(res)

	return result, nil
}

func (bq *bookQuery) Update(userID uint, bookID uint, updatedBook book.Core) (book.Core, error) {
	cnv := CoreToData(updatedBook)
	cnv.ID = uint(bookID)

	qry := bq.db.Where("id = ?", bookID).Updates(&cnv)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return book.Core{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update book query error", err.Error())
		return book.Core{}, errors.New("user not found")
	}
	return updatedBook, nil
}

func (bq *bookQuery) Delete(userID uint, bookID uint) error {
	getID := Book{}
	err := bq.db.Where("id = ? and user_id = ?", bookID, userID).First(&getID).Error
	if err != nil {
		log.Println("get book error : ", err.Error())
		return errors.New("failed to get book data")
	}

	qryDelete := bq.db.Delete(&Book{}, bookID)
	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("no rows affected")
		return errors.New("failed to delete book, data not found")
	}

	return nil
}

func (bq *bookQuery) MyBook(userID uint) ([]book.Core, error) {
	res := []Book{}
	if err := bq.db.Where("user_id = ?", userID).Order("created_at desc").Find(&res).Error; err != nil {
		log.Println("get employee approval query error : ", err.Error())
		return []book.Core{}, err
	}
	result := []book.Core{}
	for _, val := range res {
		result = append(result, DataToCore(val))
	}

	return result, nil
}
