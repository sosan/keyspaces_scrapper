package update

import (
	"fmt"
	"log"
	"main/httpclient"
	"runtime"

	"github.com/minio/selfupdate"
)

const (
	URI_DOWNLOAD = `https://github.com/sosan/keyspaces_scrapper/releases/download/latest`
	OSTIPO = "linux"
)

func AutoUpdate() (bool, bool) {
	needUpdate := false
	updateOk := false
	
	uriDownload := getDownloadURI()
	err := doUpdate(uriDownload)
	if err != nil {
		log.Fatalf("ERROR | No es posible conectarse")
	}
	log.Printf("%s", OSTIPO)

	return needUpdate, updateOk
}

func getDownloadURI() string {
	fileName := getFileName()
	return fmt.Sprintf("%s/%s", URI_DOWNLOAD, fileName)
}

func getFileName() string {
	ostype := getOsType()
	fileName := "licencias.exe"
	if ostype == "linux" {
		fileName = "licencias"
	}
	return fileName
}

func getOsType() string {
	return runtime.GOOS
}

func doUpdate(url string) error {
	statusCode := 404
	body, statusCode := httpclient.GetRequestRaw(url, "")

    if statusCode != 200 {
		log.Fatalf("ERROR | No es posible conectarse")
		return nil
    }
    err := selfupdate.Apply(*body, selfupdate.Options{})
    if err != nil {
        // error handling
		log.Fatalf("ERROR | No es posible conectarse")
    }
	defer (*body).Close()
    return err
}
