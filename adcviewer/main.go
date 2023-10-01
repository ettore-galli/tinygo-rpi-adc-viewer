package main

import "machine"

func main() {
	settings := AdcViewerSettings{
		samplingDelayMicros: 1000,
		pin:                 machine.ADC0,
	}
	RunSignalTracer(settings)

}
