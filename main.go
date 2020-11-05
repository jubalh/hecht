package main

import (
	"io/ioutil"

	"github.com/rivo/tview"
)

var book_library string = "./lib"

func main() {
	app := tview.NewApplication()

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	menu := newPrimitive("Menu")
	list := tview.NewList()

	books, err := ioutil.ReadDir(book_library)
	for _, book := range books {
		list.AddItem(book.Name(), "", 0, nil)
	}
	if err != nil {
		panic(err)
	}

	list.AddItem("Quit", "", 'q', func() {
		app.Stop()
	}).ShowSecondaryText(false)

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
