package DefectDojoApi

import (
	"github.com/freepik-company/go-defectdojo/pkg/Client"
	"os"
)

func CreateClient() Client.Client {
	client := Client.NewClient(
		os.Getenv("TEST_DEFECTDOJO_HOST"),
		os.Getenv("TEST_DEFECTDOJO_TOKEN"),
	)
	return client
}
