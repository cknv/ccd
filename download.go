package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"time"
)

func main() {
	fileLocation := "http://download.thinkbroadband.com/5MB.zip"

	start := time.Now()

	resp, err := http.Get(fileLocation)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("size of file:", len(data))

	log.Println("time taken:", time.Since(start))
}
