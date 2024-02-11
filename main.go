package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	html :=
		//STYLE :))
		"<head>" +
			"<style>" +
			"body {height: 100vh; width: 100vw; background-color: pink; display: flex; align-items: center; justify-content: center}" +
			"div {}" +
			"div {color:green; height: auto; width: fit-content; animation: 1s linear infinite rodar; margin: 0 auto; text-align:center} " +
			" @keyframes rodar { 0% { transform: rotate(0deg)}; 100% {transform: rotate(360deg)} } " +
			"</style>" +
			//DIV :)
			"<body><div><h1>ROTAÇÃO</h1></div></body>" +
			"</head>"

	http.HandleFunc("/ayusdguh", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, html)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
