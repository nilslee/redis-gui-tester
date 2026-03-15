package scenario

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aaronlee232/redis-gui-tester/internal/models"
	ui "github.com/aaronlee232/redis-gui-tester/internal/ui/components"
	"github.com/aaronlee232/redis-gui-tester/internal/utils"
)

func (h *Handler) GetScenario(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetAllScenarios(w http.ResponseWriter, r *http.Request) {
	// Get scenario with commands and expected responses
	rows, err := h.db.Query(`
	SELECT 
	s.scenario_id, s.title, s.description, s.status,
	(SELECT json_group_array(command_text) 
	FROM (SELECT command_text FROM commands 
	WHERE scenario_id = s.scenario_id ORDER BY step_order)) as commands,
	(SELECT json_group_array(response_text) 
	FROM (SELECT response_text FROM expected_responses 
	WHERE scenario_id = s.scenario_id ORDER BY step_order)) as responses
	FROM scenarios s
	ORDER BY s.created_at DESC;
	`)
	if err != nil {
		log.Println("Failed to get scenarios:", err)
		http.Error(w, "Failed to get scenarios", http.StatusInternalServerError)
	}
	defer rows.Close()

	scenarios := []*models.Scenario{}
	for rows.Next() {
		var s models.Scenario
		var cmdsJSON, respsJSON string

		// Scan column values into variables
		err := rows.Scan(&s.ID, &s.Title, &s.Description, &s.Status, &cmdsJSON, &respsJSON)
		if err != nil {
			log.Println("Scan error:", err)
			continue
		}

		// Unmarshal the JSON strings into your Go slices
		json.Unmarshal([]byte(cmdsJSON), &s.Commands)
		json.Unmarshal([]byte(respsJSON), &s.ExpectedResponses)

		scenarios = append(scenarios, &s)
	}

	// Check for errors that happened during iteration
	if err = rows.Err(); err != nil {
		http.Error(w, "Error during iteration", http.StatusInternalServerError)
		return
	}

	// Return rendered HTML
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	ui.ScenarioList(scenarios).Render(r.Context(), w)
}

func (h *Handler) CreateScenario(w http.ResponseWriter, r *http.Request) {
	var payload models.Scenario
	if err := utils.DecodeRequestJSON(w, r, &payload); err != nil {
		return
	}

	// 1. Use a Transaction (atomic operations)
	tx, err := h.db.BeginTx(r.Context(), nil)
	if err != nil {
		http.Error(w, "Failed to initiate transaction", http.StatusInternalServerError)
		return
	}
	// Safely roll back if something goes wrong
	defer tx.Rollback()

	// Insert scenario
	res, err := tx.ExecContext(r.Context(),
		`INSERT INTO scenarios (title, description, status) VALUES (?, ?, ?)`,
		payload.Title, payload.Description, payload.Status)
	if err != nil {
		log.Println("scenario insert err:", err)
		http.Error(w, "Failed to create scenario", http.StatusInternalServerError)
		return
	}

	// Record Scenario ID
	scenarioID, err := res.LastInsertId()
	if err != nil {
		log.Println("id fetch err:", err)
		http.Error(w, "Failed to fetch ID", http.StatusInternalServerError)
		return
	}

	// Insert commands
	cmdStmt, err := tx.Prepare(`INSERT INTO commands (scenario_id, step_order, command_text) VALUES (?, ?, ?)`)
	if err != nil {
		log.Println("Failed to prepare insert statement for commands:", err)
		http.Error(w, "Failed to insert commands", http.StatusInternalServerError)
		return
	}
	defer cmdStmt.Close()

	for stepOrder, commandText := range payload.Commands {
		if _, err := cmdStmt.Exec(scenarioID, stepOrder, commandText); err != nil {
			log.Println(err)
			http.Error(w, "Failed to create scenario", http.StatusInternalServerError)
			return
		}
	}

	// Insert expected responses
	rspStmt, err := tx.Prepare(`INSERT INTO expected_responses (scenario_id, step_order, response_text) VALUES (?, ?, ?)`)
	if err != nil {
		log.Println("Failed to prepare insert statement for commands:", err)
		http.Error(w, "Failed to insert commands", http.StatusInternalServerError)
		return
	}
	defer rspStmt.Close()

	for stepOrder, responseText := range payload.ExpectedResponses {
		if _, err := rspStmt.Exec(scenarioID, stepOrder, responseText); err != nil {
			log.Println("Failed to insert expected responses:", err)
			http.Error(w, "Failed to insert expected responses", http.StatusInternalServerError)
			return
		}
	}

	// Commit whole batch
	if err := tx.Commit(); err != nil {
		log.Println("Failed to commit transaction:", err)
		http.Error(w, "Transaction failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "refreshScenarioList")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateScenario(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteScenario(w http.ResponseWriter, r *http.Request) {

}
