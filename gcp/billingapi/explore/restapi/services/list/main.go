package main

// PRE-REQUISITES:
// ---------------
// 1. If not already done, enable the Google Cloud Billing API and check the quota for your project at
//    https://console.developers.google.com/apis/api/cloudbilling_component/quotas
// 2. This sample uses Application Default Credentials for authentication. If not already done, install the gcloud CLI from
//    https://cloud.google.com/sdk/ and run 'gcloud beta auth application-default login'
// 3. To install the client library, run:
//    go get -u google.golang.org/api/cloudbilling/v1

import (
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	"google.golang.org/api/cloudbilling/v1"
)

func main() {
	ctx := context.Background()
	hc, err := google.DefaultClient(ctx, cloudbilling.CloudPlatformScope)
	if err != nil {
		log.Fatalf("default client, err - %v", err)
	}
	client, err := cloudbilling.New(hc)
	if err != nil {
		log.Fatalf("default client, cloudbilling new, err - %v", err)
	}

	//
	// name := "projects/" + os.Getenv("APP_PROJECT_ID")
	// log.Printf("proj name - %v", name)

	resp, err := client.Services.List().Context(ctx).Do()
	if err != nil {
		log.Fatalf("billing, services list err - %v\n", err)
	}
	log.Printf("billing, services list - %+v", resp)
}
