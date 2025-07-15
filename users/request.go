package dto

type GetUserRequest struct {
	Email string `json:"email" binding:"required,email"`
}
