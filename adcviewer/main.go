package main

import (
	"adcviewer/adcw"
	"adcviewer/device"
	"time"
)

type TimeClock struct {
}

func (clk TimeClock) SleepMicroseconds(microseconds float64) {
	time.Sleep(time.Microsecond * time.Duration(microseconds))
}
func (clk TimeClock) SleepMilliseconds(milliseconds float64) {
	time.Sleep(time.Millisecond * time.Duration(milliseconds))
}

func main() {

	settings := adcw.AdcViewerSettings{
		SamplingDelayMicros: 1000,
	}

	var sensor adcw.ADCSensor
	sensor = device.InitializeADCSensor()

	var display adcw.SSD1306Display
	display = device.InitializeSsd1306Display()

	adcw.RunSignalTracer(adcw.RunEnvironment{Settings: settings, Sensor: sensor, Display: display, Clock: TimeClock{}})

}
