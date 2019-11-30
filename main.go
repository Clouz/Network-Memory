// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
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
	http.HandleFunc("/Scan", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	i, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
	if err != nil {                                // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, i) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {       // if there is an error
		log.Print("template executing error: ", err) //log it
	}

}

func Scan(w http.ResponseWriter, r *http.Request) {

	log.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("interface:", r.Form["interface"])
	}

	for _, iface := range net.Interfaces() {
		if iface.Name == r.Form["interface"] {
			scan(&iface)
		}
	}
}
