package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type HashtagPublisherConfig struct {
	MongoDBURI             string `env:"MONGODB_URI"`
	ContextTimeout         int    `env:"TIMEOUT"`
	KeywordDatabase        string `env:"KEYWORD_DATABASE"`
	KeywordCollection      string `env:"KEYWORD_COLLECTION"`
	RabbitMQURI            string `env:"RABBITMQ_URI"`
	HashtagJobQueueName    string `env:"RABBITMQ_HASHTAG_JOB_QUEUE"`
	HashtagJobExchange     string `env:"RABBITMQ_HASHTAG_JOB_EXCHANGE"`
	HashtagJobExchangeType string `env:"RABBITMQ_HASHTAG_JOB_EXCHANGE_TYPE"`
}

type HashtagWorkerConfig struct {
	ContextTimeout      int    `env:"TIMEOUT"`
	RabbitMQURI         string `env:"RABBITMQ_URI"`
	HashtagJobQueueName string `env:"RABBITMQ_HASHTAG_JOB_QUEUE"`
	ThreadAPIURI        string `env:"THREAD_API_URI"`
	ThreadAPIMaxPost    int    `env:"THREAD_API_MAX_POST"`
	ResultQueueName     string `env:"RABBITMQ_RESULT_QUEUE"`
	ResultExchangeName  string `env:"RABBITMQ_RESULT_EXCHANGE"`
	ResultExchangeType  string `env:"RABBITMQ_RESULT_EXCHANGE_TYPE"`
}

type QueueConfig struct {
	QueueName    string
	ExchangeName string
	ExchangeType string
}

func NewHashtagPublisherConfig() HashtagPublisherConfig {
	godotenv.Load()
	config := HashtagPublisherConfig{}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}
	return config
}

func NewHashtagWorkerConfig() HashtagWorkerConfig {
	godotenv.Load()
	config := HashtagWorkerConfig{}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}
	return config
}
