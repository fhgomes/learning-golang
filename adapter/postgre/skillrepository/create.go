package skillrepository

import (
	"context"
	"learninggolang.com/learning-golang/core/domain"
	"learninggolang.com/learning-golang/core/dto"
)

func (repository repository) Create(productRequest *dto.CreateSkillRequest) (*domain.Skill, error) {
	ctx := context.Background()
	product := domain.Skill{}

	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO skill (name, description) VALUES ($1, $2) returning *",
		productRequest.Name,
		productRequest.Description,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
