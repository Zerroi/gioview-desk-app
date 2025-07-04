package main

import (
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/oligo/gioview/menu"
	"github.com/oligo/gioview/misc"
	"github.com/oligo/gioview/navi"
	"github.com/oligo/gioview/theme"
	"github.com/oligo/gioview/view"
	"image/color"
)

// NavSection represents a section in the navigation drawer.
type NavSection interface {
	Title() string
	Layout(gtx C, th *theme.Theme) D
}

// NavDrawer is the main navigation component.
type NavDrawer struct {
	vm           view.ViewManager
	selectedItem *navi.NavTree
	listItems    []NavSection
	listState    *widget.List
	SectionInset layout.Inset
}

// NaviDrawerStyle provides styling for the NavDrawer.
type NaviDrawerStyle struct {
	*NavDrawer
	Inset layout.Inset
	Bg    color.NRGBA
	Width unit.Dp
}

type simpleItemSection struct {
	item *navi.NavTree
}

type simpleNavItem struct {
	icon *widget.Icon
	name string
}

func NewNavDrawer(vm view.ViewManager) *NavDrawer {
	return &NavDrawer{
		vm: vm,
		listState: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
	}
}

func (nv *NavDrawer) AddSection(item NavSection) {
	nv.listItems = append(nv.listItems, item)
}

func (nv *NavDrawer) Layout(gtx C, th *theme.Theme) D {
	if nv.SectionInset == (layout.Inset{}) {
		nv.SectionInset = layout.Inset{
			Bottom: unit.Dp(5),
		}
	}
	list := material.List(th.Theme, nv.listState)
	return list.Layout(gtx, len(nv.listItems), func(gtx C, index int) D {
		item := nv.listItems[index]
		return nv.SectionInset.Layout(gtx, func(gtx C) D {
			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					if item.Title() == "" {
						return layout.Dimensions{}
					}
					return layout.Inset{
						Bottom: unit.Dp(1),
					}.Layout(gtx, func(gtx C) D {
						label := material.Label(th.Theme, th.TextSize*0.8, item.Title())
						label.Color = misc.WithAlpha(th.Fg, 0xb6)
						label.Font.Weight = font.Bold
						return label.Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return item.Layout(gtx, th)
				}),
			)
		})
	})
}

func (nv *NavDrawer) OnItemSelected(item *navi.NavTree) {
	if item != nv.selectedItem {
		if nv.selectedItem != nil {
			nv.selectedItem.Unselect()
		}
		nv.selectedItem = item
	}
	nv.vm.Invalidate()
}

func (ns NaviDrawerStyle) Layout(gtx C, th *theme.Theme) D {
	gtx.Constraints.Max.X = gtx.Dp(ns.Width)
	gtx.Constraints.Min = gtx.Constraints.Max
	rect := clip.Rect{Max: gtx.Constraints.Max}
	paint.FillShape(gtx.Ops, ns.Bg, rect.Op())

	return ns.Inset.Layout(gtx, func(gtx C) D {
		return ns.NavDrawer.Layout(gtx, th)
	})
}

func (item simpleNavItem) Layout(gtx C, th *theme.Theme, textColor color.NRGBA) D {
	return layout.Flex{
		Alignment: layout.Middle,
	}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			if item.icon == nil {
				return D{}
			}
			return layout.Inset{Right: unit.Dp(8)}.Layout(gtx, func(gtx C) D {
				return misc.Icon{Icon: item.icon, Color: textColor, Size: unit.Dp(18)}.Layout(gtx, th)
			})
		}),
		layout.Rigid(func(gtx C) D {
			label := material.Label(th.Theme, th.TextSize, item.name)
			label.Color = textColor
			return label.Layout(gtx)
		}),
	)
}

func (item simpleNavItem) ContextMenuOptions(gtx C) ([][]menu.MenuOption, bool) {
	return nil, false
}

func (item simpleNavItem) Children() ([]navi.NavItem, bool) {
	return nil, false
}

func (ss simpleItemSection) Title() string {
	return ""
}

func (ss simpleItemSection) Layout(gtx C, th *theme.Theme) D {
	return ss.item.Layout(gtx, th)
}

func SimpleItemSection(icon *widget.Icon, name string, onSelect func(item *navi.NavTree)) NavSection {
	item := navi.NewNavItem(simpleNavItem{icon: icon, name: name}, onSelect)
	item.VerticalPadding = unit.Dp(8)
	item.Indention = unit.Dp(16)
	return simpleItemSection{item: item}
}
