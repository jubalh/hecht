package main

import (
	"github.com/rivo/tview"
)

var audiobooks []AudioBook
var booklibrary_path string = "./lib"
var app *tview.Application

func main() {
	audiobooks = Scan(booklibrary_path)
	app = buildUI(audiobooks)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
