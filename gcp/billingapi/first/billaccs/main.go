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

	call := client.BillingAccounts.List()
	log.Println("proj - %+v", client.BillingAccounts.Projects)
	if err := call.Pages(ctx, func(page *cloudbilling.ListBillingAccountsResponse) error {
		for _, v := range page.BillingAccounts {
			log.Printf("billing accounts: v - %v", v)
		}
		log.Printf("not account bill info")
		return nil // NOTE: returning a non-nil error stops pagination.
	}); err != nil {
		log.Fatalf("billing accounts, list, err - %v", err)
	}
}
