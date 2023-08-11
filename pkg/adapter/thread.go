package adapter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/koskuma/go-worker-workshop/pkg/entity"
)

type Thread struct {
	url      string
	pageSize int
}

func NewThreadAdapter(url string, pageSize int) *Thread {
	return &Thread{url: url, pageSize: pageSize}
}

func (t Thread) GetThread(hashtag string, cursor string) ([]entity.ThreadPost, string) {
	endpoint := t.url + "/thread/?hashtag=" + hashtag + "&page_size=" + fmt.Sprint(t.pageSize)
	if cursor != "" {
		endpoint = endpoint + "&cursor=" + cursor
	}
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint, nil)

	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	thread := entity.ThreadResponse{}
	err = json.Unmarshal(body, &thread)
	if err != nil {
		panic(err)
	}

	return thread.Posts, thread.NextPage
}

func (t Thread) GetAccount(accountId string) entity.ThreadAccount {
	endpoint := t.url + "/account/" + accountId
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint, nil)

	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	account := entity.ThreadAccount{}
	err = json.Unmarshal(body, &account)
	if err != nil {
		panic(err)
	}

	return account
}
