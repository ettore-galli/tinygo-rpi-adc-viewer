package main

import (
	"adcviewer/adcw"
	"adcviewer/device"
)

func main() {

	settings := adcw.AdcViewerSettings{
		SamplingDelayMicros: 1000,
	}

	var sensor adcw.ADCSensor
	sensor = device.InitializeADCSensor()

	var display adcw.SSD1306Display
	display = device.InitializeSsd1306Display()

	adcw.RunSignalTracer(settings, sensor, display)

}
