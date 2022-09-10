package reposytory

type ISystemReposytory interface {
	GetCurrentTemperature() (float32, error)
}
