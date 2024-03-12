package impl

import (
	"context"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/model"
	"project-sprint-marketplace/repository"
	"project-sprint-marketplace/service"
)

type productServiceImpl struct {
	repository.ProductRepository
	repository.TagRepository
}

func NewProductServiceImpl(
	productRepository *repository.ProductRepository,
	tagRepository *repository.TagRepository,
	) service.ProductService {
	return &productServiceImpl{
		ProductRepository: *productRepository,
		TagRepository: *tagRepository,
	}
}

func (productService *productServiceImpl) Create(ctx context.Context, data model.ProductCreateModel) error {
	product := entity.Product{
		UserId:		data.UserId,
		Name:		data.Name,
		Price:	data.Price,
		ImageUrl: data.ImageUrl,
		Stock: data.Stock,
		Condition: data.Condition,
		IsPurchaseable: data.IsPurchaseable,
	}

	productId, err := productService.ProductRepository.Insert(ctx, product)
	if err != nil {
		return err
	}
	
	for _,tagName := range data.Tags{
		if tagName != "" {
			tag := entity.Tag{
				ProductId: productId,
				Name: tagName,
			}

			err = productService.TagRepository.Insert(ctx, tag)
			if err != nil {
				return err
			}
		}
	}
	
	return err;
}
