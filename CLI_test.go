package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Chris wins\n")
	playerStore := &StubPlayerStore{}

	cli := &CLI{playerStore, in}
	cli.PlayerPoker()

	if len(playerStore.winCalls) != 1 {
		t.Fatal("expected a win but didn't record any")
	}

	got := playerStore.winCalls[0]
	expected := "Chris"

	if got != expected {
		t.Errorf("didn't record correct winner, got %q, expected %q", got, expected)
	}
}
