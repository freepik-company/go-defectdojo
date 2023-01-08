package DefectDojoApi

import (
	"fmt"
	"github.com/jamuriano/go-defectdojo/DefectDojo"
	"testing"
)

func Test_FindingsClient_All(t *testing.T) {
	sut := NewFindings(DefectDojo.CreateClient())
	result := sut.All()

	for _, finding := range result.Results {
		fmt.Println(finding.Title)
	}
}

func Test_FindingsClient_Create(t *testing.T) {
	sut := NewFindings(DefectDojo.CreateClient())

	finding := NewFinding(
		"Create via API",
		SeverityInformational,
		"Creada a traves de mi propio cliente go",
		true,
		false,
		NumericalSeverityInformational,
		47,
		[]int{8},
	)

	result := sut.Create(*finding)

	fmt.Println(result)
}

func TestFindingsClient_Get(t *testing.T) {
	client := DefectDojo.CreateClient()
	sut := NewFindings(client)

	result := sut.Get(63)

	fmt.Println(result)
}
