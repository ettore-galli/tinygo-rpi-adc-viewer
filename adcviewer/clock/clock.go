package clock

import "time"

type RealClock struct {
}

func (clk RealClock) SleepMicroseconds(microseconds float64) {
	time.Sleep(time.Microsecond * time.Duration(microseconds))
}
func (clk RealClock) SleepMilliseconds(milliseconds float64) {
	time.Sleep(time.Millisecond * time.Duration(milliseconds))
}
