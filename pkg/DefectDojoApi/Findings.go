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

func (this FindingsClient) All() *FindingResults {
	response := this.client.Get("/api/v2/findings/")
	result := new(FindingResults)
	_ = json.Unmarshal([]byte(response), &result)

	return result
}

func (this FindingsClient) Create(finding Finding) *Finding {
	body, _ := json.Marshal(finding)
	response := this.client.Post("/api/v2/findings/", string(body))

	return this.unmarshalFinding(response)
}

func (this FindingsClient) Get(id int) *Finding {
	response := this.client.Get(fmt.Sprintf("/api/v2/findings/%d/", id))

	return this.unmarshalFinding(response)
}

func (this FindingsClient) unmarshalFinding(response string) *Finding {
	result := new(Finding)
	_ = json.Unmarshal([]byte(response), &result)

	return result
}

func (this FindingsClient) Update(finding Finding) *Finding {
	body, _ := json.Marshal(finding)
	response := this.client.Put(fmt.Sprintf("/api/v2/findings/%d/", finding.Id), string(body))

	return this.unmarshalFinding(response)
}
