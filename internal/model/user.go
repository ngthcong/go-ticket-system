package model

import "time"

type (
	Users struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Salt      string    `json:"salt"`
		Birthday  time.Time `json:"birthday"`
		Phone     string    `json:"phone"`
		WorkPlace string    `json:"workPlace"`
		Role      int       `json:"role"`
	}
	UserRequest struct {
		Id int `json:"id"`
	}
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
