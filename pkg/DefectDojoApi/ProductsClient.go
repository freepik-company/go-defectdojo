package DefectDojoApi

import (
	"encoding/json"
	"fmt"
	"github.com/freepik-company/go-defectdojo/pkg/Client"
)

type ProductsClient struct {
	client Client.Client
}

func NewProductsClient(client Client.Client) *ProductsClient {
	return &ProductsClient{client: client}
}

func (this ProductsClient) All() *ProductResult {
	response := this.client.Get("/api/v2/products/")
	result := new(ProductResult)
	_ = json.Unmarshal([]byte(response), &result)

	return result
}

func (this ProductsClient) Create(product Product) *Product {
	response := this.client.Post("/api/v2/products/", product.String())
	return newProductFromJson(response)
}

func (this ProductsClient) Get(id int) *Product {
	response := this.client.Get(fmt.Sprintf("/api/v2/products/%d/", id))
	return newProductFromJson(response)
}

func (this ProductsClient) Update(product Product) *Product {
	response := this.client.Put(fmt.Sprintf("/api/v2/products/%d/", product.Id), product.String())
	return newProductFromJson(response)
}

func (this ProductsClient) Delete(product Product) {
	_ = this.client.Delete(fmt.Sprintf("/api/v2/products/%d/", product.Id))
}
