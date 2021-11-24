package main

type Mcq struct {
	Id        uint   `json:"id"`
	Statement string `json:"statement"`
	A         string `json:"a"`
	B         string `json:"b"`
	C         string `json:"c"`
	D         string `json:"d"`
	Answer    string `json:"ans"`
}

type contentType struct {
	key   string
	value string
}

func (c *contentType) ContentType() (string, string) {
	return c.key, c.value
}

var ContentType contentType = contentType{key: "content-type", value: "application/json"}
