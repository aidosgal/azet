package gui

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/theme"
)

func MakeGUI() fyne.CanvasObject {
  toolbar := widget.NewToolbar(
    widget.NewToolbarAction(theme.HomeIcon(), func() {}),
  )
  
  left := widget.NewLabel("left")

  right := widget.NewLabel("right")

  content := widget.NewLabel("home")

  return container.NewBorder(toolbar, nil, left, right, content)
}
