package service

import (
	"context"

	"github.com/renatocantarino/go-grpc/internals"
	"github.com/renatocantarino/go-grpc/internals/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB internals.Category
}

func NewCategoryService(categoryDb internals.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDb,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	response := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: response,
	}, nil

}
