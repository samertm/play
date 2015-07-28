package main

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	ti := time.Now()
	fmt.Println(ti.Format("2006-01-02"))
	log.Printf("%#v", ti)
}
