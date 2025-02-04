package models

type User struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
type UserRegisterRequest struct {
	Name     string
	Email    string
	Password string
	Phone    string
}

type UserLoginRequest struct {
	Email    string
	Password string
}

type GetUserInfoRequest struct {
	ID string
}

type GetUserInfoResponse struct {
	Name  string
	Email string
	Phone string
}
