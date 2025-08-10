package poker_test

import (
	"io"
	"testing"

	poker "github.com/hamzaomer11/learn-go-app"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &poker.Tape{File: file}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	expected := "abc"

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}
