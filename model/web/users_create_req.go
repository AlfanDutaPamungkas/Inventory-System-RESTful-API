package web

type UsersCreateReq struct {
	Name     string `validate:"required,max=200" json:"name"`
	Email    string `validate:"required,max=200,email" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
}
