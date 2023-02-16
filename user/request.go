package user

import "encoding/json"

type UserRequest struct {
	Email    string      `json:"email" binding:"required"`
	Password string      `json:"password" binding:"required"`
	Name     string      `json:"name" binding:"required"`
	Age      json.Number `json:"age"`
}
