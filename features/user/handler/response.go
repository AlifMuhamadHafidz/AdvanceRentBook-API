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

type ProfileResponse struct {
	ID             uint   `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	BirthOfDate    string `json:"birth_of_date"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
}

func ToProfileResponse(data user.Core) ProfileResponse {
	return ProfileResponse{
		ID:             data.ID,
		ProfilePicture: data.ProfilePicture,
		Name:           data.Name,
		Username:       data.Username,
		BirthOfDate:    data.BirthOfDate,
		Email:          data.Email,
		Phone:          data.Phone,
		Address:        data.Address,
	}
}
