package main

import (
	"fmt"
	"os"
	"raspberrypi_colling/internal/reposytory"
	"raspberrypi_colling/internal/service"
	"raspberrypi_colling/pkg/input"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	//create config
	config, err := input.ReadConfigurationFile("appsetting.json")
	if err != nil {
		config = &input.Config{
			Pin:            21,
			Delay:          1,
			MaxTemperature: 55,
		}
	}

	//create layers
	systemReposytory := reposytory.NewSystemReposytory(*config)
	systemService := service.NewSystemService(systemReposytory)

	control(config, systemService)
}

func control(config *input.Config, system *service.SystemService) {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	pin := rpio.Pin(10)
	pin.Output()
	pin.Low()
	pinState := false

	for true {
		temperature, err := system.GetCurrentTemperature()
		if err != nil {
			pin.High()
			time.Sleep(time.Second * time.Duration(config.Delay))
			continue
		}

		if temperature > float32(config.MaxTemperature) && !pinState || temperature < float32(config.MaxTemperature)-5 && pinState {
			pinState = !pinState

			if pinState {
				pin.High()
			} else {
				pin.Low()
			}
		}
	}
}
