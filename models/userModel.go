package models

import "errors"

type RegisterRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterResponseModel struct {
	Username string `json:"username"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UUID string `json:"uuid"`

}

type LoginRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseModel struct {
	UUID string `json:"uuid"`
}

func (rq *RegisterRequestModel) Validate() ( err error) {
	if len(rq.Password) < 6 {
		err = errors.New("panjang passwod harus lebih dari 6 character")
		return
	}
	return
}
