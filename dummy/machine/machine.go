package machine

// GENERAL

const Device = "dummy"

const (
	KHz = 1000
	MHz = 1000_000
	GHz = 1000_000_000
)

const (
	PinAnalog        PinMode = 1
	PinSERCOM        PinMode = 2
	PinSERCOMAlt     PinMode = 3
	PinTimer         PinMode = 4
	PinTimerAlt      PinMode = 5
	PinCom           PinMode = 6
	PinDigital       PinMode = 8
	PinInput         PinMode = 9
	PinInputPullup   PinMode = 10
	PinOutput        PinMode = 11
	PinTCC           PinMode = PinTimer
	PinTCCAlt        PinMode = PinTimerAlt
	PinInputPulldown PinMode = 12
)

// PIN

type PinMode uint8
type PinConfig struct {
	Mode PinMode
}

type Pin uint8

func (p Pin) Set(status bool)            {}
func (p Pin) High()                      {}
func (p Pin) Low()                       {}
func (p Pin) Configure(config PinConfig) {}

// ADC

type ADC struct {
	Pin Pin
}

func (a ADC) Configure(cfg ADCConfig) {}
func (a ADC) Get() uint16             { return 0 }

func InitADC() {}

// I2C

type I2CConfig struct {
	Frequency uint32
}

type I2C struct {
	Bus uint8
}

func (i2c I2C) Tx(addr uint16, w, r []byte) error {
	return nil
}

func (ic I2C) Configure(cfg I2CConfig) error {
	return nil
}

var I2C1 = I2C{Bus: 0}

// ADC

const (
	ADC0 Pin = 0
	ADC1 Pin = 0
	ADC2 Pin = 0
	ADC3 Pin = 0
	ADC4 Pin = 0
	ADC5 Pin = 0
)

type ADCConfig struct {
	Reference  uint32 // analog reference voltage (AREF) in millivolts
	Resolution uint32 // number of bits for a single conversion (e.g., 8, 10, 12)
	Samples    uint32 // number of samples for a single conversion (e.g., 4, 8, 16, 32)
	SampleTime uint32 // sample time, in microseconds (µs)
}
