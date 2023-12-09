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

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var response []*pb.Category

	for _, cat := range categories {
		item := &pb.Category{
			Id:          cat.ID,
			Name:        cat.Name,
			Description: cat.Description,
		}
		response = append(response, item)
	}

	return &pb.CategoryList{
		Categories: response,
	}, nil

}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {

	category, err := c.CategoryDB.FindById(in.Id)
	if err != nil {
		return nil, err
	}

	response := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return response, nil

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
