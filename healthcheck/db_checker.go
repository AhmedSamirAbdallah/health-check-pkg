package healthcheck

import (
	"github.com/AhmedSamirAbdallah/health-check-pkg/pkg/db"
	"time"
)

type DatabaseChecker struct {
	DBURI          string
	DatabaseName   string
	CollectionName string
}

func (d *DatabaseChecker) Name() string {
	return "database"
}

func (d *DatabaseChecker) Check() map[string]interface{} {
	startTime := time.Now()

	db.InitDB(d.DBURI)

	connection := db.CheckDatabase()
	readOK := db.CheckReadOnDB(d.DatabaseName, d.CollectionName)
	writeOK := db.CheckWriteOnDB(d.DatabaseName, d.CollectionName)

	return map[string]interface{}{
		"latency":    time.Since(startTime).String(),
		"connection": connection,
		"read":       readOK,
		"write":      writeOK,
	}

}
