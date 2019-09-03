package lib

type HealthState string

const (
	UP   HealthState = "up"
	DOWN HealthState = "down"
)

type HealthReasonSeverity string

const (
	HIGH   HealthReasonSeverity = "high"
	MEDIUM HealthReasonSeverity = "medium"
	LOW    HealthReasonSeverity = "low"
	INFO   HealthReasonSeverity = "info"
)

type HealthReason struct {
	Code        string               `json:"code"`
	Severity    HealthReasonSeverity `json:"severity"`
	Description string               `json:"desc"`
	MetaData    map[string]string    `json:"meta"`
}

type HealthStatus struct {
	State  HealthState  `json:"state"`
	Reason HealthReason `json:"reason"`
}
