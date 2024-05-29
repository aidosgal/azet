package gui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/canvas"
    "image/color"
    "github.com/aidosgal/azet/internal/speech"
)

func makeBanner() fyne.CanvasObject {
    toolbar := widget.NewLabel("Galimzhan Aidos, Kenzheakhmet Beksultan | CS-2203")
    toolbar.Alignment = fyne.TextAlignCenter
    return toolbar
}

func makeContent() fyne.CanvasObject {
    white := color.White

    speechText := canvas.NewText("Я слушаю вас", white)
    speechText.TextSize = 20
    speechText.Alignment = fyne.TextAlignCenter

    textToSpeechButton := widget.NewButton("Говорить", func() {
      text := speech.SpeechToText()
      _ = text
    })

    entry := widget.NewEntry()
    entry.PlaceHolder = "Введите текст"

    sendButton := widget.NewButton("Отправить", func() {})

    inputBox := container.NewHBox(
        entry,
        sendButton,
    )
    _ = inputBox

    content := container.NewVBox(
      speechText,
      textToSpeechButton,
      entry,
    )

    return content
}

func MakeGUI() fyne.CanvasObject {
    left := widget.NewLabel("left")
    right := widget.NewLabel("right")

    return container.NewBorder(makeBanner(), nil, left, right, makeContent())
}
