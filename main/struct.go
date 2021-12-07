package main

import "gopkg.in/mgo.v2/bson"

type Mcq struct {
	Id        uint   `json:"id"`
	Statement string `json:"statement"`
	A         string `json:"a"`
	B         string `json:"b"`
	C         string `json:"c"`
	D         string `json:"d"`
	Answer    string `json:"ans"`
}

type McqDB struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Statement string
	A         string
	B         string
	C         string
	D         string
	Answer    string
}

type contentType struct {
	key   string
	value string
}

func (c *contentType) ContentType() (string, string) {
	return c.key, c.value
}

var ContentType contentType = contentType{key: "content-type", value: "application/json"}
