package skillusecase

import "learninggolang.com/learning-golang/core/domain"

type usecase struct {
	repository domain.SkillRepository
}

// New returns contract implementation of ProductUseCase
func New(repository domain.SkillRepository) domain.SkillUseCase {
	return &usecase{
		repository: repository,
	}
}
