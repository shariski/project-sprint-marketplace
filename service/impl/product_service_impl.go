package impl

import (
	"context"
	"database/sql"
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/model"
	"project-sprint-marketplace/repository"
	"project-sprint-marketplace/service"
)

type productServiceImpl struct {
	*sql.DB
	repository.ProductRepository
	repository.TagRepository
	repository.UserRepository
}

func NewProductServiceImpl(
	DB *sql.DB,
	productRepository *repository.ProductRepository,
	tagRepository *repository.TagRepository,
	userRepository *repository.UserRepository,
) service.ProductService {
	return &productServiceImpl{
		DB:                DB,
		ProductRepository: *productRepository,
		TagRepository:     *tagRepository,
		UserRepository:    *userRepository,
	}
}

func (productService *productServiceImpl) Create(ctx context.Context, data model.ProductCreateModel) entity.Product {
	product := entity.Product{
		UserId:         data.UserId,
		Name:           data.Name,
		Price:          data.Price,
		ImageUrl:       data.ImageUrl,
		Stock:          data.Stock,
		Condition:      data.Condition,
		IsPurchaseable: data.IsPurchaseable,
	}

	tx, err := productService.DB.Begin()
	exception.PanicLogging(err)
	defer common.CommitOrRollback(tx)

	newProduct := productService.ProductRepository.Insert(ctx, tx, product)
	// fmt.Println(newProduct.Id)

	for _, tagName := range data.Tags {
		tag := entity.Tag{
			ProductId: newProduct.Id,
			Name:      tagName,
		}

		_ = productService.TagRepository.Insert(ctx, tx, tag)
	}

	return newProduct
}

func (productService *productServiceImpl) Update(ctx context.Context, data model.ProductUpdateModel) entity.Product {
	product := entity.Product{
		Id:             data.Id,
		Name:           data.Name,
		Price:          data.Price,
		ImageUrl:       data.ImageUrl,
		Condition:      data.Condition,
		IsPurchaseable: data.IsPurchaseable,
	}

	// _ = productService.ProductRepository.FindById(ctx, productService.DB, data.Id)

	tx, err := productService.DB.Begin()
	exception.PanicLogging(err)
	defer common.CommitOrRollback(tx)

	updatedProduct := productService.ProductRepository.Update(ctx, tx, product)

	productService.TagRepository.DeleteByProductId(ctx, tx, data.Id)

	for _, tagName := range data.Tags {
		tag := entity.Tag{
			ProductId: data.Id,
			Name:      tagName,
		}

		_ = productService.TagRepository.Insert(ctx, tx, tag)
	}

	return updatedProduct
}

func (productService *productServiceImpl) DeleteById(ctx context.Context, id int) {
	_ = productService.ProductRepository.FindById(ctx, productService.DB, id)

	tx, err := productService.DB.Begin()
	exception.PanicLogging(err)
	defer common.CommitOrRollback(tx)

	productService.TagRepository.DeleteByProductId(ctx, tx, id)
	productService.ProductRepository.DeleteByProductId(ctx, tx, id)
}

func (productService *productServiceImpl) FindById(ctx context.Context, id int) model.GetProductModel {
	product := productService.ProductRepository.FindById(ctx, productService.DB, id)
	seller := productService.UserRepository.FindByProductId(ctx, id)

	payload := model.GetProductModel{
		Product: product,
		Seller:  seller,
	}

	return payload
}

func (productService *productServiceImpl) UpdateStockById(ctx context.Context, data model.UpdateStockModel) entity.Product {
	product := entity.Product{
		Id:    data.Id,
		Stock: data.Stock,
	}

	_ = productService.ProductRepository.FindById(ctx, productService.DB, data.Id)

	tx, err := productService.DB.Begin()
	exception.PanicLogging(err)
	defer common.CommitOrRollback(tx)

	updatedProduct := productService.ProductRepository.UpdateStock(ctx, tx, product)

	return updatedProduct
}
