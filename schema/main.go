package main

import (
	"log"
	"net/url"

	"github.com/gorilla/schema"
)

type blahQuery struct {
	Blah string `schema:"blahblah"`
}

func main() {
	u, _ := url.Parse("something.com/samer")
	uq, _ := url.Parse(u.String() + "?blahblah=ya")
	var q1 blahQuery
	log.Println(schema.NewDecoder().Decode(&q1, u.Query()))
	log.Println(q1)
	var q2 blahQuery
	log.Println(schema.NewDecoder().Decode(&q2, uq.Query()))
	log.Println(q2)
	uqq, _ := url.Parse(uq.String() + "&lala=yo")
	var q3 blahQuery
	log.Println(schema.NewDecoder().Decode(&q3, uqq.Query()))
	log.Println(q3)
}
