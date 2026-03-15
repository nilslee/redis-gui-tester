package models

type ScenarioStatus string

const (
	StatusUntested ScenarioStatus = "untested"
	StatusPassed   ScenarioStatus = "passed"
	StatusFailed   ScenarioStatus = "failed"
)

type Scenario struct {
	ID                int            `json:"ID"`
	Title             string         `json:"title"`
	Description       string         `json:"description"`
	Commands          []string       `json:"commands"`
	ExpectedResponses []string       `json:"expected_responses"`
	Status            ScenarioStatus `json:"status"`
}

func NewScenario(title string, description string, commands []string, expectedResponses []string, status ScenarioStatus) *Scenario {
	return &Scenario{
		Title:             title,
		Description:       description,
		Commands:          commands,
		ExpectedResponses: expectedResponses,
		Status:            status,
	}
}
