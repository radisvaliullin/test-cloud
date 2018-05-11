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
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	"google.golang.org/api/cloudbilling/v1"
)

func main() {
	ctx := context.Background()
	hc, err := google.DefaultClient(ctx, cloudbilling.CloudPlatformScope)
	if err != nil {
		// TODO: Handle error.
		log.Fatalf("default client, err - %v", err)
	}
	client, err := cloudbilling.New(hc)
	if err != nil {
		// TODO: Handle error.
		log.Fatalf("default client, cloudbilling new, err - %v", err)
	}

	// The resource name of the billing account associated with the projects that you want to list. For
	// example, `billingAccounts/012345-567890-ABCDEF`.
	name := "billingAccounts/" + os.Getenv("APP_BILLING_ACCOUNT_ID") // TODO: Update placeholder value.
	log.Printf("name bill acc - %v", name)

	call := client.BillingAccounts.Projects.List(name)
	if err := call.Pages(ctx, func(page *cloudbilling.ListProjectBillingInfoResponse) error {
		for _, v := range page.ProjectBillingInfo {
			// TODO: Use v.
			log.Printf("billing accounts project list: v - %v", v)
		}
		return nil // NOTE: returning a non-nil error stops pagination.
	}); err != nil {
		// TODO: Handle error.
		log.Fatalf("billing accounts, list, err - %v", err)
	}
}
