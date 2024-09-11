package main

import (
	"fmt"
	"log"
	"main/email"
	"main/httpclient"
	"main/provider"
	"main/update"
	"main/utils"
)

func init() {
	fmt.Println(update.GetVersionBuild())
	utils.LoadEnvs()
	if utils.GetEnv("DEBUG") == "true" {
		return
	}
	needUpdate, updatedOk := update.AutoUpdate()
	if needUpdate && !updatedOk {
		log.Printf("NECESARIO ACTUALIZAR PERO NO HA SIDO POSIBLE!!")
	}
}

func main() {

	log.Printf("Creando email")
	isEmailCreated, emailGenerated := email.CreateEmail()

	if !isEmailCreated {
		return
	}

	log.Printf("Email %s creado", emailGenerated)
	log.Printf("Registrando cuenta en ESET")

	accountRegistered := provider.RegisterAccount(emailGenerated)
	if !accountRegistered {
		return
	}

	log.Printf("Cuenta registrada en ESET")
	emailsCanBeObtened, emails := email.GetAllEmails(emailGenerated)

	if !emailsCanBeObtened {
		return
	}

	if len(emails) < 1 {
		return
	}

	// obtenemos el ultimo email
	emailObtened, emailComplete := email.GetEmailById(emails[0].ID, emailGenerated)
	if !emailObtened {
		return
	}

	uriLink := email.GetLinkFromEmail(emailComplete.Text)
	log.Printf("Link para confirmar cuenta: %s", uriLink)

	log.Printf("Confirmando cuenta...")

	confirmed := httpclient.ConfirmAccount(uriLink, emailGenerated)
	if !confirmed {
		log.Printf("Error al confirmar link de la cuenta")
		return
	}

	confirmed, licencia := provider.GetLicense(emailGenerated)
	if !confirmed {
		log.Printf("Error al confirmar cuenta")
		return
	}

	log.Printf("Cuenta confirmada con el correo %s", emailGenerated)
	log.Printf("Obteniendo licencia para %s", emailGenerated)

	log.Printf("---------------------- LICENCIA %s", licencia)

	utils.SaveLicence(licencia)

}
