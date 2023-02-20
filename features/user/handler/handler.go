package handler

import (
	"advancerentbook-api/features/user"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type userControll struct {
	srv user.UserService
}

// Deactivate implements user.UserHandler
func (*userControll) Deactivate() echo.HandlerFunc {
	panic("unimplemented")
}

// Profile implements user.UserHandler
func (*userControll) Profile() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements user.UserHandler
func (*userControll) Update() echo.HandlerFunc {
	panic("unimplemented")
}

func New(srv user.UserService) user.UserHandler {
	return &userControll{
		srv: srv,
	}
}

func (uc *userControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}

		res, err := uc.srv.Register(*ReqToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "already") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already registered"})
			} else if strings.Contains(err.Error(), "is not min") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "validate: password length minimum 3 character"})
			} else if strings.Contains(err.Error(), "validate") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{"message": "success create account"})
	}
}

func (uc *userControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}

		if input.Username == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "username not allowed empty"})
		} else if input.Password == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "password not allowed empty"})
		}

		token, res, err := uc.srv.Login(input.Username, input.Password)
		if err != nil {
			if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "password not match"})
			} else {
				return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "account not registered"})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    ToResponse(res),
			"token":   token,
			"message": "success login",
		})
	}
}
