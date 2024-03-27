package provider

import (
	"context"
	"log"
	"main/models"
	"main/utils"
	"os"
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
	CON_NAVEGADOR := utils.GetEnv("CON_NAVEGADOR")

	opts := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36"),
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
	}

	if CON_NAVEGADOR == "false" || CON_NAVEGADOR == "" {
		opts = append(opts, chromedp.Headless)
	}

	actx, acancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer acancel()

	ctx, cancel := chromedp.NewContext(actx, chromedp.WithDebugf(log.Printf))
	defer cancel()

	postData := models.PostData{
		UriValue:               WEB_REGISTER_ESET,
		EmailElement:           "#email",
		EmailValue:             email,
		PasswordElement:        "#password",
		PasswordValue:          PASSWORD_DEFAULT,
		WantReceiveNewsElement: "#ReceiveNewsCheckbox-input",
		WantReceiveNewsValue:   "true",
		SubmitElement:          `button[type="submit"]`,
		ScreenshotElement:      `#root > ion-app`,
		WaitVisibleValue:       "#root",
		// SelectedCountryElement: "#SelectedCountry",
		// SelectedCountryValue:   "206",
	}

	var buf []byte
	err := chromedp.Run(ctx, submitRegisterESET(postData, &buf))

	if utils.GetEnv("SCREENSHOT") == "true" {
		if err := os.WriteFile("fullScreenshot.png", buf, 0o644); err != nil {
			log.Fatal(err)
		}
	}

	return err == nil

}

func submitRegisterESET(postData models.PostData, buf *[]byte) chromedp.Tasks {

	// buf por si queremos realizar un screenshot

	tasks := chromedp.Tasks{
		chromedp.Navigate(postData.UriValue),
		chromedp.WaitVisible(postData.WaitVisibleValue),
		chromedp.Click("#cc-accept", chromedp.NodeVisible),
		chromedp.SendKeys(postData.EmailElement, postData.EmailValue, chromedp.ByID),
		chromedp.Click(postData.SubmitElement, chromedp.BySearch),
		chromedp.Sleep(5 * time.Second),
		// second level
		chromedp.SendKeys(postData.PasswordElement, postData.PasswordValue, chromedp.ByID),
		// chromedp.SendKeys(postData.SelectedCountryElement, postData.SelectedCountryValue, chromedp.ByID),
		chromedp.SendKeys(postData.WantReceiveNewsElement, postData.WantReceiveNewsValue, chromedp.ByID),
		chromedp.Click(postData.SubmitElement, chromedp.BySearch),
		chromedp.Sleep(5 * time.Second),
	}

	if utils.GetEnv("SCREENSHOT") == "true" {
		tasks = append(tasks, chromedp.Screenshot("#root", buf))
	}

	return tasks
}

func sendChromeConfirm(email string) (bool, string) {

	opts := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36"),
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
	}

	// interesa que aparezca el navegador
	// CON_NAVEGADOR := utils.GetEnv("CON_NAVEGADOR")
	// if CON_NAVEGADOR == "false" || CON_NAVEGADOR == "" {
	// 	opts = append(opts, chromedp.Headless)
	// }

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithDebugf(log.Printf))
	defer cancel()
	// use copyselector in devtools chromium
	postData := models.PostData{
		UriValue:                  "https://home.eset.com",
		WaitVisibleValue:          "#root",
		EmailElement:              "#email",
		EmailValue:                email,
		PasswordElement:           "#password",
		PasswordValue:             PASSWORD_DEFAULT,
		SubmitElement:             `button[type="submit"]`,
		ButtonHome: `#main-content > div > div > div > ion-content > div.home-overview-section > div.home-overview-section__cards-wrap > ion-card > ion-button`,
		ButtonFreeLicense:         `#license-add-new-slides > div > ion-slide.LicenseForkSlide.md.swiper-slide.swiper-zoom-container.hydrated.swiper-slide-active > div > div > ion-card:nth-child(3)`,
		ButtonContinueFreeLicense: `#license-add-new-slides > div > ion-slide.LicenseForkSlide.md.swiper-slide.swiper-zoom-container.hydrated.swiper-slide-active > div > ion-button`,
		ButtonWindows:             `#protect-choose-os-9`,
		
		ButtonContinueSelectOS:    `#main-content > div > div > div > ion-content > div.protect-page-container--content > div.protect-page--button-container > ion-button.ion-cui-button.protect-page--continue-button.ion-color.ion-color-secondary.ios.button.button-block.button-solid.ion-activatable`,
		
		EmailToShare:              `#main-content > div > div > div > ion-content > div.protect-page-container--content.ProtectGetInstaller__content > ion-row > form > div > div.ion-cui-form-field > ion-item > div > ion-input > input`,
		
		// `#main-content > div > ion-tabs > div > ion-router-outlet > div > ion-content > div > ion-row > form > div > div.ion-cui-form-field > ion-item > div > ion-input > input`,
		ButtonToSendEmailShare:    `#main-content > div > div > div > ion-content > div.protect-page-container--content.ProtectGetInstaller__content > ion-row > form > ion-button`,
		// `#main-content > div > ion-tabs > div > ion-router-outlet > div > ion-content > div > ion-row > form > ion-button`,
		ButtonShowLicense:         `#license-list-large-previews > ion-row > ion-col > ion-card > ion-button`,
		TextLicense:               `#main-content > div > div > div > ion-content > div.license-detail-content > ion-grid.license-detail-portal-content.license-detail-portal-content__bottom-grid.md.hydrated > ion-row > ion-col:nth-child(1) > div > ion-card > ion-grid:nth-child(1) > ion-row:nth-child(2) > ion-col:nth-child(6) > div > p.DetailInfoSectionItem__value > ion-text`,
		// `#main-content > div > ion-tabs > div > ion-router-outlet > div > ion-content > div.license-detail-content > ion-grid.license-detail-portal-content.license-detail-portal-content__bottom-grid.md.hydrated > ion-row > ion-col:nth-child(1) > div > ion-card > ion-grid:nth-child(1) > ion-row:nth-child(2) > ion-col:nth-child(6) > div > p.DetailInfoSectionItem__value > ion-text`,
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
		chromedp.WaitVisible(postData.WaitVisibleValue, chromedp.ByQuery),

		// enviar formulario login
		chromedp.SendKeys(postData.EmailElement, postData.EmailValue, chromedp.ByID),
		chromedp.SendKeys(postData.PasswordElement, postData.PasswordValue, chromedp.ByID),
		chromedp.Click("cc-accept", chromedp.ByID),
		chromedp.Click(postData.SubmitElement, chromedp.BySearch),
		chromedp.WaitVisible(postData.ButtonHome, chromedp.ByQuery),

		// pagina principal y click en el boton en medio
		chromedp.Sleep(10 * time.Second),
		chromedp.Click(postData.ButtonHome, chromedp.ByQuery),
		chromedp.Sleep(10 * time.Second),

		// seleccionar boton prueba gratuita y continuar
		chromedp.Click(postData.ButtonFreeLicense, chromedp.ByQuery),
		chromedp.Sleep(5 * time.Second),
		chromedp.Click(postData.ButtonContinueFreeLicense, chromedp.ByQuery),
		chromedp.Sleep(5 * time.Second),

		// seleccion windows y continue
		chromedp.Click(postData.ButtonWindows, chromedp.ByQuery),
		chromedp.Sleep(10 * time.Second),
		chromedp.Click(postData.ButtonContinueSelectOS, chromedp.ByQuery),

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
