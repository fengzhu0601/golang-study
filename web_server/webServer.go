package main

import (
	"net/http"
	"log"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "hello golang!!!")
}

func main(){
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":9091", nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}
