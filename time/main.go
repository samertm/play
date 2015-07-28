package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format("2006-01-02"))
}
