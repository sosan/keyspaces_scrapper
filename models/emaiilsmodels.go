package models

import (
	"time"
)

type Token struct {
	Token string `json:"token"`
}

type PostData struct {
	UriValue             string
	EmailValue           string
	PasswordValue        string
	SelectedCountryValue string
	WantReceiveNewsValue string
	WaitVisibleValue     string

	EmailElement           string
	PasswordElement        string
	SelectedCountryElement string
	WantReceiveNewsElement string
	SubmitElement          string
	ScreenshotElement      string

	ButtonHome             string
	ButtonFreeLicense      string
	ButtonWindows          string
	ButtonContinue         string
	ButtonContinueSelectOS string
	EmailToShare           string
	ButtonToSendEmailShare string
	ButtonShowLicense      string
	TextLicense            string
}

type ResumenEmail struct {
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	Msgid     string `json:"msgid"`
	From      struct {
		Address string `json:"address"`
		Name    string `json:"name"`
	} `json:"from"`
	To []struct {
		Address string `json:"address"`
		Name    string `json:"name"`
	} `json:"to"`
	Subject        string    `json:"subject"`
	Intro          string    `json:"intro"`
	Seen           bool      `json:"seen"`
	IsDeleted      bool      `json:"isDeleted"`
	HasAttachments bool      `json:"hasAttachments"`
	Size           int       `json:"size"`
	DownloadURL    string    `json:"downloadUrl"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type EmailCompleto struct {
	Id      string `json:"id"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
}

type Domains struct {
	Id        string `json:"id"`
	ReadOnly  bool   `json:"readonly"`
	Domain    string `json:"domain"`
	IsActive  bool   `json:"isactive"`
	IsPrivate bool   `json:"isprivate"`
	CreatedAt string `json:"createdat"`
	UpdatedAt string `json:"updatedat"`
}
