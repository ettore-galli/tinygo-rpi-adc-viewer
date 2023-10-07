package main

import (
	"testing"
)

func TestScaleSensorValueToTraceDisplayRange(t *testing.T) {
	got := ScaleSensorValueToTraceDisplayRange(32760)

	want := byte(64)

	if want != got {
		t.Errorf("Result expected to be %v, got %v", want, got)
	}

}
