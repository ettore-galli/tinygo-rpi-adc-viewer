package clock

import "time"

type RealClock struct {
}

func GetMicrosecondsDuration(microseconds int64) time.Duration {
	return time.Microsecond * time.Duration(microseconds)
}

func GetMillisecondsDuration(milliseconds int64) time.Duration {
	return time.Millisecond * time.Duration(milliseconds)
}

func (clk RealClock) SleepMicroseconds(microseconds int64) {
	time.Sleep(GetMicrosecondsDuration(microseconds))
}

func (clk RealClock) SleepMilliseconds(milliseconds int64) {
	time.Sleep(GetMillisecondsDuration(milliseconds))
}
