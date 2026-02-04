package main

import (
	"fmt"
	"log"

	mainservers "github.com/KingrogKDR/api-gateway/main_servers"
	"github.com/KingrogKDR/api-gateway/models"
	"github.com/KingrogKDR/api-gateway/utils"
)

// what happens when this whole system starts (boot up logic -> which services will be awake and which will remain dormant) ??

func main() {
	// bootup logic:
	// load Json, create repo, create server, start server
	users, orders, err := utils.LoadJson("data/users.json", "data/orders.json")
	if err != nil {
		log.Fatal("Failed to load json data:", err)
	}

	repository := models.NewRepository(users, orders)

	server := mainservers.NewServer(repository)

	fmt.Println("Files successfully loaded onto server")
	fmt.Println(server)

}
