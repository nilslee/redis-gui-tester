package models

type ScenarioStatus string

const (
	StatusUntested ScenarioStatus = "untested"
	StatusPassed   ScenarioStatus = "passed"
	StatusFailed   ScenarioStatus = "failed"
)

type Scenario struct {
	Title            string
	Description      string
	Commands         []string
	ExpectedResponse []string
	Status           ScenarioStatus
}

func NewScenario(title string, description string, commands []string, expectedResponse []string, status ScenarioStatus) *Scenario {
	return &Scenario{
		Title:            title,
		Description:      description,
		Commands:         commands,
		ExpectedResponse: expectedResponse,
		Status:           status,
	}
}
