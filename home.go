package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material" // Import material for the button
	"gioview-desk-app/views"
	"github.com/oligo/gioview/navi"
	"github.com/oligo/gioview/theme"
	"github.com/oligo/gioview/view"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

var (
	viewIcon, _ = widget.NewIcon(icons.ActionViewModule)
)

// HomeView is the main view of the application.
type HomeView struct {
	view.ViewManager
	sidebar     *NavDrawer
	tabbar      *navi.Tabbar
	toggleTheme func()           // Function to call for toggling
	themeButton widget.Clickable // state for the button
}

// ID returns the view's unique identifier.
func (hv *HomeView) ID() string {
	return "Home"
}

// Layout renders the main view.
func (hv *HomeView) Layout(gtx C, th *theme.Theme) layout.Dimensions {
	// Process button clicks
	if hv.themeButton.Clicked(gtx) {
		if hv.toggleTheme != nil {
			hv.toggleTheme()
		}
	}

	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Start,
	}.Layout(gtx,
		// Define the sidebar region
		layout.Rigid(func(gtx C) D {
			sidebarWidth := unit.Dp(200)
			gtx.Constraints.Max.X = gtx.Dp(sidebarWidth)
			gtx.Constraints.Min.X = gtx.Constraints.Max.X

			// Paint the background for the entire sidebar
			paint.FillShape(gtx.Ops, th.Bg2, clip.Rect{Max: gtx.Constraints.Max}.Op())

			// Use a vertical flex to stack the nav list and the theme button
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				// The navigation list should expand to fill the available vertical space
				layout.Flexed(1, func(gtx C) D {
					// Apply the original padding for the list
					return layout.Inset{
						Top:    unit.Dp(20),
						Bottom: unit.Dp(20),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return hv.sidebar.Layout(gtx, th)
					})
				}),
				// The theme toggle button is rigid at the bottom
				layout.Rigid(func(gtx C) D {
					// Add padding around the button
					return layout.UniformInset(unit.Dp(12)).Layout(gtx, func(gtx C) D {
						btn := material.Button(th.Theme, &hv.themeButton, "Toggle Theme")
						btn.Background = th.ContrastBg
						btn.Color = th.ContrastFg
						btn.Inset = layout.UniformInset(unit.Dp(8))
						return btn.Layout(gtx)
					})
				}),
			)
		}),
		// This is the main content area, which remains unchanged
		layout.Flexed(1, func(gtx C) D {
			gtx.Constraints.Min = gtx.Constraints.Max
			rect := clip.Rect{Max: gtx.Constraints.Max}
			paint.FillShape(gtx.Ops, th.Bg, rect.Op())

			return layout.Flex{
				Axis:      layout.Vertical,
				Alignment: layout.Middle,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return hv.tabbar.Layout(gtx, th)
				}),
				layout.Rigid(func(gtx C) D {
					return layout.Spacer{Height: unit.Dp(1)}.Layout(gtx)
				}),
				layout.Flexed(1, func(gtx C) D {
					if hv.CurrentView() == nil {
						return view.EmptyView{}.Layout(gtx, th)
					}
					return hv.CurrentView().Layout(gtx, th)
				}),
			)
		}),
	)
}

// Update the constructor to accept the toggle function
func newHome(window *app.Window, toggleTheme func()) *HomeView {
	vm := view.DefaultViewManager(window)

	sidebar := NewNavDrawer(vm)
	sidebar.AddSection(SimpleItemSection(viewIcon, "View A", func(item *navi.NavTree) {
		sidebar.OnItemSelected(item)
		vm.RequestSwitch(view.Intent{Target: ViewA})
	}))
	sidebar.AddSection(SimpleItemSection(viewIcon, "View B", func(item *navi.NavTree) {
		sidebar.OnItemSelected(item)
		vm.RequestSwitch(view.Intent{Target: ViewB})
	}))
	sidebar.AddSection(SimpleItemSection(viewIcon, "View C", func(item *navi.NavTree) {
		sidebar.OnItemSelected(item)
		vm.RequestSwitch(view.Intent{Target: ViewC})
	}))

	// Register views
	vm.Register(ViewA, views.NewViewA)
	vm.Register(ViewB, views.NewViewB)
	vm.Register(ViewC, views.NewViewC)

	// Set initial view
	vm.RequestSwitch(view.Intent{Target: ViewA})

	return &HomeView{
		ViewManager: vm,
		tabbar:      navi.NewTabbar(vm, &navi.TabbarOptions{MaxVisibleActions: 4}),
		sidebar:     sidebar,
		toggleTheme: toggleTheme, // Store the passed-in function
	}
}
