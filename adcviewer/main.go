package main

func main() {
	settings := AdcViewerSettings{
		samplingDelayMicros: 1000,
	}
	RunSignalTracer(settings)

}
