package entity

type Hashtag struct {
	// MongoID string `example:"1234" json:"_id" bson:"_id"`
	ID      string `example:"1234" json:"id" bson:"id"`
	Keyword string `example:"golang" json:"keyword" bson:"keyword"`
}
