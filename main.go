package main

import (
	"io/ioutil"
	"os/exec"
	"path"

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

func main() {
	audiobooks = scanLibrary(booklibrary_path)
	app = buildUI(audiobooks)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
