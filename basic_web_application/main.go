package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Start...")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		n, err := fmt.Fprintf(writer, "Hello world!")
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("Bites number:", n)
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
