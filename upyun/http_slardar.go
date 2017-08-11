package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// START OMIT
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://10.0.0.193:3130", nil)
	req.Host = "dev.nginx.org"
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	// END OMIT
	fmt.Printf("%s", body)
}
