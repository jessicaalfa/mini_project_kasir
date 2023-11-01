package res

import (
	"kasir/model"
)

type ProductResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" from:"name"`
	Harga string `json:"harga" from:"harga"`
}

func ProductRes(product []model.Product) []ProductResponse {
	var result []ProductResponse
	for _, p := range product {
		response := ProductResponse{
			ID:    p.ID,
			Name:  p.Name,
			Harga: p.Harga,
		}
		result = append(result, response)
	}
	return result
}

func ProductIDRes(product *model.Product) ProductResponse {
	return ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Harga: product.Harga,
	}
}
