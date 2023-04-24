package user

import "gopkg.in/mgo.v2/bson"

// user holds data for single user
type user struct {
	ID   bson.ObjectId `json:"id" storm:"id"`
	Name string        `json:"name"`
	Role string        `json:"role"`
}
