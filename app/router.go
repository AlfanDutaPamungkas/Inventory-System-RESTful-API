package app

import (
	"inventory-system-api/controller"
	"inventory-system-api/exception"
	"inventory-system-api/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UsersController, productController controller.ProductsController, logController controller.LogController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/auth/login", userController.LoginCtrl)
	router.GET("/api/auth/profile", middleware.AuthMiddleware(userController.ProfileCtrl))
	router.PUT("/api/auth/profile", middleware.AuthMiddleware(userController.UpdateProfileCtrl))
	router.PUT("/api/auth/profile/change-password", middleware.AuthMiddleware(userController.ChangePasswordCtrl))

	router.GET("/api/superadmin/admins", middleware.AuthMiddleware(middleware.SuperAdminMiddleware(userController.FindAllAdminAccCtrl)))
	router.POST("/api/superadmin/admins", middleware.AuthMiddleware(middleware.SuperAdminMiddleware(userController.CreateAdminCtrl)))
	router.GET("/api/superadmin/admins/:id", middleware.AuthMiddleware(middleware.SuperAdminMiddleware(userController.FindAdminByIdCtrl)))
	router.PUT("/api/superadmin/admins/:id", middleware.AuthMiddleware(middleware.SuperAdminMiddleware(userController.UpdateAdminAccCtrl)))
	router.PUT("/api/superadmin/admins/:id/deactive", middleware.AuthMiddleware(middleware.SuperAdminMiddleware(userController.DeactiveAdminCtrl)))

	router.POST("/api/products", middleware.AuthMiddleware(productController.CreateController))
	router.GET("/api/products", middleware.AuthMiddleware(productController.FindAllController))
	router.GET("/api/products/:sku", middleware.AuthMiddleware(productController.FindBySKUController))
	router.PUT("/api/products/:sku", middleware.AuthMiddleware(productController.UpdateController))
	router.PUT("/api/products/:sku/out", middleware.AuthMiddleware(productController.StockOutController))
	router.PUT("/api/products/:sku/in", middleware.AuthMiddleware(productController.StockInController))
	router.PUT("/api/products/:sku/change-img", middleware.AuthMiddleware(productController.UpdateImgUrlController))
	router.PUT("/api/products/:sku/reset-expdate", middleware.AuthMiddleware(productController.NullifyExpiredDateController))
	router.DELETE("/api/products/:sku", middleware.AuthMiddleware(middleware.SuperAdminMiddleware(productController.DeleteController)))

	router.GET("/api/logs", middleware.AuthMiddleware(middleware.SuperAdminMiddleware(logController.FindAllCtrl)))

	router.PanicHandler = exception.ErrorHandler
	return router
}
