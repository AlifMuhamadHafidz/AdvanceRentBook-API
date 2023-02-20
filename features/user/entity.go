package user

import "github.com/labstack/echo/v4"

type Core struct {
	ID             uint
	Name           string
	ProfilePicture string
	Username       string
	BirthOfDate    string
	Email          string
	Password       string
	Phone          string
	Address        string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) (Core, error)
	Login(username, password string) (Core, error)
}

type UserData interface {
	Register(ewUser Core) (Core, error)
	Login(username string) (Core, error)
}
