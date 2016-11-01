package main

import (
	"log"
	"net"
	"time"
)

func main() {
	goodHost := "google.com"

	for {
		dnsResp, err := dnsLookup(goodHost)

		if err != nil {
			log.Println("could not resolve host")
		} else {
			log.Println("host resolved to:", dnsResp)
		}

		time.Sleep(5 * time.Second)
	}

}

func dnsLookup(hostname string) (string, error) {
	ipAddr, err := net.ResolveIPAddr("ip", hostname)
	return ipAddr.String(), err
}
