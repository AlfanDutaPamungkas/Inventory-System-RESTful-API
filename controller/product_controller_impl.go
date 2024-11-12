package controller

import (
	"context"
	"inventory-system-api/helper"
	"inventory-system-api/model/web"
	"inventory-system-api/service"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type ProductsControllerImpl struct {
	productService     service.ProductsService
	LogActivityService service.LogActivityService
}

func NewProductsControllerImpl(productService service.ProductsService, logActivityService service.LogActivityService) ProductsController {
	return &ProductsControllerImpl{
		productService:     productService,
		LogActivityService: logActivityService,
	}
}

func (controller *ProductsControllerImpl) CreateController(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	file, fileHeader, err := r.FormFile("Image")
	if err != nil {
		file = nil
		fileHeader = nil
	} else {
		defer file.Close()
	}

	productCreateReq := web.ProductCreateReq{}

	if r.FormValue("expired_at") != "" {
		parsedDate, err := time.Parse("2006-01-02", r.FormValue("expired_at"))
		helper.PanicError(err)
		productCreateReq.ExpiredDate = &parsedDate
	}

	r.PostForm.Del("expired_at")
	helper.FormToReq(r, &productCreateReq)

	response := controller.productService.CreateProductService(r.Context(), productCreateReq, file, fileHeader)

	message := "Create product with SKU " + productCreateReq.SKU
	controller.LogActivityService.CreateService(r.Context(), message, response.UpdateAt) 

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(w, webResponse)
}

func (controller *ProductsControllerImpl) FindAllController(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := r.URL.Query().Get("name")

	ctx := context.WithValue(r.Context(), "query_name", name)
	r = r.WithContext(ctx)

	response := controller.productService.FindAllService(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}
	helper.WriteToBody(w, webResponse)
}

func (controller *ProductsControllerImpl) FindBySKUController(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	SKU := p.ByName("sku")

	response := controller.productService.FindBySKUService(r.Context(), SKU)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}
	helper.WriteToBody(w, webResponse)
}

func (controller *ProductsControllerImpl) UpdateController(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productUpdateReq := web.ProductUpdateReq{}
	productUpdateReq.SKU = p.ByName("sku")

	if r.FormValue("expired_at") != "" {
		parsedDate, err := time.Parse("2006-01-02", r.FormValue("expired_at"))
		helper.PanicError(err)
		productUpdateReq.ExpiredDate = &parsedDate
	}

	r.PostForm.Del("expired_at")
	helper.FormToReq(r, &productUpdateReq)

	response := controller.productService.UpdateProductService(r.Context(), productUpdateReq)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	message := "Update product with SKU " + productUpdateReq.SKU
	controller.LogActivityService.CreateService(r.Context(), message, response.UpdateAt)

	helper.WriteToBody(w, webResponse)
}

func (controller *ProductsControllerImpl) StockOutController(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var message string
	stockAmountReq := web.StockAmountReq{}
	stockAmountReq.SKU = p.ByName("sku")

	helper.FormToReq(r, &stockAmountReq)

	response := controller.productService.StockOutService(r.Context(), stockAmountReq)

	if stockAmountReq.Amount == 1 {
		message = strconv.Itoa(stockAmountReq.Amount) + " unit of stock have been released"
	}else{	
		message = strconv.Itoa(stockAmountReq.Amount) + " units of stock have been released"
	}

	controller.LogActivityService.CreateService(r.Context(), message, response.UpdateAt)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(w, webResponse)
}

func (controller *ProductsControllerImpl) StockInController(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var message string
	stockAmountReq := web.StockAmountReq{}
	stockAmountReq.SKU = p.ByName("sku")

	helper.FormToReq(r, &stockAmountReq)

	response := controller.productService.StockInService(r.Context(), stockAmountReq)

	if stockAmountReq.Amount == 1 {
		message = strconv.Itoa(stockAmountReq.Amount) + " unit of stock have been received"
	}else{	
		message = strconv.Itoa(stockAmountReq.Amount) + " units of stock have been received"
	}

	controller.LogActivityService.CreateService(r.Context(), message, response.UpdateAt)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(w, webResponse)
}

func (controller *ProductsControllerImpl) UpdateImgUrlController(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	file, fileHeader, err := r.FormFile("Image")
	helper.PanicError(err)
	defer file.Close()

	response := controller.productService.UpdateImgUrlService(r.Context(), file, fileHeader, p.ByName("sku"))

	message := "Change product " + p.ByName("sku") + " image"
	controller.LogActivityService.CreateService(r.Context(), message, response.UpdateAt) 

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(w, webResponse)
}

func (controller *ProductsControllerImpl) NullifyExpiredDateController(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response := controller.productService.NullifyExpiredDateService(r.Context(), p.ByName("sku"))

	message := "Clear product " + p.ByName("sku") + " expired date"
	controller.LogActivityService.CreateService(r.Context(), message, response.UpdateAt)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(w, webResponse)
}

func (controller *ProductsControllerImpl) DeleteController(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response := controller.productService.Delete(r.Context(), p.ByName("sku"))

	message := "Delete product " + p.ByName("sku")
	controller.LogActivityService.CreateService(r.Context(), message, time.Now())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(w, webResponse)
}
