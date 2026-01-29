package wordlist

import (
	"bufio"
	"os"
	"strings"
)

// WordList contains a list of words and the supported list operations.
type WordList interface {
	// AddWord persists a new word to the existing list.
	AddWord(word string) (err error)

	// GetWords returns all of the words in the existing list.
	GetWords() (words []string, err error)
}

type wordListImpl struct {
	filename string
}

// New instantiates a new WordList.
func New(filename string) WordList {
	return &wordListImpl{filename: filename}
}

// AddWord persists a new word to the existing list.
func (w *wordListImpl) AddWord(word string) (err error) {
	var f *os.File
	// add permission to r/w
	if f, err = os.OpenFile(w.filename, os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0644); err != nil {
		return
	}
	defer f.Close()

	if _, err = f.Write([]byte(word + "\n")); err != nil {
		return
	}

	return
}

// GetWords returns all of the words in the existing list.
func (w *wordListImpl) GetWords() (words []string, err error) {
	var f *os.File
	if f, err = os.Open(w.filename); err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for i := 0; i < 100000; i++ {
		var s string
		if s, err = r.ReadString('\n'); err != nil {
			return
		}

		words = append(words, strings.TrimSpace(s))
	}

	return
}
