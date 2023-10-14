package clock

import (
	"testing"
	"time"
)

func TestGetMicrosecondsDuration(t *testing.T) {
	errorText := "Result expected to be %v, got %v"

	microsecondsDuration := GetMicrosecondsDuration(1)

	expected := time.Nanosecond * 1000 // 1000ns==1µs

	if microsecondsDuration != expected {
		t.Errorf(errorText, expected, microsecondsDuration)
	}

}

func TestGetMillisecondsDuration(t *testing.T) {
	errorText := "Result expected to be %v, got %v"

	millisecondsDuration := GetMillisecondsDuration(1)

	expected := time.Nanosecond * 1000 * 1000 // 1ms

	if millisecondsDuration != expected {
		t.Errorf(errorText, expected, millisecondsDuration)
	}

}
