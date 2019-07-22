package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/ml/v1"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	const email = "<service-account>@<your project>.iam.gserviceaccount.com"
	scopes := []string{"https://www.googleapis.com/auth/cloud-platform"}
	privateKey, err := ioutil.ReadFile("secret")
	if err != nil {
		log.Fatalf("Unable to read secret key: %v", err)
	}

	client := getClient(email, privateKey, scopes)
	service := getPlatformService(client)

	result, err := service.Projects.Jobs.List("projects/<your project>").Do()
	if err != nil {
		log.Fatalf("Unable to create training job: %v", err)
	}

	fmt.Printf("%v", result)
}

func getClient(email string, privateKey []byte, scopes []string) *http.Client {
	conf := &jwt.Config{
		Email:      email,
		PrivateKey: privateKey,
		Scopes: scopes,
		TokenURL: google.JWTTokenURL,
	}
	return conf.Client(oauth2.NoContext)
}

func getPlatformService(client *http.Client) *ml.Service {
	service, err := ml.New(client)
	if err != nil {
		log.Fatalf("Unable to create ML service: %v", err)
	}
	return service
}
