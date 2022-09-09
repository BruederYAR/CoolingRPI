package reposytory

type ISystemReposytory interface {
	GetCurrentTemperature() (float32, error)
}

type IGPIOReposytory interface {
	SetPinHigh() (error)
	SetPinLow() (error)
}