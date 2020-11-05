package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	menu := newPrimitive("Menu")
	list := tview.NewList().
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("List item 3", "Some explanatory text", 'c', nil).
		AddItem("List item 4", "Some explanatory text", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0).
		SetBorders(true).
		AddItem(newPrimitive("HECHT"), 0, 0, 1, 2, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 2, 0, 0, false)

	grid.AddItem(menu, 1, 0, 1, 1, 0, 0, false).
		AddItem(list, 1, 1, 1, 1, 0, 0, true)

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}
