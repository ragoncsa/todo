package http

import (
	"testing"
	"time"
)

func TestTimestampTimeToJSON(t *testing.T) {
	s := "2019-10-12T07:20:50.52Z"
	ttt := &TimestampTime{}
	bin := make([]byte, len(s))
	bin = append(bin, s...)

	err := ttt.UnmarshalJSON(bin)
	if err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}

	want, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if got := ttt; !got.Equal(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestTimestampTimeFromJSON(t *testing.T) {

}

func TestTaskMapping(t *testing.T) {

}
