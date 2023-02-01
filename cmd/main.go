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

	fmt.Println("CONFIG", cfg)
	if err != nil {
		log.Fatal(err)
	}

	newCl := apiclient.Start(cfg)

	//users, err := newCl.GetUsers()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//women := users.FilterByGender(female)
	//if women == nil {
	//	log.Fatal("no users found")
	//}
	//
	//for _, woman := range women {
	//	fmt.Printf("Name: %s\nEmail: %s\nGender: %s\n\n", woman.Name, woman.Email, woman.Gender)
	//}

	//newCl.CreateUser(apiclient.User{
	//	Name:   "oleg",
	//	Email:  "oleg@i.ua",
	//	Gender: "male",
	//	Status: "active",
	//})

	if err := newCl.CreateUser(apiclient.User{
		Name:   "olejka",
		Email:  "olejka@i.ua",
		Gender: "female",
		Status: "inactive",
	}); err != nil {
		log.Fatal(err)
	}

}
