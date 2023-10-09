package device

import (
	"machine"
)

func InitializeADCSensor() *machine.ADC {
	var sensor = machine.ADC{
		Pin: machine.ADC0,
	}

	machine.InitADC()
	adcCfg := machine.ADCConfig{}
	sensor.Configure(adcCfg)

	return &sensor
}
