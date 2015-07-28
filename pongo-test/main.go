package main

import (
	"fmt"

	"github.com/flosch/pongo2"
)

var baseTemplate = pongo2.Must(pongo2.FromFile("base.html"))

// Compile the template at application start (for performance reasons)
var indexTemplate = pongo2.Must(pongo2.FromFile("index.html"))

type User struct {
	Login  string
	Email  string
	Groups []Group
}

type Group struct {
	ID int
}

func main() {
	// Execute the template
	renderedTemplate, err := indexTemplate.Execute(pongo2.Context{
		"user": User{Login: "Samer"},
		"sayhi": func(u User) (string, string) {
			return "Hello, " + u.Login, "yo"
		},
	})
	if err != nil {
		// Handle any execution error (e. g. return a HTTP 500)
		panic(err)
	}
	fmt.Println(renderedTemplate)
}
