package draws

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"os"
)

func draw(w *app.Window) error {

	// ops are the operations from the UI
	var ops op.Ops

	// startButton is a clickable widget
	var startButton widget.Clickable

	// th defines the material design style
	th := material.NewTheme()

	for {
		// first grab the event
		evt := w.NextEvent()

		// then detect the type
		switch typ := evt.(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, typ)
			layout.Inset{
				Top:    0,
				Bottom: 100,
				Left:   0,
				Right:  100,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, &startButton, "Start")
				return btn.Layout(gtx)
			},
			)
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceEnd,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &startButton, "Start")
						return btn.Layout(gtx)
					},
				),
				layout.Rigid(
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)

			typ.Frame(gtx.Ops)

			// and this is sent when the application should exits
		case app.DestroyEvent:
			os.Exit(0)

		}
		
	}
	return nil
}
