package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	str := r.URL.Path[1:]
	validate(&str)
	haha := type_1{str}
	fmt.Fprintf(w, "Here I am, stuck with "+haha.str)
}

func validate(str *string) {
	if len(*str) == 0 {
		*str = "nothing"
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	response := ""
	response += header.Get("secret") + "\n"

	b := make([]byte, r.ContentLength)

	_, err := r.Body.Read(b)
	fmt.Println(string(b))
	var t test_struct
	if err != nil {
		err := json.Unmarshal(b, &t)
		if err != nil {
			panic(err)
		}
	}

	response += "haha\n" + t.One + "\n" + t.Two + "\n" + t.Three

	fmt.Println("You got SERVED!")

	fmt.Fprintf(w, response)
}
