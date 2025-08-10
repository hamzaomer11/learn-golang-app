package poker_test

import (
	"bytes"
	"strings"
	"testing"

	poker "github.com/hamzaomer11/learn-go-app"
)

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

var dummySpyAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, dummyPlayerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, dummyPlayerStore, "Cleo")
	})

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		expectedPrompt := poker.PlayerPrompt

		if gotPrompt != expectedPrompt {
			t.Errorf("got %q, expected %q", gotPrompt, expectedPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("expected Start called with 7 but got %d", game.StartedWith)
		}
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should not have started")
		}

		gotPrompt := stdout.String()

		expectedPrompt := poker.PlayerPrompt + "you're so silly"

		if gotPrompt != expectedPrompt {
			t.Errorf("got %q, expected %q", gotPrompt, expectedPrompt)
		}
	})
}

func assertScheduledAlert(t *testing.T, got, expected poker.ScheduledAlert) {
	t.Helper()
	if got != expected {
		t.Errorf("got %+v, expected %+v", got, expected)
	}
}
