package service

import (
	"context"
	"inventory-system-api/model/web"
)

type UsersService interface {
	CreateAdminService(ctx context.Context, request web.UsersCreateReq) web.UsersResponse
	LoginService(ctx context.Context, request web.UsersLoginReq) web.TokenResponse
	ProfileService(ctx context.Context, id string) web.UsersResponse
	UpdateAdminAccService(ctx context.Context, request web.UsersUpdateReq, id string) web.UsersResponse
	FindAllAdminAccService(ctx context.Context) []web.UsersResponse
	DeactiveAdminService(ctx context.Context, id string) web.UsersResponse
	ChangePasswordService(ctx context.Context, request web.UserUpdatePasswordReq) string
}
