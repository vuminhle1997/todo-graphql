package models

type Todo struct {
	ID          string `bson:"_id,omitempty"`
	Title       string `bson:"title,omitempty"`
	Description string `bson:"description,omitempty"`
	Done        bool   `bson:"done,omitempty"`
}
