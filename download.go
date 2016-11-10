package main

import (
	"io/ioutil"
	"log"
	"net/http"
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

	fileSize := float64(len(data))
	log.Println("size of file:", fileSize)

	elapsed := float64(time.Since(start).Seconds())
	log.Println("time taken:", elapsed)

	log.Println("speed is:", fileSize / elapsed, "bytes/second")
}
