package skillservice

import (
	"encoding/json"
	"learninggolang.com/learning-golang/core/dto"
	"net/http"
)

func (skService service) Create(response http.ResponseWriter, request *http.Request) {
	productRequest, err := dto.FromJSONCreateProductRequest(request.Body)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	product, err := skService.usecase.Create(productRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(product)
}
