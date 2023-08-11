package main

import (
	"encoding/json"
	"fmt"

	"github.com/koskuma/go-worker-workshop/config"
	"github.com/koskuma/go-worker-workshop/pkg/adapter"
	"github.com/koskuma/go-worker-workshop/pkg/entity"
	"github.com/koskuma/go-worker-workshop/pkg/repository"
	"github.com/koskuma/go-worker-workshop/pkg/usecase"
)

func main() {
	cfg := config.NewHashtagWorkerConfig()

	resultQueueConfig := config.QueueConfig{
		QueueName:    cfg.ResultQueueName,
		ExchangeName: cfg.ResultExchangeName,
		ExchangeType: cfg.ResultExchangeType,
	}

	rabbitMQAdapter := adapter.NewQueue(cfg.RabbitMQURI)
	rabbitMQAdapter.CreateAndBindQueue(resultQueueConfig)
	hashtagResultRepo := repository.NewHashtagResult(rabbitMQAdapter, &resultQueueConfig)

	threadAdapter := adapter.NewThreadAdapter(cfg.ThreadAPIURI, cfg.ThreadAPIMaxPost)
	threadRepo := repository.NewThreadRepo(threadAdapter)

	threadTimelineUsecase := usecase.NewTimeline(threadRepo, hashtagResultRepo)

	fmt.Println("Subscribing to queue")
	rabbitMQAdapter.Subscribe(cfg.HashtagJobQueueName, func(message []byte) {
		fmt.Println("Received message: ", string(message))
		job := entity.Hashtag{}
		json.Unmarshal(message, &job)
		threadTimelineUsecase.GetTimeline(job)
		fmt.Println("Finished processing message")
	})
}
