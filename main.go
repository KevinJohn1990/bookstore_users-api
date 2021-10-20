package main

import (
	"fmt"

	"github.com/KevinJohn1990/bookstore_users-api/app"
)

func init() {
	//set env variables
	// environ.SetConfig(true)
	fmt.Println("Inside main init() func")
}

func main() {
	//start applications
	app.StartApplication()
}
