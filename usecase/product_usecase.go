package usecase

import (
	"go-api/model"
	"go-api/repository"
)

// Se o nome da estrutura iniciar com letra minuscula ela age como estrutura privada dentro do pacote, se iniciar com maiuscula é visivel global

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

// funcao retornando uma lista de produtos
func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}
