package machine

import "errors"

var (
	ErrTimeoutRNG         = errors.New("machine: RNG Timeout")
	ErrClockRNG           = errors.New("machine: RNG Clock Error")
	ErrSeedRNG            = errors.New("machine: RNG Seed Error")
	ErrInvalidInputPin    = errors.New("machine: invalid input pin")
	ErrInvalidOutputPin   = errors.New("machine: invalid output pin")
	ErrInvalidClockPin    = errors.New("machine: invalid clock pin")
	ErrInvalidDataPin     = errors.New("machine: invalid data pin")
	ErrNoPinChangeChannel = errors.New("machine: no channel available for pin interrupt")
)

// Device is the running program's chip name, such as "ATSAMD51J19A" or
// "nrf52840". It is not the same as the CPU name.
//
// The constant is some hardcoded default value if the program does not target a
// particular chip but instead runs in WebAssembly for example.
const Device = "dummy"

// Generic constants.
const (
	KHz = 1000
	MHz = 1000_000
	GHz = 1000_000_000
)

// PinMode sets the direction and pull mode of the pin. For example, PinOutput
// sets the pin as an output and PinInputPullup sets the pin as an input with a
// pull-up.
type PinMode uint8

type PinConfig struct {
	Mode PinMode
}

const (
	PinAnalog    PinMode = 1
	PinSERCOM    PinMode = 2
	PinSERCOMAlt PinMode = 3
	PinTimer     PinMode = 4
	PinTimerAlt  PinMode = 5
	PinCom       PinMode = 6
	//PinAC_CLK        PinMode = 7
	PinDigital       PinMode = 8
	PinInput         PinMode = 9
	PinInputPullup   PinMode = 10
	PinOutput        PinMode = 11
	PinTCC           PinMode = PinTimer
	PinTCCAlt        PinMode = PinTimerAlt
	PinInputPulldown PinMode = 12
)

// Pin is a single pin on a chip, which may be connected to other hardware
// devices. It can either be used directly as GPIO pin or it can be used in
// other peripherals like ADC, I2C, etc.
type Pin uint8

// NoPin explicitly indicates "not a pin". Use this pin if you want to leave one
// of the pins in a peripheral unconfigured (if supported by the hardware).
const NoPin = Pin(0xff)

func (p Pin) Set(status bool) {}

// High sets this GPIO pin to high, assuming it has been configured as an output
// pin. It is hardware dependent (and often undefined) what happens if you set a
// pin to high that is not configured as an output pin.
func (p Pin) High() {}

// Low sets this GPIO pin to low, assuming it has been configured as an output
// pin. It is hardware dependent (and often undefined) what happens if you set a
// pin to low that is not configured as an output pin.
func (p Pin) Low() {}

func (p Pin) Configure(config PinConfig) {}

type ADC struct {
	Pin Pin
}

func (a ADC) Configure(cfg ADCConfig) {}

func (a ADC) Get() uint16 { return 0 }

type I2C struct {
	Bus uint8
}

var I2C1 = I2C{Bus: 0}

type I2CConfig struct {
	Frequency uint32
}

func (i2c I2C) Tx(addr uint16, w, r []byte) error { return nil }

func (ic I2C) Configure(cfg I2CConfig) error {
	return nil
}

// ADC on the Arduino
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

func InitADC() {}
