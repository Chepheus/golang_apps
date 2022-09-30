package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = "8081"

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello world!")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Bites number:", n)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	fmt.Println("Start...")

	http.HandleFunc("/hello", HelloWorld)

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)

	fmt.Println("Listen port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
