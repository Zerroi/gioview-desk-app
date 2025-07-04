package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/oligo/gioview/view"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

// UI holds all of the application's state
type UI struct {
	window *app.Window
	theme  *Theme
	vm     *HomeView
}

// Loop runs the application's main event loop.
func (ui *UI) Loop() error {
	var ops op.Ops
	for {
		e := ui.window.Event()
		switch e := e.(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			ui.layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}

// layout renders the application's UI.
func (ui *UI) layout(gtx C) D {
	if ui.vm == nil {
		ui.vm = newHome(ui.window)
	}

	return ui.vm.Layout(gtx, ui.theme)
}

func main() {
	go func() {
		w := &app.Window{}
		w.Option(app.Title("Gio Desktop App"))
		th := NewTheme()
		th.TextSize = unit.Sp(14)
		th.Bg2 = color.NRGBA{R: 0x2e, G: 0x34, B: 0x40, A: 255}
		th.Bg = color.NRGBA{R: 0x24, G: 0x28, B: 0x32, A: 255}
		th.Fg = color.NRGBA{R: 0xe5, G: 0xe9, B: 0xf0, A: 255}
		th.ContrastBg = color.NRGBA{R: 0x88, G: 0xc0, B: 0xd0, A: 255}
		th.ContrastFg = color.NRGBA{R: 0x24, G: 0x28, B: 0x32, A: 255}

		ui := &UI{theme: th, window: w}
		if err := ui.Loop(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

var (
	ViewA = view.NewViewID("ViewA")
	ViewB = view.NewViewID("ViewB")
	ViewC = view.NewViewID("ViewC")
)
