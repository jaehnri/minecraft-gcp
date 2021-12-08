// Package function contains an HTTP Cloud Function.
package function

import (
	"fmt"
	"google.golang.org/api/compute/v1"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
)

type StartInstanceRequest struct {

	// Google Cloud Provider Zone for this request
	Zone string

	// Google Cloud Provider ProjectID for this request
	ProjectID string

	// Compute Engine Instance to start
	Instance string
}

func NewStartInstanceRequest() *StartInstanceRequest {
	return &StartInstanceRequest{
		Zone:      os.Getenv("INSTANCE_ZONE"),
		ProjectID: os.Getenv("PROJECT_ID"),
		Instance:  os.Getenv("INSTANCE_NAME"),
	}
}

func LaunchMineServer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	request := NewStartInstanceRequest()
	resp, err := computeService.Instances.Start(request.ProjectID, request.Zone, request.Instance).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, buildResponse(resp.Status))
}

func buildResponse(status string) string {
	if status == "DONE" {
		return "Servidor já está aberto!"
	}
	
	return "Servidor está iniciando. Espere um minuto."
}
