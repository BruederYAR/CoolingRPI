package reposytory

import (
	"raspberrypi_colling/pkg/input"

	"github.com/stianeikeland/go-rpio/v4"
)

type GPIOReposytory struct {
	Config *input.Config
}

func NewGPIOReposytory(config input.Config) *GPIOReposytory {
	gPIOReposytory := &GPIOReposytory{
		Config: &config,
	}
	return gPIOReposytory
}

func (gPIOReposytory *GPIOReposytory) SetPinHigh() error {
	pin := rpio.Pin(gPIOReposytory.Config.Pin)

	if err := rpio.Open(); err != nil {
		return err
	}

	defer rpio.Close()

	pin.Output()
	pin.High()

	rpio.Close()
	return nil
}

func (gPIOReposytory *GPIOReposytory) SetPinLow() error {
	pin := rpio.Pin(gPIOReposytory.Config.Pin)

	if err := rpio.Open(); err != nil {
		return err
	}

	defer rpio.Close()

	pin.Output()
	pin.Low()

	rpio.Close()
	return nil
}
