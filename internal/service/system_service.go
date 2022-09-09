package service

import (
	"raspberrypi_colling/internal/reposytory"
)

type SystemService struct {
	reposytory reposytory.ISystemReposytory
}

func NewSystemService(reposytory reposytory.ISystemReposytory) *SystemService {
	service := &SystemService{
		reposytory: reposytory,
	}
	return service
}

func (systemService *SystemService) GetCurrentTemperature() (float32, error) {
	return systemService.reposytory.GetCurrentTemperature()
}
