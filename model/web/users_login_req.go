package web

type UsersLoginReq struct {
	Email    string `validate:"required,max=200" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
}
