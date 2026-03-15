package scenario

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aaronlee232/redis-gui-tester/internal/models"
	ui "github.com/aaronlee232/redis-gui-tester/internal/ui/components"
)

func GetScenario(w http.ResponseWriter, r *http.Request) {

}

func GetAllScenarios(w http.ResponseWriter, r *http.Request) {
	scenario1 := models.NewScenario("Some Title", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", []string{"command 1", "command 2"}, []string{"response 1", "response 2"}, models.StatusUntested)
	scenario2 := models.NewScenario("Another Title", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", []string{"command 1", "command 2"}, []string{"response 1", "response 2"}, models.StatusPassed)
	scenario3 := models.NewScenario("Another Title", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", []string{"command 1", "command 2"}, []string{"response 1", "response 2"}, models.StatusFailed)
	scenarios := []*models.Scenario{scenario1, scenario2, scenario3}

	// 1. Set a custom header
	w.Header().Set("Content-Type", "text/plain")

	// 2. Set the HTTP status code (must be called before w.Write)
	w.WriteHeader(http.StatusOK)

	// 3. Write the response body to include html scenario list
	ui.ScenarioList(scenarios).Render(r.Context(), w)
}

func CreateScenario(w http.ResponseWriter, r *http.Request) {
	var payload models.Scenario
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Inserts scenario object into sql
	// return error if failed

	w.Header().Set("HX-Trigger", "refreshScenarioList")
	w.WriteHeader(http.StatusOK)
}

func UpdateScenario(w http.ResponseWriter, r *http.Request) {

}

func DeleteScenario(w http.ResponseWriter, r *http.Request) {

}
