package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/oligo/gioview/view"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

// UI holds all of the application's state
type UI struct {
	window      *app.Window
	darkTheme   *Theme
	lightTheme  *Theme
	activeTheme *Theme // 指向当前使用的主题
	vm          *HomeView
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

	return ui.vm.Layout(gtx, ui.activeTheme)
}

func main() {
	go func() {
		w := &app.Window{}
		w.Option(app.Decorated(true))
		appTitle := "Gio Desktop App"
		w.Option(app.Title(appTitle))

		// --- 已修正: 初始化两个主题 ---
		darkTheme := NewDarkTheme()
		lightTheme := NewLightTheme()

		ui := &UI{
			window: w,
			// 设置主题，并默认激活深色主题
			darkTheme:   darkTheme,
			lightTheme:  lightTheme,
			activeTheme: lightTheme,
		}

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
