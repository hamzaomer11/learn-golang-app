package poker

import (
	"fmt"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int, to io.Writer) {
	s.Alerts = append(s.Alerts, ScheduledAlert{at, amount})
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.WinCalls[0], winner)
	}
}

func AssertStatus(t testing.TB, got *httptest.ResponseRecorder, expected int) {
	t.Helper()
	if got.Code != expected {
		t.Errorf("did not get correct status, got %d, expected %d", got.Code, expected)
	}
}

func AssertResponseBody(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("response body is wrong, got %q, expected %q", got, expected)
	}
}

func AssertLeague(t testing.TB, got, expected []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expect %v", got, expected)
	}
}

func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, expected string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != expected {
		t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
	}
}
