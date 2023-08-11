package repository

import (
	"encoding/json"

	"github.com/koskuma/go-worker-workshop/config"
	"github.com/koskuma/go-worker-workshop/pkg/adapter"
	"github.com/koskuma/go-worker-workshop/pkg/entity"
)

type IHashtagJob interface {
	PublishAllHashtags(hashtags []entity.Hashtag)
}

type hashtagJob struct {
	rabbitMQAdapter adapter.IRabbitMQAdapter
	queueConfig     *config.QueueConfig
}

func NewHashtagJob(rabbitMQAdapter adapter.IRabbitMQAdapter, queueConfig *config.QueueConfig) IHashtagJob {
	return hashtagJob{
		rabbitMQAdapter: rabbitMQAdapter,
		queueConfig:     queueConfig,
	}
}

func (h hashtagJob) PublishAllHashtags(hashtags []entity.Hashtag) {
	for _, hashtag := range hashtags {
		ht, err := json.Marshal(hashtag)
		if err != nil {
			panic(err)
		}
		h.rabbitMQAdapter.Publish(ht, h.queueConfig.ExchangeName, h.queueConfig.QueueName)
	}
}
