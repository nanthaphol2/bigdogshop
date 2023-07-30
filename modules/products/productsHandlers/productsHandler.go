package productsHandlers

import (
	"github.com/nanthaphol2/bigdogshop/config"
	"github.com/nanthaphol2/bigdogshop/modules/files/filesUsecases"
	"github.com/nanthaphol2/bigdogshop/modules/products/productsUsecases"
)

type productsHandlersErrCode string

const (
	findOneProductErr productsHandlersErrCode = "products-001"
	findProductErr    productsHandlersErrCode = "products-002"
	insertProductErr  productsHandlersErrCode = "products-003"
	deleteProductErr  productsHandlersErrCode = "products-004"
	updateProductErr  productsHandlersErrCode = "products-005"
)

type IProductsHandler interface {
}

type productsHandler struct {
	cfg             config.IConfig
	productsUsecase productsUsecases.IProductsUsecase
	filesUsecase    filesUsecases.IFilesUsecase
}

func ProductsHandler(cfg config.IConfig, productsUsecase productsUsecases.IProductsUsecase, filesUsecase filesUsecases.IFilesUsecase) IProductsHandler {
	return &productsHandler{
		cfg:             cfg,
		productsUsecase: productsUsecase,
		filesUsecase:    filesUsecase,
	}
}
