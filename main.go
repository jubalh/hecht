package main

import (
	"io/ioutil"
	"path"

	"github.com/rivo/tview"
)

var booklibrary_path string = "./lib"
var booklist_view *tview.List
var chapterlist_view *tview.List

func updateChapters() {
	selected := booklist_view.GetCurrentItem()
	text, _ := booklist_view.GetItemText(selected)

	chapterlist_view.Clear()

	chapters, err := ioutil.ReadDir(path.Join(booklibrary_path, text))
	if err != nil {
		panic(err)
	}
	for _, chapter := range chapters {
		chapterlist_view.AddItem(chapter.Name(), "", 0, nil)
	}
}

func main() {
	app := tview.NewApplication()

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	booklist_view = tview.NewList()
	chapterlist_view = tview.NewList().ShowSecondaryText(false)

	books, err := ioutil.ReadDir(booklibrary_path)
	for _, book := range books {
		if book.IsDir() {
			booklist_view.AddItem(book.Name(), "", 0, updateChapters)
		}
	}
	if err != nil {
		panic(err)
	}

	booklist_view.AddItem("Quit", "", 'q', func() {
		app.Stop()
	}).ShowSecondaryText(false)
	booklist_view.SetCurrentItem(0)

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0).
		SetBorders(true).
		AddItem(newPrimitive("HECHT"), 0, 0, 1, 2, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 2, 0, 0, false)

	grid.AddItem(booklist_view, 1, 0, 1, 1, 0, 0, true).
		AddItem(chapterlist_view, 1, 1, 1, 1, 0, 0, false)

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}
