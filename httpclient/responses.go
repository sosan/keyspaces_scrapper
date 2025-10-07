package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func GeneratePostJsonRequest(uri string, postData []byte) ([]byte, int) {

	clientHttp := http.Client{Timeout: time.Duration(60) * time.Second}

	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(postData))
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", "api_anti_vir")
	req.Header.Add("Content-type", "application/json")

	resp, err := clientHttp.Do(req)

	if err != nil {
		log.Fatal(err)
		return nil, 404
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, 404
	}

	return body, resp.StatusCode

}

func GenerateGetRequest(uri string, token string) ([]byte, int) {

	clientHttp := http.Client{Timeout: time.Duration(60) * time.Second}

	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-type", "application/json")

	resp, err := clientHttp.Do(req)

	if err != nil {
		log.Fatal(err)
		return nil, 404
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, 404
	}

	return body, resp.StatusCode

}

func ConfirmAccount(uri string, email string) bool {

	clientHttp := http.Client{Timeout: time.Duration(60) * time.Second}

	req, _ := http.NewRequest("GET", uri, nil)

	resp, err := clientHttp.Do(req)

	if err != nil {
		log.Fatal(err)
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == 200

}

func GetRequest(uri string, token string) ([]byte, int) {

	clientHttp := http.Client{Timeout: time.Duration(60) * time.Second}

	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Accept", `application/json`)
	if token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	req.Header.Add("Content-type", "application/json")

	resp, err := clientHttp.Do(req)

	if err != nil {
		log.Fatal(err)
		return nil, 404
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, 404
	}

	return body, resp.StatusCode

}

func GetRequestRaw(uri string, token string) (*io.ReadCloser, int) {
	clientHttp := http.Client{Timeout: time.Duration(60) * time.Second}

	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Accept", `application/json`)
	if token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	req.Header.Add("Content-type", "application/json")

	resp, err := clientHttp.Do(req)

	if err != nil {
		log.Fatal(err)
		return nil, 404
	}
	// defer resp.Body.Close()

	return &resp.Body, resp.StatusCode

}