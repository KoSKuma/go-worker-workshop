package repository

import (
	"github.com/koskuma/go-worker-workshop/pkg/adapter"
	"github.com/koskuma/go-worker-workshop/pkg/entity"
)

type IThread interface {
	GetThread(hashtag string) []entity.ThreadPost
	GetAccount(userID string) entity.ThreadAccount
}

type thread struct {
	adapter *adapter.Thread
}

func NewThreadRepo(threadAdapter *adapter.Thread) *thread {
	return &thread{adapter: threadAdapter}
}

func (t thread) GetThread(hashtag string) []entity.ThreadPost {
	var posts []entity.ThreadPost
	res, cursor := t.adapter.GetThread(hashtag, "")
	for cursor != "" {
		posts, cursor = t.adapter.GetThread(hashtag, "")
		res = append(res, posts...)
	}
	return res
}

func (t thread) GetAccount(userID string) entity.ThreadAccount {
	return t.adapter.GetAccount(userID)
}
