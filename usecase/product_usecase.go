package usecase

import "go-api/model"

type ProductUsecase struct {
}

func NewProductUseCase() ProductUsecase {
	return ProductUsecase{}
}

func (pu *ProductUsecase) GetProduct() ([]model.Product, error) {
	return []model.Product{}, nil
}
