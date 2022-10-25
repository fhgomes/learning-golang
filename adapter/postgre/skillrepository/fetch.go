package skillrepository

import (
	"context"
	"github.com/booscaaa/go-paginate/paginate"
	"learninggolang.com/learning-golang/core/domain"
	"learninggolang.com/learning-golang/core/dto"
)

func (repository repository) Fetch(pagination *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Skill], error) {
	ctx := context.Background()
	skillsResult := []domain.Skill{}
	total := int32(0)

	query, queryCount, err := paginate.Paginate("SELECT * FROM skill").
		Page(pagination.Page).
		Desc(pagination.Descending).
		Sort(pagination.Sort).
		RowsPerPage(pagination.ItemsPerPage).
		SearchBy(pagination.Search, "name", "description").
		Query()

	if err != nil {
		return nil, err
	}

	{
		rows, err := repository.db.Query(
			ctx,
			*query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			product := domain.Skill{}

			rows.Scan(
				&product.ID,
				&product.Name,
				&product.Description,
			)

			skillsResult = append(skillsResult, product)
		}
	}

	{
		err := repository.db.QueryRow(ctx, *queryCount).Scan(&total)

		if err != nil {
			return nil, err
		}
	}

	return &domain.Pagination[[]domain.Skill]{
		Items: skillsResult,
		Total: total,
	}, nil
}
