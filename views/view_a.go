package views

import (
	"log"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/oligo/gioview/page"
	"github.com/oligo/gioview/theme"
	"github.com/oligo/gioview/view"
	gvw "github.com/oligo/gioview/widget"
)

var (
	ViewAId = view.NewViewID("ViewA")
)

type ViewA struct {
	*view.BaseView
	page.PageStyle
	input      gvw.TextField
	button     widget.Clickable
	buttonText string
}

func NewViewA() view.View {
	return &ViewA{
		BaseView:   &view.BaseView{},
		buttonText: "Click Me",
	}
}

func (v *ViewA) ID() view.ViewID {
	return ViewAId
}

func (v *ViewA) Title() string {
	return "View A"
}

func (v *ViewA) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if v.button.Clicked(gtx) {
		log.Printf("Button clicked! Input text: %s", v.input.Text())
		v.buttonText = "Clicked!"
	}

	return v.PageStyle.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(material.H4(th.Theme, "This is View A").Layout),
			layout.Rigid(layout.Spacer{Height: th.FingerSize}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				v.input.HelperText = "Type something..."
				return v.input.Layout(gtx, th, "Input Label")
			}),
			layout.Rigid(layout.Spacer{Height: th.FingerSize}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th.Theme, &v.button, v.buttonText)
				return btn.Layout(gtx)
			}),
		)
	})
}
