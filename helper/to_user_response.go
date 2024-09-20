package helper

import (
	"inventory-system-api/model/domain"
	"inventory-system-api/model/web"
)

func ToUserResponse(user domain.Users) web.UsersResponse {
	return web.UsersResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
