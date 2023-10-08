package adcw

import (
	"testing"
)

func TestScaleSensorValueToTraceDisplayRange(t *testing.T) {
	got := ScaleSensorValueToTraceDisplayRange(32760)

	want := byte(63)

	if want != got {
		t.Errorf("Result expected to be %v, got %v", want, got)
	}

}
