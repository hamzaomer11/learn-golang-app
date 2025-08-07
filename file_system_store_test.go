package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()
		expected := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, expected)

		got = store.GetLeague()
		assertLeague(t, got, expected)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Chris")
		expected := 33

		if got != expected {
			t.Errorf("got %d, expected %d", got, expected)
		}
	})
}
