package main

import (
	"fmt"
	"github.com/olegvolkov91/api-client-hw/package/apiclient"
	"github.com/olegvolkov91/api-client-hw/package/config"
	"log"
)

const (
	female string = "female"
)

func main() {
	cfg, err := config.Init()

	if err != nil {
		log.Fatal(err)
	}

	newCl := apiclient.Start(cfg)

	u, err := newCl.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	users := u.FilterByGender(female)
	if users == nil {
		log.Fatal("no users found")
	}

	for _, woman := range users[female] {
		fmt.Printf("Name: %s\nEmail: %s\nStatus: %s\n\n", woman.Name, woman.Email, woman.Status)
	}
}
