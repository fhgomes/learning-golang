package skillservice

import "learninggolang.com/learning-golang/core/domain"

type service struct {
	usecase domain.SkillUseCase
}

func New(usecase domain.SkillUseCase) domain.SkillService {
	return &service{
		usecase: usecase,
	}
}
