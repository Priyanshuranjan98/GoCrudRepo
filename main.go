package main

import (
	"fmt"
	"log"
	"mongoexample/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDb Api..")
	r := router.Router()
	fmt.Println("Server is Getting Started ...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listining on port 4000")
}
