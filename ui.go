package main

import (
	"log"
	"os/exec"
	"path"
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/jubalh/hecht/library"
	"github.com/rivo/tview"
)

var booklist_view *tview.List
var chapterlist_view *tview.List

var isPlaying bool = false
var cmd *exec.Cmd

func playFile() {
	if isPlaying {
		err := cmd.Process.Kill()
		if err != nil {
			log.Fatal(err)
		}
		isPlaying = false
	} else {
		selected := booklist_view.GetCurrentItem()
		bookname, _ := booklist_view.GetItemText(selected)
		selected = chapterlist_view.GetCurrentItem()
		chaptername, _ := chapterlist_view.GetItemText(selected)

		audiopath := path.Join(booklibrary_path, bookname, chaptername)

		cmd = exec.Command("mpv", audiopath)
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		isPlaying = true
	}
}

func updateChapters() {
	selected := booklist_view.GetCurrentItem()

	chapterlist_view.Clear()

	for _, chapter := range audiobooks[selected].Chapters {
		chapterlist_view.AddItem(chapter.Name, "", 0, playFile)
	}

	app.SetFocus(chapterlist_view)
}

func navigationHandler(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyRight:
		app.SetFocus(chapterlist_view)
		return nil
	case tcell.KeyLeft:
		app.SetFocus(booklist_view)
		return nil
	case tcell.KeyEsc:
		app.Stop()
	}
	return event
}

func buildUI(audiobooks []library.AudioBook) *tview.Application {
	app = tview.NewApplication()

	booklist_view = tview.NewList()
	chapterlist_view = tview.NewList().ShowSecondaryText(false)

	for _, book := range audiobooks {
		booklist_view.AddItem(book.Name, strconv.Itoa(book.Length)+" minutes", 0, updateChapters)
	}

	booklist_view.AddItem("Quit", "", 'q', func() {
		app.Stop()
	}).ShowSecondaryText(true)

	booklist_view.SetCurrentItem(0)

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0).
		SetBorders(true).
		AddItem(newPrimitive("HECHT"), 0, 0, 1, 2, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 2, 0, 0, false)

	grid.AddItem(booklist_view, 1, 0, 1, 1, 0, 0, true).
		AddItem(chapterlist_view, 1, 1, 1, 1, 0, 0, false)

	app.SetInputCapture(navigationHandler)

	app.SetRoot(grid, true).SetFocus(grid)

	return app
}
