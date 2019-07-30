package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/ml/v1"
	"log"
)

func main() {
	client, err := google.DefaultClient(oauth2.NoContext, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		log.Fatalf("Unable to create a client using default token source: %v", err)
	}

	service, err := ml.New(client)
	if err != nil {
		log.Fatalf("Unable to create ML service: %v", err)
	}

	result, err := service.Projects.Jobs.List("projects/<your project>").Do()
	if err != nil {
		log.Fatalf("Unable to create training job: %v", err)
	}

	fmt.Printf("%v", result)
}
