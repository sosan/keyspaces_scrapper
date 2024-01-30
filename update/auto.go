package update

import (
	"log"
	"main/httpclient"

	"github.com/minio/selfupdate"
)

func AutoUpdate() (bool, bool) {
	needUpdate := false
	updateOk := false

	err := doUpdate("https://github.com/sosan/keyspaces_scrapper/releases/download/main/licencias")
	if err != nil {
		log.Fatalf("ERROR | No es posible conectarse")
	}

	return needUpdate, updateOk
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
