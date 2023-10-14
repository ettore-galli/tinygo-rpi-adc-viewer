package adcw

import (
	"image/color"
	"sync"
	"time"
)

type AdcViewerSettings struct {
	SamplingDelayMicros float64
}

type signalTrace struct {
	curX   int
	values [128]uint16
}

type ADCSensor interface {
	Get() uint16
}

type SSD1306Display interface {
	SetPixel(x int16, y int16, c color.RGBA)
	Display() error
}

type RunEnvironment struct {
	Settings AdcViewerSettings
	Sensor   ADCSensor
	Display  SSD1306Display
}

func calculateResidualSleepTimeMicroseconds(samplingDelayMicros float64, knownSamplingTimeMicros float64) time.Duration {
	return time.Duration(time.Duration(samplingDelayMicros - knownSamplingTimeMicros).Microseconds())
}

func AdcLoop(sensor ADCSensor, samplingDelayMicros float64, displayValueCallback func(uint16)) {

	var sensorValue uint16

	for {
		sensorValue = sensor.Get()
		displayValueCallback(sensorValue)
		time.Sleep(time.Microsecond * calculateResidualSleepTimeMicroseconds(samplingDelayMicros, 0.2))
	}
}

func IAmAliveLoop() {
	for {
		println("I am alive...")
		time.Sleep(time.Millisecond * 1000)
	}
}

func ScaleSensorValueToTraceDisplayRange(value uint16) byte {
	return byte(value >> 10) // 0-65535 --> 0-64
}

func writeTraceOnDisplay(display SSD1306Display, trace *signalTrace, value uint16) {
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	yValue := ScaleSensorValueToTraceDisplayRange(value)

	display.SetPixel(int16(trace.curX), int16(trace.values[trace.curX]), black)
	display.SetPixel(int16(trace.curX), int16(yValue), white)

	trace.values[trace.curX] = uint16(yValue)

	trace.curX = (trace.curX + 1) % len(trace.values)

	if trace.curX == 0 {
		display.Display()
	}

}

func RunSignalTracer(runEnvironment RunEnvironment) {
	var mainWg sync.WaitGroup

	mainWg.Add(1)

	trace := signalTrace{curX: 0}

	displayValueCallback := func(value uint16) {
		writeTraceOnDisplay(runEnvironment.Display, &trace, value)
	}

	go AdcLoop(runEnvironment.Sensor, runEnvironment.Settings.SamplingDelayMicros, displayValueCallback)

	go IAmAliveLoop()

	mainWg.Wait()
}
