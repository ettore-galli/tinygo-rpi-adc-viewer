package adcw

import (
	"testing"
)

func TestScaleSensorValueToTraceDisplayRange(t *testing.T) {
	got := ScaleSensorValueToTraceDisplayRange(32768)

	want := byte(32)

	if want != got {
		t.Errorf("Result expected to be %v, got %v", want, got)
	}

}
