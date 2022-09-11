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
			TemperaturePath: "/usr/local/bin/thermal/temp",
		}
	}

	//create layers
	systemReposytory := reposytory.NewSystemReposytory(*config)
	systemService := service.NewSystemService(systemReposytory)

	control(config, systemService)
}

func control(config *input.Config, system *service.SystemService) {
	if err := rpio.Open(); err != nil {
		fmt.Println("[ERR] " + err.Error())
		os.Exit(1)
	}
	defer rpio.Close()

	pin := rpio.Pin(config.Pin)
	pin.Output()
	pin.Low()
	pinState := false

	for true {
		temperature, err := system.GetCurrentTemperature()
		if err != nil {
			fmt.Println("[ERR] " + err.Error())
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

		time.Sleep(time.Second * time.Duration(config.Delay))
	}
}
