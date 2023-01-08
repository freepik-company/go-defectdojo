package cmd

import (
	"github.com/freepik-company/go-defectdojo/pkg/Client"
	"github.com/freepik-company/go-defectdojo/pkg/DefectDojoApi"
)

type DefectDojo struct {
	client      Client.Client
	engagements *DefectDojoApi.EngagementsClient
	findings    *DefectDojoApi.FindingsClient
	tests       *DefectDojoApi.TestsClient
}

func NewDefectDojo(client Client.Client) *DefectDojo {
	return &DefectDojo{client: client}
}

func (this DefectDojo) Engagement() *DefectDojoApi.EngagementsClient {
	if this.engagements == nil {
		this.engagements = DefectDojoApi.NewEngagementsClient(this.client)
	}

	return this.engagements
}

func (this DefectDojo) Findings() *DefectDojoApi.FindingsClient {
	if this.findings == nil {
		this.findings = DefectDojoApi.NewFindingsClient(this.client)
	}
	return this.findings
}

func (this DefectDojo) Tests() *DefectDojoApi.TestsClient {
	if this.tests == nil {
		this.tests = DefectDojoApi.NewTestsClient(this.client)
	}

	return this.tests
}
