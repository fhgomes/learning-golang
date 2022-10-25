package skillusecase

import (
	"learninggolang.com/learning-golang/core/domain"
	"learninggolang.com/learning-golang/core/dto"
)

func (ucSkill usecase) Fetch(paginationRequest *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Skill], error) {
	products, err := ucSkill.repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return products, nil
}
