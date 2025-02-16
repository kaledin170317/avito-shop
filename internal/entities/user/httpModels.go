package user

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type UserResponse struct {
	Username string `json:"username"`
	Coins    int    `json:"coins"`
}
