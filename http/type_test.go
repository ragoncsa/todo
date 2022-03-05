package http

import (
	"testing"
	"time"
)

func TestTimestampTimeToJSON(t *testing.T) {
	const timeVal = "2019-10-12T07:20:50Z"
	time, err := time.Parse(time.RFC3339, timeVal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ttt := &TimestampTime{
		Time: time,
	}

	bin, err := ttt.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}

	if got, want := string(bin), "\""+timeVal+"\""; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

}

func TestTimestampTimeFromJSON(t *testing.T) {
	const timeVal = "2019-10-12T07:20:50Z"
	const jsonTimeVal = "\"" + timeVal + "\""
	ttt := &TimestampTime{}
	bin := make([]byte, len(jsonTimeVal))
	bin = append(bin, []byte(jsonTimeVal)...)

	err := ttt.UnmarshalJSON(bin)
	if err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}

	want, err := time.Parse(time.RFC3339, timeVal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := ttt; !got.Equal(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
