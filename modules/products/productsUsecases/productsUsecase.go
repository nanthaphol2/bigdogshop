package productsUsecases

import "github.com/nanthaphol2/bigdogshop/modules/products/productsRepositories"

type IProductsUsecase interface {
}

type productsUsecase struct {
	productsRepository productsRepositories.IProductsRepository
}

func ProductsUsecase(productsRepository productsRepositories.IProductsRepository) IProductsUsecase {
	return &productsUsecase{
		productsRepository: productsRepository,
	}
}
