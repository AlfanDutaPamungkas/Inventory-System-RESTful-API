package web

type UsersUpdateReq struct {
	Id    string `json:"id"`
	Name  string `validate:"max=200" json:"name"`
	Email string `validate:"max=200" json:"email"`
}
