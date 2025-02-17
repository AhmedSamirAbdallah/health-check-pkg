package main

import (
	"healthcheck/healthcheck"
	"healthcheck/internal/config"
)

func main() {
	cfg, _ := config.LoadConfig()
	h := healthcheck.NewHealthCheckManager()

	h.Register(&healthcheck.DatabaseChecker{DBURI: cfg.DBURI, DatabaseName: cfg.DatabaseName, CollectionName: cfg.CollectionName})
	h.Register(&healthcheck.RedisChecker{Addr: cfg.RedisURL, Password: cfg.RedisPassword, DB: cfg.RedisDB, Key: cfg.RedisKey, Val: cfg.RedisVal})
	h.Register(&healthcheck.KafkaChecker{Brokers: cfg.KafkaBroker, Topic: cfg.KafkaTopic})
	h.Register(&healthcheck.TemporalCheker{TemporalUrl: cfg.TemporalUrl, WithTls: cfg.WithTLS})
	h.Serve(":8020")
}
