package main

import (
	"io/ioutil"
	"strconv"

	"github.com/rivo/tview"
)

var book_library string = "./lib"
var book_list *tview.List
var chapter_list *tview.List

func updateChapters() {
	sel := book_list.GetCurrentItem()
	s := strconv.Itoa(sel)
	chapter_list.Clear()
	chapter_list.AddItem(s, "", 0, nil)
}

func main() {
	app := tview.NewApplication()

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	book_list = tview.NewList()
	chapter_list = tview.NewList()

	books, err := ioutil.ReadDir(book_library)
	for _, book := range books {
		book_list.AddItem(book.Name(), "", 0, updateChapters)
	}
	if err != nil {
		panic(err)
	}

	book_list.AddItem("Quit", "", 'q', func() {
		app.Stop()
	}).ShowSecondaryText(false)
	book_list.SetCurrentItem(0)

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0).
		SetBorders(true).
		AddItem(newPrimitive("HECHT"), 0, 0, 1, 2, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 2, 0, 0, false)

	grid.AddItem(book_list, 1, 0, 1, 1, 0, 0, true).
		AddItem(chapter_list, 1, 1, 1, 1, 0, 0, false)

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}
