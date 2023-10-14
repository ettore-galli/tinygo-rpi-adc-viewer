package main

import (
	"adcviewer/adcw"
	"adcviewer/clock"
	"adcviewer/device"
)

func main() {

	settings := adcw.AdcViewerSettings{
		SamplingDelayMicros: 100,
	}

	var sensor adcw.ADCSensor
	sensor = device.InitializeADCSensor()

	var display adcw.SSD1306Display
	display = device.InitializeSsd1306Display()

	adcw.RunSignalTracer(adcw.RunEnvironment{Settings: settings, Sensor: sensor, Display: display, Clock: clock.RealClock{}})

}
