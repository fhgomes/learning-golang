package skillservice

import (
	"encoding/json"
	"learninggolang.com/learning-golang/core/dto"
	"net/http"
)

func (skService service) Fetch(response http.ResponseWriter, request *http.Request) {
	paginationRequest, err := dto.FromValuePaginationRequestParams(request)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	products, err := skService.usecase.Fetch(paginationRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(products)
}
