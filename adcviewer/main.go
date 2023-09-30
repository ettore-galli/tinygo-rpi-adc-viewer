package main

import (
	"image/color"
	"machine"
	"sync"
	"time"

	"tinygo.org/x/drivers/ssd1306"
)

func LedLoop(led machine.Pin, delayMs float64) {
	for {
		led.Low()
		time.Sleep(time.Millisecond * time.Duration(delayMs))
		led.High()
		time.Sleep(time.Millisecond * time.Duration(delayMs))
	}
}

func AdcLoop(sensor machine.ADC, samplingDelayMs float64, valueCallback func(uint16)) {
	var val uint16
	for {
		val = sensor.Get()
		valueCallback(val)
		time.Sleep(time.Millisecond * time.Duration(samplingDelayMs))
	}
}

func IAmAliveLoop() {
	for {
		println("I am alive...")
		time.Sleep(time.Millisecond * 1000)
	}
}

func createImageBufferFromValue(value uint16) []byte {
	const BufferLength int = 1024
	const Pages int = 8
	const PageLengthBytes int = 128

	buffer := make([]byte, BufferLength)

	barValue := byte(value >> 9) // 0-65535 --> 0-127
	for page := 0; page < Pages; page++ {
		for i := 0; i < PageLengthBytes; i++ {
			bufferPosition := page*PageLengthBytes + i
			if i < int(barValue) {
				buffer[bufferPosition] = 0xff
			} else {
				buffer[bufferPosition] = 0x00
			}
		}
	}

	return buffer
}

func writeBufferOnDisplay(display ssd1306.Device, imgBuffer []byte) {

	err := display.SetBuffer(imgBuffer)
	if err != nil {
		println(err)
	}

	display.Display()
}

func writeValueOnDisplay(display ssd1306.Device, value uint16) {
	imgBuffer := createImageBufferFromValue(value)
	writeBufferOnDisplay(display, imgBuffer)
}

func initializeADCSensor() machine.ADC {
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

type signalTrace struct {
	curX   int
	values [128]uint16
}

func writeTraceOnDisplay(display ssd1306.Device, trace *signalTrace, value uint16) {
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	yValue := byte(value >> 10) // 0-65535 --> 0-64

	display.SetPixel(int16(trace.curX), int16(trace.values[trace.curX]), black)
	display.SetPixel(int16(trace.curX), int16(yValue), white)

	trace.values[trace.curX] = uint16(yValue)

	trace.curX = (trace.curX + 1) % len(trace.values)

	display.Display()
}

func main() {

	var mainWg sync.WaitGroup
	mainWg.Add(1)

	var samplingDelayMs float64 = 1

	led := machine.Pin(0)
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	var sensor machine.ADC = initializeADCSensor()

	var display ssd1306.Device = initializeSsd1306Display()

	trace := signalTrace{curX: 0}

	valueCallback := func(value uint16) {
		// writeValueOnDisplay(display, value)
		writeTraceOnDisplay(display, &trace, value)
	}

	go AdcLoop(sensor, samplingDelayMs, valueCallback)

	// var delayMs float64 = 300
	// go LedLoop(led, delayMs)

	go IAmAliveLoop()

	mainWg.Wait()

}
