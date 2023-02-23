package services

import (
	"advancerentbook-api/features/user"
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

func TestRegister(t *testing.T) {
	repo := mocks.NewUserData(t)
	inputData := user.Core{Name: "Alif", Username: "alif23", Email: "alif@gmail.com", Phone: "081234"}
	resData := user.Core{ID: uint(1), Name: "Alif", Username: "alif23", Email: "alif@gmail.com", Phone: "081234"}

	t.Run("success create account", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.Register(inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Phone, res.Phone)
		repo.AssertExpectations(t)
	})

	t.Run("invalid validation", func(t *testing.T) {
		wrongInput := user.Core{Name: "Alif%&*", Password: "123a#$%"}
		// repo.On("Register", uint(1), mock.Anything).Return(user.Core{}, errors.New("email duplicated")).Once()
		srv := New(repo)
		res, err := srv.Register(wrongInput)
		assert.ErrorContains(t, err, "validate")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})

	t.Run("Duplicated", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(user.Core{}, errors.New("email duplicated")).Once()
		srv := New(repo)
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "email already registered")
		repo.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(user.Core{}, errors.New("server error")).Once()
		srv := New(repo)
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server error")
		repo.AssertExpectations(t)
	})

}

func TestLogin(t *testing.T) {
	repo := mocks.NewUserData(t)
	inputUsername := "alif23"
	passwordHashed := helper.GeneratePassword("123")
	resData := user.Core{ID: uint(1), Name: "Alif", Email: "alif@gmail.com", Password: passwordHashed}

	t.Run("login success", func(t *testing.T) {
		repo.On("Login", inputUsername).Return(resData, nil).Once()
		srv := New(repo)
		token, res, err := srv.Login(inputUsername, "123")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.Name, res.Name)
		repo.AssertExpectations(t)
	})

	t.Run("account not found", func(t *testing.T) {
		repo.On("Login", inputUsername).Return(user.Core{}, errors.New("data not found")).Once()
		srv := New(repo)
		token, res, err := srv.Login(inputUsername, "123")
		assert.NotNil(t, token)
		assert.ErrorContains(t, err, "not")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("password not matched", func(t *testing.T) {
		inputUsername := "lip23"
		repo.On("Login", inputUsername).Return(resData, nil).Once()
		srv := New(repo)
		_, res, err := srv.Login(inputUsername, "342")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "password")
		assert.Empty(t, nil)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		inputUsername := "alif23"
		repo.On("Login", inputUsername).Return(user.Core{}, errors.New("There is a problem with the server")).Once()

		srv := New(repo)
		token, res, err := srv.Login(inputUsername, "342")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Empty(t, token)
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestProfile(t *testing.T) {
	repo := mocks.NewUserData(t)
	resData := user.Core{ID: uint(1), Name: "Alif", Username: "alif23", Email: "alif@gmail.com", Phone: "081234"}

	t.Run("success show profile", func(t *testing.T) {
		repo.On("Profile", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("account not found", func(t *testing.T) {
		repo.On("Profile", uint(1)).Return(user.Core{}, errors.New("query error, problem with server")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewUserData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.jpg")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	inputData := user.Core{ID: 1, Name: "Alif", Phone: "08123"}
	resData := user.Core{ID: 1, Name: "Alif", Phone: "08123"}

	t.Run("success updating account", func(t *testing.T) {
		repo.On("Update", uint(1), mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, *imageTrueCnv, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("fail updating account", func(t *testing.T) {
		repo.On("Update", uint(1), mock.Anything).Return(user.Core{}, errors.New("user not found")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not registered")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})
	t.Run("email duplicated", func(t *testing.T) {
		repo.On("Update", uint(1), mock.Anything).Return(user.Core{}, errors.New("email duplicated")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "email duplicated")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})
	t.Run("access denied", func(t *testing.T) {
		repo.On("Update", uint(1), mock.Anything).Return(user.Core{}, errors.New("access denied")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "access denied")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("invalid file validation", func(t *testing.T) {
		filePathFake := filepath.Join("..", "..", "..", "test.csv")
		headerFake, err := helper.UnitTestingUploadFileMock(filePathFake)
		if err != nil {
			log.Panic("dari file header", err.Error())
		}
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, *headerFake, inputData)
		assert.ErrorContains(t, err, "validate")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})
}

func TestDeactivate(t *testing.T) {
	repo := mocks.NewUserData(t)
	t.Run("deleting account successful", func(t *testing.T) {
		repo.On("Deactivate", uint(1)).Return(nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Deactivate(pToken)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("internal server error, account fail to delete", func(t *testing.T) {
		repo.On("Deactivate", uint(1)).Return(errors.New("no user has delete")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Deactivate(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		repo.AssertExpectations(t)
	})
}
