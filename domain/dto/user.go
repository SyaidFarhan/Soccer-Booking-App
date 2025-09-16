package dto

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phone_number"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type RegisterRequest struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required,min=6"`
	RoleID          uint   `json:"role_id" binding:"required"`
	PhoneNumber     string `json:"phone_number" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type RegisterResponse struct {
	User UserResponse `json:"user"`
}

type UpdateUserRequest struct {
	Name            string `json:"name" binding:"omitempty"`
	Email           string `json:"email" binding:"omitempty,email"`
	Username        string `json:"username" binding:"omitempty"`
	Password        string `json:"password" binding:"omitempty,min=6"`
	RoleID          uint   `json:"role_id" binding:"omitempty"`
	PhoneNumber     string `json:"phone_number" binding:"omitempty"`
	ConfirmPassword string `json:"confirm_password" binding:"omitempty,eqfield=Password"`
}

type UpdateUserResponse struct {
	User UserResponse `json:"user"`
}

type ChangePasswordRequest struct {
	OldPassword        string `json:"old_password" binding:"required"`
	NewPassword        string `json:"new_password" binding:"required,min=6"`
	ConfirmNewPassword string `json:"confirm_new_password" binding:"required,eqfield=NewPassword"`
}
