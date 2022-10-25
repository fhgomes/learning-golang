package dto

import (
	"encoding/json"
	"io"
)

type CreateSkillRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func FromJSONCreateProductRequest(body io.Reader) (*CreateSkillRequest, error) {
	createSkillRequest := CreateSkillRequest{}
	if err := json.NewDecoder(body).Decode(&createSkillRequest); err != nil {
		return nil, err
	}

	return &createSkillRequest, nil
}
