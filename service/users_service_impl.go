package service

import (
	"context"
	"database/sql"
	"inventory-system-api/exception"
	"inventory-system-api/helper"
	"inventory-system-api/model/domain"
	"inventory-system-api/model/web"
	"inventory-system-api/repository"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	DB             *sql.DB
	Validate       *validator.Validate
	UserRepository repository.UsersRepository
}

func NewUsersServiceImpl(DB *sql.DB, validate *validator.Validate, userRepository repository.UsersRepository) UsersService {
	return &UsersServiceImpl{
		DB:             DB,
		Validate:       validate,
		UserRepository: userRepository,
	}
}

func (service *UsersServiceImpl) CreateAdminService(ctx context.Context, request web.UsersCreateReq) web.UsersResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err == nil {
		panic(exception.NewBadReqErr("Email already exists!"))
	}

	hashedPassword, err := helper.HashPassword(request.Password)
	helper.PanicError(err)

	user := domain.Users{
		Id:       helper.Uuid(),
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
	}

	user = service.UserRepository.CreateAdmin(ctx, tx, user)
	return helper.ToUserResponse(user)
}
func (service *UsersServiceImpl) LoginService(ctx context.Context, request web.UsersLoginReq) web.TokenResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewUnauthorizedError("invalid credentials"))
	}

	if user.Status == "inactive" {
		panic(exception.NewForbiddenErr("your account is inactive"))
	}

	valid := helper.VerifyPassword(request.Password, user.Password)
	if !valid {
		panic(exception.NewUnauthorizedError("invalid credentials"))
	}

	jwtExpiredTimeToken, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_TOKEN"))
	helper.PanicError(err)

	tokenCreateReq := web.TokenCreateReq{
		UserId: user.Id,
		Role:   user.Role,
	}

	token := web.TokenResponse{
		Token: helper.CreateToken(tokenCreateReq, time.Duration(jwtExpiredTimeToken)),
	}

	return token
}

func (service *UsersServiceImpl) ProfileService(ctx context.Context, id string) web.UsersResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundErr("user not found"))
	}

	return helper.ToUserResponse(user)
}

func (service *UsersServiceImpl) UpdateAdminAccService(ctx context.Context, request web.UsersUpdateReq, id string) web.UsersResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundErr("user not found"))
	}

	if user.Role == "super admin" && user.Id != ctx.Value("userData").(*web.TokenClaims).UserId {
		panic(exception.NewUnauthorizedError("you can't update super admin account"))
	}

	if request.Email == "" && request.Name == "" {
		panic(exception.NewBadReqErr("email or name can't be empty"))
	}

	if request.Name == "" {
		user.Email = request.Email
	} else {
		user.Name = request.Name
	}

	user = service.UserRepository.UpdateAdminAcc(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UsersServiceImpl) FindAllAdminAccService(ctx context.Context) []web.UsersResponse {
	queries := ctx.Value("queries")
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx, queries.(map[string]string))

	var userResponses []web.UsersResponse
	for _, user := range users {
		userResponses = append(userResponses, helper.ToUserResponse(user))
	}

	return userResponses
}

func (service *UsersServiceImpl) DeactiveAdminService(ctx context.Context, id string) web.UsersResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundErr("user not found"))
	}

	user = service.UserRepository.DeactiveAdmin(ctx, tx, id)

	return helper.ToUserResponse(user)
}

func (service *UsersServiceImpl) ChangePasswordService(ctx context.Context, request web.UserUpdatePasswordReq) string {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, ctx.Value("userData").(*web.TokenClaims).UserId)
	if err != nil {
		panic(exception.NewNotFoundErr("user not found"))
	}

	hashedPassword, err := helper.HashPassword(request.Password)
	helper.PanicError(err)

	user.Password = hashedPassword

	service.UserRepository.ChangePassword(ctx, tx, user)

	return "success change password"
}
