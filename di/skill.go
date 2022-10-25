package di

import (
	"learninggolang.com/learning-golang/adapter/http/skillservice"
	postgres "learninggolang.com/learning-golang/adapter/postgre"
	"learninggolang.com/learning-golang/adapter/postgre/skillrepository"
	"learninggolang.com/learning-golang/core/domain"
	"learninggolang.com/learning-golang/core/domain/usecase/skillusecase"
)

func ConfigProductDI(conn postgres.PoolInterface) domain.SkillService {
	productRepository := skillrepository.New(conn)
	productUseCase := skillusecase.New(productRepository)
	productService := skillservice.New(productUseCase)

	return productService
}
