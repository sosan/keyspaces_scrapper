package provider

import (
	"context"
	"log"
	"main/models"
	"main/utils"
	"time"

	"github.com/chromedp/chromedp"
)

const (
	WEB_REGISTER_ESET = "https://login.eset.com/Register/Index"
	PASSWORD_DEFAULT  = "Kle_pto_r123#"
)

func RegisterAccount(email string) bool {

	isRegistered := sendChromeRequest(email)
	return isRegistered

}

func GetLicense(email string) (bool, string) {

	isConfirmed, licencia := sendChromeConfirm(email)
	return isConfirmed, licencia

}

func sendChromeRequest(email string) bool {

	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()

	postData := models.PostData{
		UriValue:               WEB_REGISTER_ESET,
		EmailValue:             email,
		PasswordValue:          PASSWORD_DEFAULT,
		SelectedCountryValue:   "206",
		WantReceiveNewsValue:   "true",
		WaitVisibleValue:       "#fullpage",
		EmailElement:           "#Email",
		PasswordElement:        "#Password",
		SelectedCountryElement: "#SelectedCountry",
		WantReceiveNewsElement: "#WantReceiveNews",
		SubmitElement:          "input.account__entry.btn.btn-normal",
		ScreenshotElement:      `#root > ion-app`,
	}

	var buf []byte
	err := chromedp.Run(ctx, submitRegisterESET(postData, &buf))
	return err == nil

}

func submitRegisterESET(postData models.PostData, buf *[]byte) chromedp.Tasks {

	// buf por si queremos realizar un screenshot

	tasks := chromedp.Tasks{
		chromedp.Navigate(postData.UriValue),
		chromedp.WaitVisible(postData.WaitVisibleValue),
		chromedp.Click("#cc-accept", chromedp.NodeVisible),
		chromedp.SendKeys(postData.EmailElement, postData.EmailValue, chromedp.ByID),
		chromedp.Submit(postData.SubmitElement, chromedp.BySearch),
		chromedp.Sleep(5 * time.Second),
		chromedp.SendKeys(postData.PasswordElement, postData.PasswordValue, chromedp.ByID),
		chromedp.SendKeys(postData.SelectedCountryElement, postData.SelectedCountryValue, chromedp.ByID),
		chromedp.SendKeys(postData.WantReceiveNewsElement, postData.WantReceiveNewsValue, chromedp.ByID),
		chromedp.Submit(postData.SubmitElement, chromedp.BySearch),
		chromedp.Sleep(12 * time.Second),
	}

	return tasks
}

func sendChromeConfirm(email string) (bool, string) {

	// opts := append(chromedp.DefaultExecAllocatorOptions[:],
	// 	chromedp.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3239.108 Safari/537.36"),
	// 	chromedp.NoDefaultBrowserCheck,
	// 	// chromedp.Headless,

	// )

	CON_NAVEGADOR := utils.GetEnv("CON_NAVEGADOR")

	opts := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36"),
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
	}

	if CON_NAVEGADOR == "false" {
		opts = append(opts, chromedp.Headless)
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithDebugf(log.Printf))
	defer cancel()

	postData := models.PostData{
		UriValue:               "https://home.eset.com",
		WaitVisibleValue:       "#fullpage",
		EmailValue:             email,
		PasswordValue:          PASSWORD_DEFAULT,
		EmailElement:           "Email",
		PasswordElement:        "Password",
		SubmitElement:          "input.account__entry.btn.btn-normal",
		ButtonHome:             `#main-content > div > ion-tabs > div > ion-router-outlet > div > ion-content > div.home-overview-section > div.home-overview-section__cards-wrap > ion-card > ion-button`,
		ButtonFreeLicense:      `#license-add-new-slide > form > div.ion-cui-btn-above-link-wrapper > ion-button:nth-child(3)`,
		ButtonWindows:          `#protect-choose-os-9`,
		ButtonContinue:         `#main-content > div > ion-tabs > div > ion-router-outlet > div > ion-content > div > ion-button`,
		EmailToShare:           `#main-content > div > ion-tabs > div > ion-router-outlet > div > ion-content > div > ion-row > form > div > div.ion-cui-form-field > ion-item > div > ion-input > input`,
		ButtonToSendEmailShare: `#main-content > div > ion-tabs > div > ion-router-outlet > div > ion-content > div > ion-row > form > ion-button`,
		ButtonShowLicense:      `#license-list-large-previews > ion-row > ion-col > ion-card > ion-button`,
		TextLicense:            `#main-content > div > ion-tabs > div > ion-router-outlet > div > ion-content > div.license-detail-content > ion-grid.license-detail-portal-content.license-detail-portal-content__bottom-grid.md.hydrated > ion-row > ion-col:nth-child(1) > div > ion-card:nth-child(1) > ion-grid > ion-row:nth-child(2) > ion-col:nth-child(6) > div > div > p.DetailInfoSectionItem__value > ion-text`,
	}

	var buf []byte
	var licencia string
	err := chromedp.Run(ctx, submitConfirmAccount(postData, &buf, &licencia))

	// if err := os.WriteFile("screenshot_confirmacion.png", buf, 0o644); err != nil {
	// 	log.Fatal(err)
	// }

	return err == nil, licencia

}

func submitConfirmAccount(postData models.PostData, buf *[]byte, licencia *string) chromedp.Tasks {

	tasks := chromedp.Tasks{
		chromedp.Navigate(postData.UriValue),

		// enviar formulario login
		chromedp.SendKeys(postData.EmailElement, postData.EmailValue, chromedp.ByID),
		chromedp.SendKeys(postData.PasswordElement, postData.PasswordValue, chromedp.ByID),
		chromedp.Click("cc-accept", chromedp.ByID),
		chromedp.Submit(postData.SubmitElement, chromedp.BySearch),
		chromedp.WaitVisible(postData.ButtonHome, chromedp.ByQuery),

		// pagina principal y click en el boton en medio
		chromedp.Sleep(10 * time.Second),
		chromedp.Click(postData.ButtonHome, chromedp.ByQuery),
		chromedp.Sleep(10 * time.Second),

		// click en modal
		chromedp.Click(postData.ButtonFreeLicense, chromedp.ByQuery),

		// seleccion windows y continue
		chromedp.Sleep(5 * time.Second),
		chromedp.Click(postData.ButtonWindows, chromedp.ByQuery),
		chromedp.Sleep(2 * time.Second),
		chromedp.Click(postData.ButtonContinue, chromedp.ByQuery),

		// envio key al email
		chromedp.SendKeys(postData.EmailToShare, postData.EmailValue, chromedp.ByQuery),
		chromedp.Click(postData.ButtonToSendEmailShare, chromedp.ByQuery),

		// cortesia a esperar que recibimos el correo
		chromedp.Sleep(60 * time.Second),

		// ir licencias
		chromedp.Navigate("https://home.eset.com/licenses"),
		chromedp.Sleep(5 * time.Second),

		chromedp.Click(postData.ButtonShowLicense, chromedp.ByQuery),
		chromedp.Text(postData.TextLicense, licencia, chromedp.ByQuery),
	}

	return tasks
}
