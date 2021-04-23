package main

import (
	"fmt"

	"github.com/Blu5Morpheus/FirstGoWebService/models"
)

func main() {
	u := models.User{
		ID: 2,
		FirstName: "Raven",
		LastName: "Darkholme",
	}
	fmt.Println(u)
}