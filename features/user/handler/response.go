package handler

import (
	"advancerentbook-api/features/user"
	"errors"
)

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

func ConvertUpdateResponse(input user.Core) (interface{}, error) {
	ResponseFilter := user.Core{}
	ResponseFilter = input
	result := make(map[string]interface{})
	if ResponseFilter.ID != 0 {
		result["id"] = ResponseFilter.ID
	}
	if ResponseFilter.ProfilePicture != "" {
		result["profile_picture"] = ResponseFilter.ProfilePicture
	}
	if ResponseFilter.Name != "" {
		result["name"] = ResponseFilter.Name
	}
	if ResponseFilter.Username != "" {
		result["username"] = ResponseFilter.Username
	}
	if ResponseFilter.BirthOfDate != "" {
		result["birth_of_date"] = ResponseFilter.BirthOfDate
	}
	if ResponseFilter.Email != "" {
		result["email"] = ResponseFilter.Email
	}
	if ResponseFilter.Phone != "" {
		result["phone"] = ResponseFilter.Phone
	}
	if ResponseFilter.Address != "" {
		result["address"] = ResponseFilter.Address
	}
	if ResponseFilter.Password != "" {
		result["password"] = ResponseFilter.Password
	}

	if len(result) <= 1 {
		return user.Core{}, errors.New("no data was change")
	}
	return result, nil
}
