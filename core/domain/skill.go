package domain

import (
	"learninggolang.com/learning-golang/core/dto"
	"net/http"
)

type Skill struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SkillService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
}

type SkillUseCase interface {
	Create(skillRequest *dto.CreateSkillRequest) (*Skill, error)
	Fetch(paginationRequest *dto.PaginationRequestParams) (*Pagination[[]Skill], error)
}

// SkillRepository is a contract of database connection adapter layer
type SkillRepository interface {
	Create(skillRequest *dto.CreateSkillRequest) (*Skill, error)
	Fetch(paginationRequest *dto.PaginationRequestParams) (*Pagination[[]Skill], error)
}
