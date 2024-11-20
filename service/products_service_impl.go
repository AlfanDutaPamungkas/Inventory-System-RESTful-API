package service

import (
	"context"
	"database/sql"
	"inventory-system-api/exception"
	"inventory-system-api/helper"
	"inventory-system-api/model/domain"
	"inventory-system-api/model/web"
	"inventory-system-api/repository"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
)

type ProductsServiceImpl struct {
	DB                 *sql.DB
	ProductsRepository repository.ProductsRepository
	StockRepository    repository.StockRepository
	cld                *cloudinary.Cloudinary
	Validate           *validator.Validate
}

func NewProductServiceImpl(DB *sql.DB, productsRepository repository.ProductsRepository, stockRepository repository.StockRepository, cld *cloudinary.Cloudinary, validate *validator.Validate) ProductsService {
	return &ProductsServiceImpl{
		DB:                 DB,
		ProductsRepository: productsRepository,
		StockRepository:    stockRepository,
		cld:                cld,
		Validate:           validate,
	}
}

func (service *ProductsServiceImpl) CreateProductService(ctx context.Context, request web.ProductCreateReq, file multipart.File, fileHeader *multipart.FileHeader) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	var imageUrl string

	if file != nil && fileHeader != nil {
		imageUrl = helper.UploadImage(ctx, service.cld, file, fileHeader)
	} else {
		imageUrl = ""
	}

	var expDate *time.Time

	if request.ExpiredDate != nil {
		expDate = request.ExpiredDate
	}

	product := domain.Products{
		SKU:         request.SKU,
		Name:        request.Name,
		Brand:       request.Brand,
		Category:    request.Category,
		ImageUrl:    imageUrl,
		Price:       request.Price,
		Amount:      request.Amount,
		ExpiredDate: expDate,
	}

	product = service.ProductsRepository.Create(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductsServiceImpl) FindAllService(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductsRepository.FindAll(ctx, tx)

	var productResponses []web.ProductResponse

	for _, product := range products {
		productResponse := helper.ToProductResponse(product)
		productResponses = append(productResponses, productResponse)
	}

	return productResponses
}

func (service *ProductsServiceImpl) FindBySKUService(ctx context.Context, SKU string) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductsRepository.FindBySKU(ctx, tx, SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	return helper.ToProductResponse(product)
}

func (service *ProductsServiceImpl) UpdateProductService(ctx context.Context, request web.ProductUpdateReq) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductsRepository.FindBySKU(ctx, tx, request.SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	if request.Name == "" && request.Brand == "" && request.Category == "" && request.Price == 0 && request.Amount == 0 {
		panic(exception.NewBadReqErr("no fields updated"))
	}

	if request.Name != "" {
		product.Name = request.Name
	}

	if request.Brand != "" {
		product.Brand = request.Brand
	}

	if request.Category != "" {
		product.Category = request.Category
	}

	if request.Price != 0 {
		product.Price = request.Price
	}

	if request.Amount != 0 {
		product.Amount = request.Amount
	}

	if request.ExpiredDate != nil {
		product.ExpiredDate = request.ExpiredDate
	}
	product = service.ProductsRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductsServiceImpl) StockOutService(ctx context.Context, request web.StockAmountReq) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductsRepository.FindBySKU(ctx, tx, request.SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	if request.Amount > product.Amount {
		panic(exception.NewBadReqErr("the limit amount issued from stock is 0"))
	}

	product.Amount = -request.Amount

	product = service.StockRepository.StockOut(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductsServiceImpl) StockInService(ctx context.Context, request web.StockAmountReq) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductsRepository.FindBySKU(ctx, tx, request.SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	product.Amount = request.Amount

	product = service.StockRepository.StockOut(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductsServiceImpl) UpdateImgUrlService(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, SKU string) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductsRepository.FindBySKU(ctx, tx, SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	imageUrl := helper.UploadImage(ctx, service.cld, file, fileHeader)

	product.ImageUrl = imageUrl

	product = service.ProductsRepository.UpdateImgUrl(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductsServiceImpl) NullifyExpiredDateService(ctx context.Context, SKU string) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductsRepository.FindBySKU(ctx, tx, SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	product = service.StockRepository.NullifyExpiredDate(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductsServiceImpl) Delete(ctx context.Context, SKU string) string {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.ProductsRepository.FindBySKU(ctx, tx, SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	service.ProductsRepository.Delete(ctx, tx, SKU)

	return "product succesfully deleted"
}
