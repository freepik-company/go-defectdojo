package DefectDojoApi

import (
	"encoding/json"
	"fmt"
	"github.com/freepik-company/go-defectdojo/pkg/Client"
)

type TestsClient struct {
	client Client.Client
}

func NewTestsClient(client Client.Client) *TestsClient {
	return &TestsClient{client: client}
}

func (this TestsClient) All() *TestResult {
	response := this.client.Get("/api/v2/tests/")
	result := new(TestResult)
	_ = json.Unmarshal([]byte(response), &result)

	return result
}

func (this TestsClient) Create(test Test) *Test {
	response := this.client.Post("/api/v2/tests/", test.String())
	return newTestFromJson(response)
}

func (this TestsClient) Get(id int) *Test {
	response := this.client.Get(fmt.Sprintf("/api/v2/tests/%d/", id))
	return newTestFromJson(response)
}

func (this TestsClient) Update(test Test) *Test {
	response := this.client.Put(fmt.Sprintf("/api/v2/tests/%d/", test.Id), test.String())
	return newTestFromJson(response)
}

func (this TestsClient) Delete(test Test) {
	_ = this.client.Delete(fmt.Sprintf("/api/v2/tests/%d/", test.Id))
}
