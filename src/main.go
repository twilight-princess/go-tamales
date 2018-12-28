package main

import (
	"fmt"
	"log"
	"net/http"
	// "io/ioutil"
)

func main() {
	PORT := ":8080"
	log.Fatal(http.ListenAndServe(PORT, nil))
	fmt.Printf("Server is running on %s", PORT)
}