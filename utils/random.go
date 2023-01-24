package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"main/httpclient"
	"main/models"

	"github.com/google/uuid"
)

func GenerateRandomEmail() string {

	firstPart := rand.Int()
	secondPart := uuid.New()
	domain, statusCode := getDomains()

	if statusCode != 200 {
		log.Panic("NO ES POSIBLE OBTENER LA URLS DE DOMINIOS")
	}

	email := fmt.Sprintf("%d-%s@%s", firstPart, secondPart.String(), domain ) 
	
	return email


}


func getDomains() (string, int) {

	domainsBinary, statusCode := httpclient.GetRequest("https://api.mail.tm/domains?page=1", "")
	if statusCode != 200 {
		log.Panic("NO ES POSIBLE OBTENER LA URLS DE DOMINIOS")
	}

	var domainsJSON []models.Domains
	_ = json.Unmarshal(domainsBinary, &domainsJSON)

	return domainsJSON[0].Domain, statusCode

}