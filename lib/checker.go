package lib

type Checker interface {
	GetHealth() HealthStatus
}
