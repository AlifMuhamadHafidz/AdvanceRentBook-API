package services

import (
	"advancerentbook-api/features/book"
	"advancerentbook-api/helper"
	"advancerentbook-api/mocks"
	"errors"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewBookData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.jpg")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	inputData := book.Core{
		ID:        0,
		Title:     "One Piece",
		Published: 1997,
		Author:    "Bang Oda",
		Publisher: "Shueisha",
		RentPrice: 2000,
	}
	resData := book.Core{
		ID:        1,
		Title:     "One Piece",
		Published: 1997,
		Author:    "Bang Oda",
		Publisher: "Shueisha",
		RentPrice: 2000,
	}

	t.Run("success post book", func(t *testing.T) {
		repo.On("Add", uint(1), inputData).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, *imageTrueCnv, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("invalid jwt", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		_, err := srv.Add(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})

	t.Run("cannot post book", func(t *testing.T) {
		repo.On("Add", uint(1), mock.Anything).Return(book.Core{}, errors.New("server error"))
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("invalid file validation", func(t *testing.T) {
		filePathFake := filepath.Join("..", "..", "..", "test.csv")
		headerFake, err := helper.UnitTestingUploadFileMock(filePathFake)
		if err != nil {
			log.Panic("from file header", err.Error())
		}
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, *headerFake, inputData)
		assert.ErrorContains(t, err, "type")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})

}

func TestGetAllBook(t *testing.T) {
	repo := mocks.NewBookData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.jpg")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	resData := []book.Core{{
		ID:        1,
		Title:     "One Piece",
		Published: 1997,
		Author:    "Bang Oda",
		Publisher: "Shueisha",
		Image:     imageTrueCnv.Filename,
		RentPrice: 2000,
	}}
	q := "naruto"

	t.Run("success get all book", func(t *testing.T) {
		repo.On("GetAllBook", q).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetAllBook(q)
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

	t.Run("book not found", func(t *testing.T) {
		repo.On("GetAllBook", q).Return([]book.Core{}, errors.New("book not found")).Once()
		srv := New(repo)
		res, err := srv.GetAllBook(q)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("GetAllBook", q).Return([]book.Core{}, errors.New("server problem")).Once()
		srv := New(repo)
		res, err := srv.GetAllBook(q)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})
}

func TestBookDetail(t *testing.T) {
	repo := mocks.NewBookData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.jpg")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	resData := book.Core{
		ID:        1,
		Title:     "One Piece",
		Published: 1997,
		Author:    "Bang Oda",
		Publisher: "Shueisha",
		Image:     imageTrueCnv.Filename,
		RentPrice: 2000,
	}

	t.Run("success get book detail", func(t *testing.T) {
		repo.On("BookDetail", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.BookDetail(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("BookDetail", uint(1)).Return(book.Core{}, errors.New("data not found")).Once()
		srv := New(repo)
		res, err := srv.BookDetail(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.NotEqual(t, 0, res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("server problem", func(t *testing.T) {
		repo.On("BookDetail", uint(1)).Return(book.Core{}, errors.New("server problem")).Once()
		srv := New(repo)
		res, err := srv.BookDetail(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.NotEqual(t, 0, res.ID)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewBookData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.jpg")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	inputData := book.Core{
		ID:        1,
		Title:     "One Piece",
		Published: 1997,
		Author:    "Bang Oda",
	}

	resData := book.Core{
		ID:        1,
		Title:     "One Piece",
		Published: 1997,
		Author:    "Bang Oda",
	}

	t.Run("success update book", func(t *testing.T) {
		repo.On("Update", uint(1), uint(1), inputData).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, uint(1), *imageTrueCnv, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Update", uint(1), uint(1), inputData).Return(book.Core{}, errors.New("data not found")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("invalid jwt", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		_, err := srv.Update(pToken, uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})

	t.Run("invalid file validation", func(t *testing.T) {
		filePathFake := filepath.Join("..", "..", "..", "test.csv")
		headerFake, err := helper.UnitTestingUploadFileMock(filePathFake)
		if err != nil {
			log.Panic("from file header", err.Error())
		}
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, uint(1), *headerFake, inputData)
		assert.ErrorContains(t, err, "validate")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("Update", uint(1), uint(1), inputData).Return(book.Core{}, errors.New("server problem")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.NotEqual(t, 0, res.ID)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewBookData(t)

	t.Run("success delete book", func(t *testing.T) {
		repo.On("Delete", uint(1), uint(1)).Return(nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)

	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Delete", uint(2), uint(2)).Return(errors.New("data not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(2)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 2)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})
}

func TestMyBook(t *testing.T) {
	repo := mocks.NewBookData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.jpg")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	resData := []book.Core{{
		ID:        1,
		Title:     "One Piece",
		Published: 1997,
		Author:    "Bang Oda",
		Publisher: "Shueisha",
		Image:     imageTrueCnv.Filename,
		RentPrice: 2000,
	}}

	t.Run("success get user book", func(t *testing.T) {
		repo.On("MyBook", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.MyBook(pToken)
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

	t.Run("book not found", func(t *testing.T) {
		repo.On("MyBook", uint(1)).Return([]book.Core{}, errors.New("book not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.MyBook(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("MyBook", uint(1)).Return([]book.Core{}, errors.New("server problem")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.MyBook(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})
}
