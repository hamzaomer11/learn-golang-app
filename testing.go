package poker

import (
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
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

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{at, amount})
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
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
