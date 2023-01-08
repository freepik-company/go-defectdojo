package DefectDojo

import (
	"github.com/jamuriano/go-defectdojo/pkg/Client"
	"os"
)

func CreateClient() Client.Client {
	client := Client.NewClient(
		os.Getenv("TEST_DEFECTDOJO_HOST"),
		os.Getenv("TEST_DEFECTDOJO_TOKEN"),
	)
	return client
}
