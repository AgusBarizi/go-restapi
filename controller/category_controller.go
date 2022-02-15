package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryController interface {
	Create(writer http.ResponseWriter, router *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, router *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, router *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, router *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, router *http.Request, params httprouter.Params)
}
