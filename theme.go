package main

import (
	"gioui.org/unit"
	"image/color"
	"log"

	"gioui.org/font/gofont"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/oligo/gioview/theme"
)

type Theme = theme.Theme

// NewDarkTheme 创建一个深色主题（原 NewTheme 函数）
func NewDarkTheme() *Theme {
	t := theme.NewTheme("", nil, false)
	t.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	t.Theme = material.NewTheme()
	t.Theme.Shaper = t.Shaper

	t.TextSize = unit.Sp(14)
	t.Bg = color.NRGBA{R: 0x24, G: 0x28, B: 0x32, A: 255}
	t.Bg2 = color.NRGBA{R: 0x2e, G: 0x34, B: 0x40, A: 255}
	t.Fg = color.NRGBA{R: 0xe5, G: 0xe9, B: 0xf0, A: 255}
	t.ContrastBg = color.NRGBA{R: 0x88, G: 0xc0, B: 0xd0, A: 255}
	t.ContrastFg = color.NRGBA{R: 0x24, G: 0x28, B: 0x32, A: 255}
	t.HoverAlpha = 48
	t.SelectedAlpha = 96
	log.Println("Dark theme loaded")
	return t
}

// NewLightTheme 创建一个浅色主题
func NewLightTheme() *Theme {
	t := theme.NewTheme("", nil, false)
	t.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	t.Theme = material.NewTheme()
	t.Theme.Shaper = t.Shaper

	t.TextSize = unit.Sp(14)
	t.Bg = color.NRGBA{R: 0xfa, G: 0xfa, B: 0xfa, A: 255}         // 浅灰背景
	t.Bg2 = color.NRGBA{R: 0xf0, G: 0xf0, B: 0xf0, A: 255}        // 稍深的灰色用于侧边栏/标题栏
	t.Fg = color.NRGBA{R: 0x21, G: 0x21, B: 0x21, A: 255}         // 深色文字
	t.ContrastBg = color.NRGBA{R: 0x03, G: 0xa9, B: 0xf4, A: 255} // 亮蓝色作为对比色
	t.ContrastFg = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 255} // 对比色上的白色文字
	t.HoverAlpha = 48
	t.SelectedAlpha = 96
	log.Println("Light theme loaded")
	return t
}
