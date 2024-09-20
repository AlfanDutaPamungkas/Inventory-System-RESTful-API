package web

type UserUpdatePasswordReq struct{
	Id string `json:"id"`
	Password string `json:"password" validate:"required"`
}

