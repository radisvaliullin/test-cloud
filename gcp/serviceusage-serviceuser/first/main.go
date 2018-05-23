package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	serviceusage "google.golang.org/api/serviceusage/v1"
	serviceuser "google.golang.org/api/serviceuser/v1"
)

func main() {

	client, err := google.DefaultClient(oauth2.NoContext, "https://www.googleapis.com/auth/bigquery")
	if err != nil {
		log.Fatal(err)
	}

	// service user
	apiServUser, err := serviceuser.New(client)
	if err != nil {
		log.Fatal(err)
	}

	projID := os.Getenv("APP_PROJECT_ID")
	log.Println("proj id - ", projID)
	pslCall := apiServUser.Projects.Services.List(projID)
	fmt.Printf("proj serv list - %v\n", pslCall)

	// l, err := pslCall.Do()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("l - %+v\n", l)

	// service usage
	apiServUsage, err := serviceusage.New(client)
	if err != nil {
		log.Fatal(err)
	}

	sleCall := apiServUsage.Services.ListEnabled(os.Getenv("APP_PROJECT_ID"))
	fmt.Printf("serv list enabled - %+v\n", sleCall)
}
