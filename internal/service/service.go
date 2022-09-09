package service

type ISystemService interface {
	GetCurrentTemperature() (float32, error)
}

type IGPIOService interface {
	SetPinHigh() (error)
	SetPinLow() (error)
}