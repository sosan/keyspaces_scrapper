package email

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"main/httpclient"
	"main/models"
	"main/utils"
)

const (
	API_MAIL         = "https://api.mail.tm"
	DEFAULT_PASSWORD = "string123_123"
)

func CreateEmail() (bool, string) {

	currentEmail := utils.GenerateRandomEmail()

	postData, _ := json.Marshal(map[string]string{
		"address":  currentEmail,
		"password": DEFAULT_PASSWORD,
	})

	uri := fmt.Sprintf("%s/accounts", API_MAIL)
	body, statusCode := httpclient.GeneratePostJsonRequest(uri, postData)

	if statusCode == 201 {
		fmt.Println("creado correctamente")
	} else {
		log.Printf("NO creado correctamente | StatusCode %d | Body %s", statusCode, body)

		return false, ""
	}

	return true, currentEmail

}

func GetAllEmails(email string) (bool, []models.ResumenEmail) {

	isToeknGetted, token := getMailToken(email, DEFAULT_PASSWORD)
	if !isToeknGetted {
		return false, nil
	}

	uri := fmt.Sprintf("%s/messages", API_MAIL)
	bodyRaw, statusCode := httpclient.GenerateGetRequest(uri, token)

	if statusCode != 200 {
		return false, nil
	}

	var emails []models.ResumenEmail
	json.Unmarshal(bodyRaw, &emails)

	return true, emails

}

func GetEmailById(id string, email string) (bool, models.EmailCompleto) {

	var emailsComplete models.EmailCompleto
	isToeknGetted, token := getMailToken(email, DEFAULT_PASSWORD)
	if !isToeknGetted {
		return false, emailsComplete
	}

	uri := fmt.Sprintf("%s/messages/%s", API_MAIL, id)
	bodyRaw, statusCode := httpclient.GenerateGetRequest(uri, token)

	if statusCode != 200 {
		return false, emailsComplete
	}

	json.Unmarshal(bodyRaw, &emailsComplete)

	return true, emailsComplete

}

func GetLinkFromEmail(textEmail string) string {

	firstPart := strings.Split(textEmail, "[")

	link := strings.Split(firstPart[1], "]")[0]

	return link

}

func getMailToken(email string, password string) (bool, string) {

	uri := fmt.Sprintf("%s/token", API_MAIL)
	postData, _ := json.Marshal(map[string]string{
		"address":  email,
		"password": DEFAULT_PASSWORD,
	})

	bodyRaw, statusCode := httpclient.GeneratePostJsonRequest(uri, postData)
	var jsonResponse *models.Token
	json.Unmarshal(bodyRaw, &jsonResponse)

	if statusCode != 200 {
		log.Printf("NO creado correctamente | StatusCode %d | Body %s ", statusCode, bodyRaw)
		return false, ""
	}

	return true, jsonResponse.Token

}

