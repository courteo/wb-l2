package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	downloadSite(os.Args[1]) // url from command line    
}

func downloadSite(url string) {
	response, err := http.Get(url) 
	if err != nil {
		fmt.Printf("Can't download site: %v\n", err)
		return
	}
	defer response.Body.Close()

	// Get file's name from url
	fileName := filepath.Base(url)
	fmt.Println("filename ", fileName)
	
	// create file
	file, err := os.Create(fileName + ".html")
	if err != nil {
		fmt.Printf("Cannot create file: %v\n", err)
		return
	}
	defer file.Close()

	// copy response data to file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("Can't write data to file: %v\n", err)
		return
	}

	fmt.Printf("file %s is downloaded.\n", fileName)
}