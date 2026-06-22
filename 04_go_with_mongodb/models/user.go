package models

type User struct {
	Id     bson.ObjectId `bson:"_id",josn:"id"`
	Name   string        `json:"name",bson:"name"`
	Gender string        `json:"gender,omitempty",bson:"gender"`
	Age    int           `json:"age",bson:"age"`
}

func (u *User) GetId() bson.ObjectId {
	return u.Id
}
