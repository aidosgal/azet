package main

import (
	"fyne.io/fyne/v2/app"
	_ "fyne.io/fyne/v2/container"
	_ "fyne.io/fyne/v2/widget"
  "github.com/aidosgal/azet/internal/gui"
  "github.com/aidosgal/azet/internal/theme"
)

func main() {
	app := app.New()
	window := app.NewWindow("Hello")

  app.Settings().SetTheme(theme.NewTheme())

  window.SetContent(gui.MakeGUI())

	window.ShowAndRun()
}
