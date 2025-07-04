package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget" // <--- 已修正：确保导入了正确的 widget 包
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
	sidebar *NavDrawer
	tabbar  *navi.Tabbar
}

// ID returns the view's unique identifier.
func (hv *HomeView) ID() string {
	return "Home"
}

// Layout renders the main view.
func (hv *HomeView) Layout(gtx C, th *theme.Theme) layout.Dimensions {
	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Start,
	}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return NaviDrawerStyle{
				NavDrawer: hv.sidebar,
				Inset: layout.Inset{
					Top:    unit.Dp(20),
					Bottom: unit.Dp(20),
					Left:   unit.Dp(2),
				},
				Bg:    th.Bg2,
				Width: unit.Dp(200),
			}.Layout(gtx, th)
		}),
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

func newHome(window *app.Window) *HomeView {
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
	}
}
