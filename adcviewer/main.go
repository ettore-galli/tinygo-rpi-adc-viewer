package main

import "adcviewer/adcw"

func main() {
	settings := adcw.AdcViewerSettings{
		SamplingDelayMicros: 1000,
	}
	adcw.RunSignalTracer(settings)

}
