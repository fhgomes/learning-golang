package skillrepository

import (
	postgres "learninggolang.com/learning-golang/adapter/postgre"
	"learninggolang.com/learning-golang/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of SkillRepository
func New(db postgres.PoolInterface) domain.SkillRepository {
	return &repository{
		db: db,
	}
}
