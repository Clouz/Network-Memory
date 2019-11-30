// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net"
	"net/http"
	"text/template"
)

type Page struct {
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	HomePageVars := Page{
		Title: "Titolo 1",
		Body:  "Corpo 1",
	}

	_, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
	if err != nil {                                // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
