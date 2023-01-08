package cmd

import (
	"github.com/jamuriano/go-defectdojo/pkg/Client"
	"github.com/jamuriano/go-defectdojo/pkg/DefectDojoApi"
)

type DefectDojo struct {
	client   Client.Client
	findings *DefectDojoApi.FindingsClient
	tests    *DefectDojoApi.TestsClient
}

func NewDefectDojo(client Client.Client) *DefectDojo {
	return &DefectDojo{client: client}
}

func (dojo DefectDojo) Findings() *DefectDojoApi.FindingsClient {
	if dojo.findings == nil {
		dojo.findings = DefectDojoApi.NewFindings(dojo.client)
	}
	return dojo.findings
}
