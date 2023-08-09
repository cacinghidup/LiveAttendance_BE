package auth

import "time"

type RegisterRequest struct {
	Name     string    `json:"name" validation:"required"`
	Picture  string    `json:"file" validation:"required"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
