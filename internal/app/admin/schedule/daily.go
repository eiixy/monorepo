package schedule

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/eiixy/monorepo/internal/app/admin/data"
	"github.com/eiixy/monorepo/internal/data/example/ent/account"
)

type Daily struct {
	data *data.Data
}

func NewDaily(data *data.Data) *Daily {
	return &Daily{data: data}
}

// Run mock daily send stat data to accounts
func (r Daily) Run(ctx context.Context) error {
	startID := 0
	for {
		accounts, err := r.data.EntClient.Account.Query().Where(account.IDGT(startID)).Limit(200).All(ctx)
		if err != nil {
			return err
		}
		if len(accounts) == 0 {
			break
		}
		var messages []*sarama.ProducerMessage
		for i := range accounts {
			messages = append(messages, &sarama.ProducerMessage{
				Topic: "send_daily_stat",
				Key:   sarama.StringEncoder(fmt.Sprintf("account:%d", accounts[i].ID)),
				Value: sarama.StringEncoder("stat data..."),
			})
		}
		err = r.data.Producer.SendMessages(messages)
		if err != nil {
			return err
		}
		startID = accounts[len(accounts)-1].ID
	}
	return nil
}
