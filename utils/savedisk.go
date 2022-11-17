package utils

import (
	"log"
	"os"
)

func SaveLicence(licencia string ) {

	if err := os.WriteFile("licencia.txt", []byte(licencia), 0o644); err != nil {
		log.Fatal(err)
	}

}