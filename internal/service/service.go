package service

type ISystemService interface {
	GetCurrentTemperature() (float32, error)
}