package models

type User struct {
	UserID    string `bson:"_id,omitempty" json:"userId,omitempty"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password" bson:"-"`
}
