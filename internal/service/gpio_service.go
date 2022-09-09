package service

import "raspberrypi_colling/internal/reposytory"

type GPIOService struct {
	Reposytory reposytory.IGPIOReposytory
}

func NewGPIOService(reposytory reposytory.IGPIOReposytory) *GPIOService {
	gPIOService := &GPIOService{
		Reposytory: reposytory,
	}
	return gPIOService
}


func (gPIOService *GPIOService)	SetPinHigh() (error) {
	return gPIOService.Reposytory.SetPinHigh()
}
func (gPIOService *GPIOService) SetPinLow() (error) {
	return gPIOService.SetPinLow()
}
