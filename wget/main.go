package main

import (
	"fmt"
	"net"
	"time"
	"net/http"
	"io"
	"os"
	"strings"
)

func checkNetworkConnectivity() {
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", "google.com:80", timeout)
	if err != nil {
		fmt.Println("Error: Network is not reachable")
	}
	conn.Close()
}

func downloadFile(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: Failed to download file")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Failed to download file")
		return
	}

	// Extract filename from URL
	fileName := url[strings.LastIndex(url, "/")+1:]

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error: Failed to create file")
	}
	defer file.Close()

	io.Copy(file, response.Body)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wgetgo <url>")
		return
	}

	url := os.Args[1]
	checkNetworkConnectivity()
	downloadFile(url)
}
