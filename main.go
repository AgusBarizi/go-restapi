package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"m3gaplazma/go-restapi/app"
	"m3gaplazma/go-restapi/controller"
	"m3gaplazma/go-restapi/helper"
	"m3gaplazma/go-restapi/middleware"
	"m3gaplazma/go-restapi/repository"
	"m3gaplazma/go-restapi/service"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
