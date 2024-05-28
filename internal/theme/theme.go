package theme

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/theme"
  "image/color"
)

type Theme struct {
  fyne.Theme
}

func NewTheme() fyne.Theme {
  return &Theme{Theme: theme.DefaultTheme()}
}

func (t *Theme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
  return t.Theme.Color(name, variant)
}

func (t *Theme) Size(name fyne.ThemeSizeName) float32 {
  if name == theme.SizeNamePadding {
    return 12
  }

  return t.Theme.Size(name)
}
