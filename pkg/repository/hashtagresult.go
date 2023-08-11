package repository

import (
	"encoding/json"

	"github.com/koskuma/go-worker-workshop/config"
	"github.com/koskuma/go-worker-workshop/pkg/adapter"
	"github.com/koskuma/go-worker-workshop/pkg/entity"
)

type IHashtagResult interface {
	Publish(hashtags entity.Result)
}

type hashtagResult struct {
	rabbitMQAdapter adapter.IRabbitMQAdapter
	queueConfig     *config.QueueConfig
}

func NewHashtagResult(rabbitMQAdapter adapter.IRabbitMQAdapter, queueConfig *config.QueueConfig) IHashtagResult {
	return hashtagResult{
		rabbitMQAdapter: rabbitMQAdapter,
		queueConfig:     queueConfig,
	}
}

func (h hashtagResult) Publish(hashtag entity.Result) {
	ht, err := json.Marshal(hashtag)
	if err != nil {
		panic(err)
	}
	h.rabbitMQAdapter.Publish(ht, h.queueConfig.ExchangeName, h.queueConfig.QueueName)
}
