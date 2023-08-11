package usecase

import (
	"github.com/koskuma/go-worker-workshop/pkg/entity"
	"github.com/koskuma/go-worker-workshop/pkg/repository"
)

type ITimeline interface {
	GetTimeline(job entity.Hashtag)
}

type timeline struct {
	threadRepo        repository.IThread
	hashtagResultRepo repository.IHashtagResult
}

func NewTimeline(threadRepo repository.IThread, hashtagResultRepo repository.IHashtagResult) ITimeline {
	return timeline{
		threadRepo:        threadRepo,
		hashtagResultRepo: hashtagResultRepo,
	}
}

func (t timeline) GetTimeline(job entity.Hashtag) {
	posts := t.threadRepo.GetThread(job.Keyword)
	for _, p := range posts {
		go t.GetAccountAndPublish(p)
	}
}

func (t timeline) GetAccountAndPublish(post entity.ThreadPost) {
	account := t.threadRepo.GetAccount(post.UserID)
	result := entity.Result{
		ID:           post.ID,
		Text:         post.Text,
		UserID:       post.UserID,
		Likes:        post.Likes,
		ParentThread: post.ParentThread,
		RepostCount:  post.RepostCount,
		Account:      account,
	}
	t.hashtagResultRepo.Publish(result)
}
