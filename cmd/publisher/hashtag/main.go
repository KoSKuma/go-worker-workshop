package main

import (
	"context"
	"fmt"
	"time"

	"github.com/koskuma/go-worker-workshop/config"
	"github.com/koskuma/go-worker-workshop/pkg/adapter"
	"github.com/koskuma/go-worker-workshop/pkg/repository"
	"github.com/koskuma/go-worker-workshop/pkg/usecase"
)

func main() {
	cfg := config.NewHashtagPublisherConfig()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ContextTimeout)*time.Second)
	defer cancel()

	hashtagJobQueueConfig := config.QueueConfig{
		QueueName:    cfg.HashtagJobQueueName,
		ExchangeName: cfg.HashtagJobExchange,
		ExchangeType: cfg.HashtagJobExchangeType,
	}

	rabbitMQAdapter := adapter.NewQueue(cfg.RabbitMQURI)
	rabbitMQAdapter.CreateAndBindQueue(hashtagJobQueueConfig)
	hashtagJobRepo := repository.NewHashtagJob(rabbitMQAdapter, &hashtagJobQueueConfig) // Should config be passed here?

	mongodbClient, err := adapter.NewMongoDBConnection(ctx, cfg.MongoDBURI)
	if err != nil {
		panic(err)
	}
	defer mongodbClient.Disconnect(ctx)
	mongodbCollection := mongodbClient.Database(cfg.KeywordDatabase).Collection(cfg.KeywordCollection)
	mongoAdapter := adapter.NewMongoDBAdapter(mongodbClient)
	keywordRepo := repository.NewKeyword(mongoAdapter, mongodbCollection, time.Duration(cfg.ContextTimeout)*time.Second)

	hashtagUsecase := usecase.NewHashtagUsecase(keywordRepo, hashtagJobRepo)
	hashtagUsecase.PublishAllHashtags()

	fmt.Println("Published all hashtags successfully!")
}
