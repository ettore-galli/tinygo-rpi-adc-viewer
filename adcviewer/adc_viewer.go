package main

import (
	"image/color"
	"machine"
	"sync"
	"time"

	"tinygo.org/x/drivers/ssd1306"
)

type AdcViewerSettings struct {
	samplingDelayMicros float64
	pin                 machine.Pin
}

type signalTrace struct {
	curX   int
	values [128]uint16
}

func calculateResidualSleepTimeMicroseconds(samplingDelayMicros float64, knownSamplingTimeMicros float64) time.Duration {
	return time.Duration(time.Duration(samplingDelayMicros - knownSamplingTimeMicros).Microseconds())
}

func AdcLoop(sensor machine.ADC, samplingDelayMicros float64, displayValueCallback func(uint16)) {

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

func initializeADCSensor(pin machine.Pin) machine.ADC {
	var sensor = machine.ADC{
		Pin: machine.ADC0,
	}

	machine.InitADC()
	adcCfg := machine.ADCConfig{}
	sensor.Configure(adcCfg)

	return sensor
}

func initializeSsd1306Display() ssd1306.Device {
	machine.I2C1.Configure(machine.I2CConfig{Frequency: 400 * machine.KHz})
	display := ssd1306.NewI2C(machine.I2C1)
	display.Configure(ssd1306.Config{Width: 128, Height: 64, Address: ssd1306.Address_128_32, VccState: ssd1306.SWITCHCAPVCC})

	display.ClearBuffer()
	display.ClearDisplay()

	return display
}

func scaleSensorValueToTraceDisplayRange(value uint16) byte {
	return byte(value >> 10) // 0-65535 --> 0-64
}

func writeTraceOnDisplay(display ssd1306.Device, trace *signalTrace, value uint16) {
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	yValue := scaleSensorValueToTraceDisplayRange(value)

	display.SetPixel(int16(trace.curX), int16(trace.values[trace.curX]), black)
	display.SetPixel(int16(trace.curX), int16(yValue), white)

	trace.values[trace.curX] = uint16(yValue)

	trace.curX = (trace.curX + 1) % len(trace.values)

	if trace.curX == 0 {
		display.Display()
	}

}

func RunSignalTracer(settings AdcViewerSettings) {
	var mainWg sync.WaitGroup

	mainWg.Add(1)

	var sensor machine.ADC = initializeADCSensor(settings.pin)

	var display ssd1306.Device = initializeSsd1306Display()

	trace := signalTrace{curX: 0}
	displayValueCallback := func(value uint16) {
		writeTraceOnDisplay(display, &trace, value)
	}

	go AdcLoop(sensor, settings.samplingDelayMicros, displayValueCallback)

	go IAmAliveLoop()

	mainWg.Wait()
}
