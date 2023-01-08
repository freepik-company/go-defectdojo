package DefectDojoApi

import "github.com/jamuriano/go-defectdojo/pkg/Client"

type TestsClient struct {
	client Client.Client
}

func NewTestsClient(client Client.Client) *TestsClient {
	return &TestsClient{client: client}
}
