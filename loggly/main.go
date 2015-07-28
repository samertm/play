package main

import (
	"log"
	"os"

	"github.com/segmentio/go-loggly"
)

var LoggyToken = "fffffffffffffffffffffffffffffffffffffff"

func main() {
	c := loggly.New(LoggyToken)
	c.Writer = os.Stdout
	log.Println(c.Info("test", loggly.Message{"hi": "there"}))
	log.Println(c.Flush())
}
