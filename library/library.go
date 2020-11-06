package library

import (
	"io/ioutil"
	"path"
)

type Chapter struct {
	Name   string
	Length int
}

type AudioBook struct {
	Name     string
	Length   int
	Chapters []Chapter
}

/* scanBook scans the (audio) files of a folder (book) and returns the book as a slice of Chapter */
func scanBook(path string) []Chapter {
	var chapters []Chapter

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		chapter := Chapter{Name: file.Name(), Length: 1}
		chapters = append(chapters, chapter)
	}

	return chapters
}

/* scan scans all the books into the library and returns it as a slice of AudioBook */
func Scan(libpath string) []AudioBook {
	var audiobooks []AudioBook

	folders, err := ioutil.ReadDir(libpath)
	if err != nil {
		panic(err)
	}

	for _, file := range folders {
		if file.IsDir() {
			book := AudioBook{Name: file.Name()}
			bookpath := path.Join(libpath, file.Name())
			book.Chapters = scanBook(bookpath)

			audiobooks = append(audiobooks, book)
		}
	}

	return audiobooks
}
