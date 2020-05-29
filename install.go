package main

import (
	"os"
	"strings"
)

func install(packageURL string) (err error) {
	//Download
	urlParts := strings.Split(packageURL, "/")
	fileName := urlParts[len(urlParts)-1]
	err = downloadFile(fileName, packageURL)
	if err != nil {
		return
	}

	//Install
	err = executeCommand("sudo", "installer", "-pkg", fileName, "-target", "/", "-allowUntrusted")
	if err != nil {
		return
	}

	//Cleanup
	os.RemoveAll(fileName)
	return
}

func installAppFromArchive(archiveURL string) (err error) {
	//Download
	urlParts := strings.Split(archiveURL, "/")
	fileName := urlParts[len(urlParts)-1]
	err = downloadFile(fileName, archiveURL)
	if err != nil {
		return
	}

	//Install/Extract
	err = executeCommand("sudo", "unzip", "-q", "-o", fileName, "-d", "/Applications/")
	if err != nil {
		return
	}

	//Cleanup
	os.RemoveAll(fileName)
	return
}
