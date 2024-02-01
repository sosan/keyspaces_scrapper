package update

import (
	"log"

	"runtime"
	"github.com/mouuff/go-rocket-update/pkg/provider"
	"github.com/mouuff/go-rocket-update/pkg/updater"
)

const (
	URI_DOWNLOAD = `github.com/sosan/keyspaces_scrapper`//`https://github.com/sosan/keyspaces_scrapper/releases/download/latest`
	OSTIPO = "windows"
)

func AutoUpdate() (bool, bool) {
	needUpdate := false
	updateOk := false
	
	// uriDownload := getDownloadURI()
	err := doUpdate()
	if err != nil {
		log.Fatalf("ERROR | No es posible conectarse")
	}
	log.Printf("%s", OSTIPO)
	log.Fatalf("TERMINADO")
	return needUpdate, updateOk
}

func getFileName() string {
	ostype := getOsType()
	fileName := "licencias.exe.zip"
	if ostype == "linux" {
		fileName = "licencias.zip"
	}
	return fileName
}

func getOsType() string {
	return runtime.GOOS
}

func doUpdate() error {
	fileName := getFileName()
	upD := &updater.Updater{
		Provider: &provider.Github{
			RepositoryURL: URI_DOWNLOAD,
			ArchiveName:   fileName,
		},
		ExecutableName: fileName,
		Version:        "v0.0.0",
	}

	log.Println(upD.Version)
	statusCode, err := upD.Update(); 
	if err != nil {
		log.Println(err)
		log.Println(statusCode)
	}

	// body, statusCode := httpclient.GetRequestRaw(url, "")

    // if statusCode != 200 {
	// 	log.Fatalf("ERROR | No es posible conectarse")
	// 	return nil
    // }


    // err := selfupdate.Apply(*body, selfupdate.Options{})
    // if err != nil {
    //     // error handling
	// 	log.Fatalf("ERROR | No es posible conectarse")
    // }
	// defer (*body).Close()
    return err
}
