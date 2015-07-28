package main

import (
	"log"
	"net/url"
)

func main() {
	u, _ := url.Parse("http://blahblah.com/whatup?yo=bro")
	log.Println(u.RequestURI(), "\n", u.Path, "\n", u.RawQuery, "\n", u.Fragment)
	log.Println(url.QueryUnescape(url.QueryEscape("samertm.com/something?wkaa=flaka&flay=blah")))
}
