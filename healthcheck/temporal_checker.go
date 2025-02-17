package healthcheck

import (
	"healthcheck/internal/temporal"
	"time"
)

type TemporalCheker struct {
	TemporalUrl string
	WithTls     bool
}

func (t *TemporalCheker) Name() string {
	return "temporal"
}

func (t *TemporalCheker) Check() map[string]interface{} {
	startTime := time.Now()
	connection := temporal.CheckTemporalConnection(t.TemporalUrl, t.WithTls)

	return map[string]interface{}{
		"connection": connection,
		"latency":    time.Since(startTime).String(),
	}
}
