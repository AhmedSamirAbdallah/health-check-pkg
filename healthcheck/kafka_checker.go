package healthcheck

import (
	"github.com/AhmedSamirAbdallah/health-check-pkg/pkg/kafka"
	"time"
)

type KafkaChecker struct {
	Brokers string //localhost:1222,localhost:1223
	Topic   string
}

func (k *KafkaChecker) Name() string {
	return "kafka"
}

func (k *KafkaChecker) Check() map[string]interface{} {
	startTime := time.Now()

	kafka.InitKafka(k.Brokers)

	consumeOK := kafka.CheckConsumer(k.Topic)
	produceOK := kafka.CheckProduce(k.Topic)

	return map[string]interface{}{
		"latency": time.Since(startTime).String(),
		"produce": produceOK,
		"consume": consumeOK,
	}
}
