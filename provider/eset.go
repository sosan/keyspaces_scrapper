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
		ButtonHome:                `[data-label="home-overview-empty-add-license-btn"]`,
		ButtonFreeLicense:         `#license-add-new-slides > div > ion-slide.LicenseForkSlide.md.swiper-slide.swiper-zoom-container.hydrated.swiper-slide-active > div > div > div:nth-child(3) > div > button`,
		ButtonContinueFreeLicense: `[data-label="license-fork-slide-continue-button"]`,
		ButtonWindows:             `#protect-choose-os-9`,
		ButtonContinueSelectOS:    `[data-label="device-protect-choose-platform-continue-btn"]`,
		EmailToShare:              `[data-label="device-protect-get-installer-email-input-input"]`,
		ButtonToSendEmailShare:    `[data-label="device-protect-get-installer-send-email-btn"]`,
		ButtonShowLicense:         `[data-label="license-list-open-detail-page-btn"]`,
		TextLicense:               `#main-content > div > div > div > ion-content > div.license-detail-content > ion-grid.license-detail-portal-content.license-detail-portal-content__bottom-grid.md.hydrated > ion-row > ion-col:nth-child(1) > div > div > div > div > ion-grid:nth-child(1) > ion-row:nth-child(2) > ion-col:nth-child(6) > div > div.DetailInfoSectionItem__value > p`,
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
		chromedp.WaitVisible(postData.ButtonHome, chromedp.BySearch),

		// pagina principal y click en el boton en medio
		chromedp.Sleep(2 * time.Second),
		chromedp.Click(postData.ButtonHome, chromedp.BySearch),
		chromedp.Sleep(2 * time.Second),

		// seleccionar boton prueba gratuita y continuar
		chromedp.Click(postData.ButtonFreeLicense, chromedp.ByQuery),
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(postData.ButtonContinueFreeLicense, chromedp.ByQuery),
		chromedp.Sleep(1 * time.Second),

		// seleccion windows y continue
		chromedp.Click(postData.ButtonWindows, chromedp.ByQuery),
		chromedp.Sleep(3 * time.Second),
		chromedp.Click(postData.ButtonContinueSelectOS, chromedp.ByQuery),

		// envio key al email
		chromedp.SendKeys(postData.EmailToShare, postData.EmailValue, chromedp.ByQuery),
		chromedp.Click(postData.ButtonToSendEmailShare, chromedp.ByQuery),

		// cortesia a esperar que recibimos el correo
		chromedp.Sleep(10 * time.Second),

		// ir licencias
		chromedp.Navigate("https://home.eset.com/subscriptions"), //  https://home.eset.com/licenses
		chromedp.Sleep(1 * time.Second),

		chromedp.Click(postData.ButtonShowLicense, chromedp.ByQuery),
		chromedp.Text(postData.TextLicense, licencia, chromedp.ByQuery),
	}

	return tasks
}
