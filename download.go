package main

import (
	"io"
	"net/http"
	"os"
)

func downloadFile(filepath string, url string) (err error) {

	// Get the data
	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	var out *os.File
	out, err = os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
