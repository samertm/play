package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format("2006-01-02"))
	log.Printf("%#v", t)
}
