package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("03:04:05 PM")
	clock.SetText(formatted)
}

func main() {
	// Create a new application
	myApp := app.New()
	myWindow := myApp.NewWindow("Go Clock")

	// Set window size and make it non-resizable
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.SetFixedSize(true)

	// Create a label for the time
	timeLabel := widget.NewLabel("")
	timeLabel.TextStyle = fyne.TextStyle{Bold: true}
	timeLabel.Alignment = fyne.TextAlignCenter
	timeLabel.TextSize = 48

	// Create a label for the date
	dateLabel := widget.NewLabel("")
	dateLabel.TextSize = 18
	dateLabel.Alignment = fyne.TextAlignCenter

	// Update time immediately
	updateTime(timeLabel)
	dateLabel.SetText(time.Now().Format("Monday, January 2, 2006"))

	// Update time every second
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			case <-time.After(time.Second):
				updateTime(timeLabel)
				dateLabel.SetText(time.Now().Format("Monday, January 2, 2006"))
			}
		}
	}()

	// Create a container with the time and date
	content := container.NewVBox(
		layout.NewSpacer(),
		timeLabel,
		dateLabel,
		layout.NewSpacer(),
	)

	// Set the content and show the window
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
	close(quit)
}
