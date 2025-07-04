package main

import (
	"image/color"
	"log"

	"gioui.org/font/gofont"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/oligo/gioview/theme"
)

type Theme = theme.Theme

func NewTheme() *Theme {
	t := theme.NewTheme("", nil, false)
	// 使用 text.NewShaper 并通过 text.WithCollection 提供字体
	t.Shaper = text.NewShaper(text.WithCollection(gofont.Collection())) // <--- 修正此行
	t.Theme = material.NewTheme()
	t.Theme.Shaper = t.Shaper

	// A modern, code-editor-like color scheme
	t.Palette = material.Palette{
		Bg:         color.NRGBA{R: 0x24, G: 0x28, B: 0x32, A: 255}, // Dark background
		Fg:         color.NRGBA{R: 0xe5, G: 0xe9, B: 0xf0, A: 255}, // Light foreground
		ContrastBg: color.NRGBA{R: 0x88, G: 0xc0, B: 0xd0, A: 255}, // A soft blue for accents
		ContrastFg: color.NRGBA{R: 0x24, G: 0x28, B: 0x32, A: 255}, // Dark text on accents
	}
	t.Bg2 = color.NRGBA{R: 0x2e, G: 0x34, B: 0x40, A: 255} // Slightly lighter background for sidebars/panels
	t.HoverAlpha = 48
	t.SelectedAlpha = 96
	log.Println("Theme loaded")
	return t
}
