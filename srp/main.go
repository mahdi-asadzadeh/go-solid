package main

import (
	"io/ioutil"
	"strings"
)

type NoteBook struct {
	Notes []string
}

func AddANote(note *NoteBook, noteStr string) string {
	note.Notes = append(note.Notes, noteStr)
	return noteStr
}

func SaveFile(note *NoteBook, filename string) bool {
	err := ioutil.WriteFile(filename, []byte(strings.Join(note.Notes, "\n")), 0644)
	if err != nil {
		return false
	}
	return true
}

func main() {
	note := NoteBook{}

	AddANote(&note, "ABC")
	AddANote(&note, "DEF")
	AddANote(&note, "JKL")
	SaveFile(&note, "a.txt")
}
