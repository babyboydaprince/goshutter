package app

import (
	"image"
	"image/png"
	"log"
	"os"

	"github.com/getlantern/systray"
	"github.com/kbinani/screenshot"
	"github.com/sqweek/dialog"
)

func main() {
	// Run the systray application
	systray.Run(onReady, onExit)
}

func onReady() {
	// Set up the system tray icon
	systray.SetIcon(getIcon())
	systray.SetTitle("Go Shutter")
	systray.SetTooltip("Take Screenshots Easily")

	// Create "Take Screenshot" menu item
	mScreenshot := systray.AddMenuItem("Take Screenshot", "Start screenshot mode")

	// Create "Quit" menu item
	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	go func() {
		for {
			select {
			case <-mScreenshot.ClickedCh:
				takeScreenshot()
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// Cleanup tasks if necessary when the app exits
}

func takeScreenshot() {
	// Get the bounds of the primary monitor
	bounds := screenshot.GetDisplayBounds(0)

	// Capture the screenshot
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Fatal(err)
	}

	// Prompt the user to save the file
	filePath, err := dialog.File().Title("Save Screenshot").Filter("PNG file", "png").Save()
	if err != nil {
		return // User canceled
	}

	// Save the file
	saveImage(img, filePath)
}

func saveImage(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Encode the image to PNG
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
}

func getIcon() []byte {
	// Replace with the path to your icon
	iconFile := "./img/logo.png"
	iconData, err := os.ReadFile(iconFile)
	if err != nil {
		log.Fatal(err)
	}
	return iconData
}
