package tester

import (
	"database/sql"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) RegisterRoutes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("POST /run-scenario", h.RunAllScenarios)
	r.HandleFunc("POST /run-scenario/{id}", h.RunScenario)
	r.HandleFunc("GET /run-scenario/{id}", h.GetScenarioOutput)

	return r
}
