package controllers

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"fullName" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type UpdateUserRequest struct {
	Password *string `json:"password,omitempty"`
	FullName *string `json:"fullName,omitempty"`
	Email    *string `json:"email,omitempty"`
}
