package DefectDojoApi

import (
	"encoding/json"
	"fmt"
	"github.com/jamuriano/go-defectdojo/pkg/Client"
)

const (
	SeverityCritical               = "Critical"
	SeverityHigh                   = "High"
	SeverityMedium                 = "Medium"
	SeverityLog                    = "Low"
	SeverityInformational          = "Informational"
	NumericalSeverityCritical      = "S0"
	NumericalSeverityHigh          = "S1"
	NumericalSeverityMedium        = "S2"
	NumericalSeverityLow           = "S3"
	NumericalSeverityInformational = "S4"
)

type FindingsClient struct {
	client Client.Client
}

func NewFindings(client Client.Client) *FindingsClient {
	return &FindingsClient{client: client}
}

func (fc FindingsClient) All() FindingResults {
	response := fc.client.Get("/api/v2/findings/")
	result := FindingResults{}
	_ = json.Unmarshal([]byte(response), &result)

	return result
}

func (fc FindingsClient) Create(finding Finding) Finding {
	body, _ := json.Marshal(finding)
	response := fc.client.Post("/api/v2/findings/", string(body))

	result := Finding{}
	_ = json.Unmarshal([]byte(response), &result)

	return result
}

func (fc FindingsClient) Get(id int) Finding {
	response := fc.client.Get(fmt.Sprintf("/api/v2/findings/%d/", id))

	return fc.unmarshalFinding(response)
}

func (fc FindingsClient) unmarshalFinding(response string) Finding {
	result := Finding{}
	_ = json.Unmarshal([]byte(response), &result)

	return result
}

func (fc FindingsClient) Update(finding Finding) Finding {
	body, _ := json.Marshal(finding)
	response := fc.client.Put(fmt.Sprintf("/api/v2/findings/%d/", finding.Id), string(body))

	return fc.unmarshalFinding(response)
}
