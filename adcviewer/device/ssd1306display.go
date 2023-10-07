package device

import (
	"machine"

	"tinygo.org/x/drivers/ssd1306"
)

func InitializeSsd1306Display() *ssd1306.Device {
	machine.I2C1.Configure(machine.I2CConfig{Frequency: 400 * machine.KHz})
	display := ssd1306.NewI2C(machine.I2C1)
	display.Configure(ssd1306.Config{Width: 128, Height: 64, Address: ssd1306.Address_128_32, VccState: ssd1306.SWITCHCAPVCC})

	display.ClearBuffer()
	display.ClearDisplay()

	return &display
}
