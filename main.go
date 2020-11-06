package main

import (
	"github.com/rivo/tview"

	"github.com/jubalh/hecht/library"
)

var audiobooks []library.AudioBook
var booklibrary_path string = "./lib"
var app *tview.Application

func main() {
	audiobooks = library.Scan(booklibrary_path)
	app = buildUI(audiobooks)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
