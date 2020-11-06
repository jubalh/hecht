package main

import (
	"io/ioutil"
	"os/exec"
	"path"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Chapter struct {
	name   string
	length int
}

type AudioBook struct {
	name     string
	length   int
	chapters []Chapter
}

var audiobooks []AudioBook
var booklibrary_path string = "./lib"
var booklist_view *tview.List
var chapterlist_view *tview.List
var app *tview.Application
var isPlaying bool = false
var cmd *exec.Cmd

func scanBook(path string) []Chapter {
	var chapters []Chapter

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		chapter := Chapter{name: file.Name(), length: 1}
		chapters = append(chapters, chapter)
	}

	return chapters
}

func scanLibrary(libpath string) []AudioBook {
	var audiobooks []AudioBook

	folders, err := ioutil.ReadDir(libpath)
	if err != nil {
		panic(err)
	}

	for _, file := range folders {
		if file.IsDir() {
			book := AudioBook{name: file.Name()}
			bookpath := path.Join(libpath, file.Name())
			book.chapters = scanBook(bookpath)

			audiobooks = append(audiobooks, book)
		}
	}

	return audiobooks
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

func playFile() {
	if isPlaying {
		cmd.Process.Kill()
		isPlaying = false
	} else {
		selected := booklist_view.GetCurrentItem()
		bookname, _ := booklist_view.GetItemText(selected)
		selected = chapterlist_view.GetCurrentItem()
		chaptername, _ := chapterlist_view.GetItemText(selected)

		audiopath := path.Join(booklibrary_path, bookname, chaptername)

		cmd = exec.Command("mpv", audiopath)
		cmd.Start()
		isPlaying = true
	}
}

func updateChapters() {
	selected := booklist_view.GetCurrentItem()

	chapterlist_view.Clear()

	for _, chapter := range audiobooks[selected].chapters {
		chapterlist_view.AddItem(chapter.name, "", 0, playFile)
	}

	app.SetFocus(chapterlist_view)
}

func main() {
	app = tview.NewApplication()

	audiobooks = scanLibrary(booklibrary_path)

	booklist_view = tview.NewList()
	chapterlist_view = tview.NewList().ShowSecondaryText(false)

	for _, book := range audiobooks {
		booklist_view.AddItem(book.name, "", 0, updateChapters)
	}

	booklist_view.AddItem("Quit", "", 'q', func() {
		app.Stop()
	}).ShowSecondaryText(false)
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

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}
