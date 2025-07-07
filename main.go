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
	activeTheme *Theme // Points to the currently used theme
	vm          *HomeView
}

// Add a method to UI to toggle the theme
func (ui *UI) toggleTheme() {
	if ui.activeTheme == ui.darkTheme {
		ui.activeTheme = ui.lightTheme
	} else {
		ui.activeTheme = ui.darkTheme
	}
	// Invalidate the window to trigger a redraw with the new theme
	ui.window.Invalidate()
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
		// Pass the theme toggle function to the home view constructor
		ui.vm = newHome(ui.window, ui.toggleTheme)
	}

	return ui.vm.Layout(gtx, ui.activeTheme)
}

func main() {
	go func() {
		w := &app.Window{}
		w.Option(app.Decorated(true))
		appTitle := "Gio Desktop App"
		w.Option(app.Title(appTitle))

		darkTheme := NewDarkTheme()
		lightTheme := NewLightTheme()

		ui := &UI{
			window: w,
			// Set up both themes and activate one by default
			darkTheme:   darkTheme,
			lightTheme:  lightTheme,
			activeTheme: lightTheme, // Start with the light theme
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
