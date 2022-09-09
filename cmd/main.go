package main

import (
	"fmt"
	"raspberrypi_colling/internal/reposytory"
	"raspberrypi_colling/internal/service"
	"raspberrypi_colling/pkg/input"
	"time"
)

func main() {
	//create config
	config, err := input.ReadConfigurationFile("appsetting.json")
	if err != nil {
		config = &input.Config{
			Pin:            40,
			Delay:          1,
			MaxTemperature: 55,
		}
	}

	//create layers
	systemReposytory := reposytory.NewSystemReposytory(*config)
	systemService := service.NewSystemService(systemReposytory)

	gPIOReposytory := reposytory.NewGPIOReposytory(*config)
	gPIOService := service.NewGPIOService(gPIOReposytory)

	fmt.Println(systemService.GetCurrentTemperature())

	for i := 0; i < 10; i++ {
		gPIOService.SetPinHigh()
		time.Sleep(time.Second / 5)
		gPIOService.SetPinHigh()
	}
}
