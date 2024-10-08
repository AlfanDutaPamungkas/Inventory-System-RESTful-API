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
)

type ProductsServiceImpl struct {
	DB                 *sql.DB
	ProductsRepository repository.ProductsRepository
	StockRepository    repository.StockRepository
	cld                *cloudinary.Cloudinary
}

func NewProductServiceImpl(DB *sql.DB, productsRepository repository.ProductsRepository, stockRepository repository.StockRepository, cld *cloudinary.Cloudinary) ProductsService {
	return &ProductsServiceImpl{
		DB:                 DB,
		ProductsRepository: productsRepository,
		StockRepository:    stockRepository,
		cld:                cld,
	}
}

func (service *ProductsServiceImpl) CreateProductService(ctx context.Context, request web.ProductCreateReq, file multipart.File, fileHeader *multipart.FileHeader) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	imageUrl := helper.UploadImage(ctx, service.cld, file, fileHeader)

	product := domain.Products{
		SKU:      request.SKU,
		Name:     request.Name,
		Brand:    request.Brand,
		Category: request.Category,
		ImageUrl: imageUrl,
		Price:    request.Price,
	}

	product = service.ProductsRepository.Create(ctx, tx, product)

	var expDate *time.Time

	if request.ExpiredDate != nil {
		expDate = request.ExpiredDate
	}

	stock := domain.ProductStock{
		SKU:         request.SKU,
		Amount:      request.Amount,
		ExpiredDate: expDate,
	}

	stock = service.StockRepository.Create(ctx, tx, stock)

	return helper.ToProductResponse(product, stock)
}

func (service *ProductsServiceImpl) FindAllService(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductsRepository.FindAll(ctx, tx)

	stocks := service.StockRepository.FindAll(ctx, tx)
	stockMap := make(map[string]domain.ProductStock)
	for _, stock := range stocks {
		stockMap[stock.SKU] = stock
	}

	var productResponses []web.ProductResponse

	for _, product := range products {
		stock, exist := stockMap[product.SKU]
		if !exist {
			panic(exception.NewBadReqErr("Error"))
		}

		productResponse := helper.ToProductResponse(product, stock)
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

	stock, err := service.StockRepository.FindBySKU(ctx, tx, SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	return helper.ToProductResponse(product, stock)
}

func (service *ProductsServiceImpl) UpdateProductService(ctx context.Context, request web.ProductUpdateReq) web.ProductResponse {
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

	product = service.ProductsRepository.Update(ctx, tx, product)

	stock, err := service.StockRepository.FindBySKU(ctx, tx, request.SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	if request.Amount != 0 {
		stock.Amount = request.Amount
	}

	if request.ExpiredDate != nil {
		stock.ExpiredDate = request.ExpiredDate
	}

	stock = service.StockRepository.Update(ctx, tx, stock)

	return helper.ToProductResponse(product, stock)
}

func (service *ProductsServiceImpl) StockOutService(ctx context.Context, request web.StockAmountReq) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductsRepository.FindBySKU(ctx, tx, request.SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	stock, err := service.StockRepository.FindBySKU(ctx, tx, request.SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	if request.Amount > stock.Amount {
		panic(exception.NewBadReqErr("the limit amount issued from stock is 0"))
	}

	stock.Amount = -request.Amount

	stock = service.StockRepository.StockOut(ctx, tx, stock)

	return helper.ToProductResponse(product, stock)
}

func (service *ProductsServiceImpl) StockInService(ctx context.Context, request web.StockAmountReq) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductsRepository.FindBySKU(ctx, tx, request.SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	stock, err := service.StockRepository.FindBySKU(ctx, tx, request.SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	stock.Amount = request.Amount

	stock = service.StockRepository.StockOut(ctx, tx, stock)

	return helper.ToProductResponse(product, stock)
}

func (service *ProductsServiceImpl) UpdateImgUrlService(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, SKU string) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	stock, err := service.StockRepository.FindBySKU(ctx, tx, SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	product, err := service.ProductsRepository.FindBySKU(ctx, tx, SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	imageUrl := helper.UploadImage(ctx, service.cld, file, fileHeader)

	product.ImageUrl = imageUrl

	product = service.ProductsRepository.UpdateImgUrl(ctx, tx, product)

	return helper.ToProductResponse(product, stock)
}

func (service *ProductsServiceImpl) NullifyExpiredDateService(ctx context.Context, SKU string) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	stock, err := service.StockRepository.FindBySKU(ctx, tx, SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	product, err := service.ProductsRepository.FindBySKU(ctx, tx, SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	stock = service.StockRepository.NullifyExpiredDate(ctx, tx, stock)

	return helper.ToProductResponse(product, stock)
}

func (service *ProductsServiceImpl) Delete(ctx context.Context, SKU string) string {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.StockRepository.FindBySKU(ctx, tx, SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	_, err = service.ProductsRepository.FindBySKU(ctx, tx, SKU)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	service.StockRepository.Delete(ctx, tx, SKU)
	service.ProductsRepository.Delete(ctx, tx, SKU)

	return "product succesfully deleted"
}
