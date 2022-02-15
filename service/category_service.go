package service

import (
	"context"
	"m3gaplazma/go-restapi/model/dto"
)

type CategoryService interface {
	Create(ctx context.Context, request dto.CategoryCreateRequest) dto.CategoryResponse
	Update(ctx context.Context, request dto.CategoryUpdateRequest) dto.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) dto.CategoryResponse
	FindAll(ctx context.Context) []dto.CategoryResponse
}
