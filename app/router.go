package app

import (
	"github.com/julienschmidt/httprouter"
	"m3gaplazma/go-restapi/controller"
	"m3gaplazma/go-restapi/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/category/:categoryId", categoryController.FindById)
	router.POST("/api/category", categoryController.Create)
	router.PUT("/api/category/:categoryId", categoryController.Update)
	router.DELETE("/api/category/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
