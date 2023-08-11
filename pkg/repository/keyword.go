package repository

import (
	"context"
	"time"

	"github.com/koskuma/go-worker-workshop/pkg/adapter"
	"github.com/koskuma/go-worker-workshop/pkg/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type IKeyword interface {
	GetAll() ([]entity.Hashtag, error)
}

type keyword struct {
	databaseAdapter   adapter.IMongoDBAdapter
	keywordCollection adapter.IMongoCollection
	timeout           time.Duration
}

func NewKeyword(databaseAdapter adapter.IMongoDBAdapter, keywordCollection adapter.IMongoCollection, timeout time.Duration) IKeyword {
	return keyword{
		databaseAdapter:   databaseAdapter,
		keywordCollection: keywordCollection,
		timeout:           timeout,
	}
}

func (k keyword) GetAll() ([]entity.Hashtag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), k.timeout)
	defer cancel()

	var hashtags []entity.Hashtag

	err := k.databaseAdapter.Find(ctx, k.keywordCollection, &hashtags, bson.D{})

	if err != nil {
		return nil, err
	}

	return hashtags, nil
}
