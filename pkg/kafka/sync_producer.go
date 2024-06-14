package kafka

import (
	"github.com/IBM/sarama"
)

func NewSyncProducer(addrs []string, opts ...ConfigOption) (sarama.SyncProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V2_1_1_0
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Return.Successes = true
	cfg.Producer.Partitioner = sarama.NewHashPartitioner
	for _, opt := range opts {
		opt(cfg)
	}

	client, err := sarama.NewSyncProducer(addrs, cfg)
	if err != nil {
		return nil, err
	}

	return client, nil
}
