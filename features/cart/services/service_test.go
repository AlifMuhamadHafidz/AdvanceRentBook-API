package services

import (
	"advancerentbook-api/features/cart"
	"advancerentbook-api/helper"
	"advancerentbook-api/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAddCart(t *testing.T) {
	repo := mocks.NewCartData(t)
	inputData := cart.Core{ID: 1, BookName: "Naruto", Owner: "alif", RentPrice: 5000}
	resData := cart.Core{ID: 1, BookName: "Naruto", Owner: "alif", RentPrice: 5000}

	t.Run("success add cart", func(t *testing.T) {
		repo.On("AddCart", uint(1), uint(1), inputData).Return(resData, nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.AddCart(pToken, uint(1), inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.BookName, res.BookName)
		repo.AssertExpectations(t)
	})

	t.Run("invalid jwt", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		_, err := srv.AddCart(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})

	t.Run("cart not found", func(t *testing.T) {
		repo.On("AddCart", uint(1), uint(1), inputData).Return(cart.Core{}, errors.New("cart not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.AddCart(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("AddCart", uint(1), uint(1), inputData).Return(cart.Core{}, errors.New("server problem")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.AddCart(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}

func TestShowCart(t *testing.T) {
	repo := mocks.NewCartData(t)
	resData := []cart.Core{{ID: 1, BookName: "Naruto", Owner: "alif", RentPrice: 5000}}

	t.Run("success show cart", func(t *testing.T) {
		repo.On("ShowCart", uint(1)).Return(resData, nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.ShowCart(pToken)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("ShowCart", uint(1)).Return([]cart.Core{}, errors.New("server problem")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.ShowCart(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})
}

func TestDeleteCart(t *testing.T) {
	repo := mocks.NewCartData(t)

	t.Run("success delete data", func(t *testing.T) {
		repo.On("DeleteCart", uint(1), uint(1)).Return(nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.DeleteCart(pToken, uint(1))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("problem with server", func(t *testing.T) {
		repo.On("DeleteCart", uint(1), uint(1)).Return(errors.New("problem with server")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.DeleteCart(pToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}
