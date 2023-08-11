package entity

type ThreadResponse struct {
	Posts    []ThreadPost `json:"data"`
	NextPage string       `json:"next_page"`
}

type ThreadPost struct {
	ID           string `json:"_id"`
	Text         string `json:"text"`
	UserID       string `json:"user_id"`
	Likes        int    `json:"likes"`
	ParentThread any    `json:"parent_thread"`
	RepostCount  int    `json:"repost_count"`
}

type ThreadAccount struct {
	ID              string `json:"_id"`
	DisplayName     string `json:"display_name"`
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
	Description     string `json:"description"`
	Follower        int    `json:"follower"`
	Following       int    `json:"following"`
}

type Result struct {
	ID           string        `json:"_id"`
	Text         string        `json:"text"`
	UserID       string        `json:"user_id"`
	Likes        int           `json:"likes"`
	ParentThread any           `json:"parent_thread"`
	RepostCount  int           `json:"repost_count"`
	Account      ThreadAccount `json:"account"`
}
