package usecase

import (
	"github.com/koskuma/go-worker-workshop/pkg/repository"
)

type Hashtag struct {
	keywordRepository    repository.IKeyword
	hashtagJobRepository repository.IHashtagJob
}

func NewHashtagUsecase(keywordRepository repository.IKeyword, hashtagJobRepository repository.IHashtagJob) *Hashtag {
	return &Hashtag{
		keywordRepository:    keywordRepository,
		hashtagJobRepository: hashtagJobRepository,
	}
}

func (h Hashtag) PublishAllHashtags() {
	hashtags, err := h.keywordRepository.GetAll()
	if err != nil {
		panic(err)
	}
	h.hashtagJobRepository.PublishAllHashtags(hashtags)
}
