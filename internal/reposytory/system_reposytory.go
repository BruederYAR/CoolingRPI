package reposytory

import (
	"raspberrypi_colling/pkg/input"
	"strconv"
)

type SystemReposytory struct {
	Config *input.Config
}

func NewSystemReposytory(config input.Config) *SystemReposytory {
	systemReposytory := &SystemReposytory{
		Config: &config,
	}
	return systemReposytory
}

func (systemReposytory *SystemReposytory) GetCurrentTemperature() (float32, error) {
	date, err := input.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return 0, err
	}

	result, err := strconv.Atoi(date)
	if err != nil {
		return 0, err
	}

	return float32(result/1000), nil
}
