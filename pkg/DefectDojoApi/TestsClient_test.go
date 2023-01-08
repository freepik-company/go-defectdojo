package DefectDojoApi

import (
	"fmt"
	"github.com/jamuriano/go-defectdojo/DefectDojo"
	"testing"
)

func TestTestsClient_All(t *testing.T) {
	setUp()
	defer tearDown()

	sut := NewTestsClient(DefectDojo.CreateClient())

	result := sut.All()

	for _, test := range result.Results {
		fmt.Println(test.Title)
	}
}

func tearDown() {

}

func setUp() {

}
