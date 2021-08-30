package main

import (
	"fmt"

	"github.com/cctechwiz/go-core-language-path/go_getting_started/webservice/models"
)

func main() {
	u := models.User{
		ID: 2,
		FirstName: "Tricia",
		LastName: "McMillan",
	}

	fmt.Println(u)
}

