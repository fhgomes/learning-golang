package skillusecase

import (
	"learninggolang.com/learning-golang/core/domain"
	"learninggolang.com/learning-golang/core/dto"
)

func (ucSkill usecase) Create(req *dto.CreateSkillRequest) (*domain.Skill, error) {
	product, err := ucSkill.repository.Create(req)

	if err != nil {
		return nil, err
	}

	return product, nil
}
