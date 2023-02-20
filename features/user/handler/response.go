package handler

import "advancerentbook-api/features/user"

type UserReponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ToResponse(data user.Core) UserReponse {
	return UserReponse{
		ID:       data.ID,
		Name:     data.Name,
		Username: data.Username,
		Email:    data.Email,
	}
}
