package DefectDojoApi

import (
	"encoding/json"
	"fmt"
	"github.com/jamuriano/go-defectdojo/pkg/Client"
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
	body, _ := json.Marshal(test)
	response := this.client.Post("/api/v2/tests/", string(body))

	return this.unmarshalTest(response)
}

func (this TestsClient) unmarshalTest(response string) *Test {
	result := new(Test)
	_ = json.Unmarshal([]byte(response), &result)

	return result
}

func (this TestsClient) Get(id int) *Test {
	response := this.client.Get(fmt.Sprintf("/api/v2/tests/%d/", id))

	return this.unmarshalTest(response)
}

func (this TestsClient) Update(test Test) *Test {
	body, _ := json.Marshal(test)
	response := this.client.Put(fmt.Sprintf("/api/v2/tests/%d/", test.Id), string(body))

	return this.unmarshalTest(response)
}
