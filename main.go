package main

import (
	_ "github.com/go-sql-driver/mysql"
	"m3gaplazma/go-restapi/helper"
	"m3gaplazma/go-restapi/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializeServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
