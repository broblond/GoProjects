package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://$paste_your_URL_here")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
}
//end
