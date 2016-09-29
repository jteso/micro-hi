package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func healthController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UP")
}

func hiController(w http.ResponseWriter, r *http.Request) {
	ip := resolveLocalIpAddr()
	log.Printf("IP Address: %s", ip)
	fmt.Fprintf(w, ip)
}

func main() {
	http.HandleFunc("/health", healthController)
	http.HandleFunc("/api/hi", hiController)
	printWelcome()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func resolveLocalIpAddr() string {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			return ipnet.IP.String()
		}
	}
	return "local_address unknown"
}

func printWelcome() {
	fmt.Printf("Server listening in port 8080...\n")
	fmt.Printf("Routes available:\n")
	fmt.Printf("---------------------\n")
	fmt.Printf("==> [/health] Health check endpoint\n")
	fmt.Printf("==> [/api/hi] Hello service\n")
}
