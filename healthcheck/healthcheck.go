package healthcheck

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthCheckManager struct {
	startTime time.Time
	checkers  []Checker
}

// NewHealthCheckManager initializes a new manager
func NewHealthCheckManager() *HealthCheckManager {
	return &HealthCheckManager{
		startTime: time.Now(),
		checkers:  []Checker{},
	}
}

// Register allows adding a new dependency check
func (h *HealthCheckManager) Register(checker Checker) {
	h.checkers = append(h.checkers, checker)
}

// Handler returns an HTTP handler for health checks
func (h *HealthCheckManager) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"status":       "UP",
			"uptime":       time.Since(h.startTime).String(),
			"dependencies": map[string]interface{}{},
		}
		for _, checker := range h.checkers {
			resp["dependencies"].(map[string]interface{})[checker.Name()] = checker.Check()
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}

}

// Serve starts the HTTP server with health check route
func (h *HealthCheckManager) Serve(port string) {
	http.Handle("GET /api/health-check", h.Handler())
	http.ListenAndServe(port, nil)
}
