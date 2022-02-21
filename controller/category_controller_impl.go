package controller

import (
	"github.com/julienschmidt/httprouter"
	"m3gaplazma/go-restapi/helper"
	"m3gaplazma/go-restapi/model/dto"
	"m3gaplazma/go-restapi/service"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := dto.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)

	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "Data was created",
		Data:   categoryResponse,
	}
	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := dto.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)

	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "Data was updated",
		Data:   categoryResponse,
	}
	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)

	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "Data was deleted",
	}
	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(request.Context())

	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	helper.WriteToResponseBody(writer, apiResponse)
}
