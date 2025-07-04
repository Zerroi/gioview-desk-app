package views

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/oligo/gioview/page"
	"github.com/oligo/gioview/theme"
	"github.com/oligo/gioview/view"
)

var (
	ViewBId = view.NewViewID("ViewB")
	ViewCId = view.NewViewID("ViewC")
)

type ViewB struct {
	*view.BaseView
	page.PageStyle
}

func NewViewB() view.View {
	return &ViewB{BaseView: &view.BaseView{}}
}

func (v *ViewB) ID() view.ViewID {
	return ViewBId
}

func (v *ViewB) Title() string {
	return "View B"
}

func (v *ViewB) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return v.PageStyle.Layout(gtx, th,
		material.H4(th.Theme, "This is View B").Layout,
	)
}

type ViewC struct {
	*view.BaseView
	page.PageStyle
}

func NewViewC() view.View {
	return &ViewC{BaseView: &view.BaseView{}}
}

func (v *ViewC) ID() view.ViewID {
	return ViewCId
}

func (v *ViewC) Title() string {
	return "View C"
}

func (v *ViewC) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return v.PageStyle.Layout(gtx, th,
		material.H4(th.Theme, "This is View C").Layout,
	)
}
