// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"m3gaplazma/go-restapi/app"
	"m3gaplazma/go-restapi/controller"
	"m3gaplazma/go-restapi/middleware"
	"m3gaplazma/go-restapi/repository"
	"m3gaplazma/go-restapi/service"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitializeServer() *http.Server {
	categoryRepositoryImpl := repository.NewCategoryRepository()
	db := app.NewDB()
	validate := validator.New()
	categoryServiceImpl := service.NewCategoryService(categoryRepositoryImpl, db, validate)
	categoryControllerImpl := controller.NewCategoryController(categoryServiceImpl)
	router := app.NewRouter(categoryControllerImpl)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var categorySet = wire.NewSet(repository.NewCategoryRepository, wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)), service.NewCategoryService, wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)), controller.NewCategoryController, wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)))
