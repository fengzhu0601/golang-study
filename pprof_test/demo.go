package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"fengzhu/golang-study/pprof_test/data"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("fengzhu"))
		}
	}()

	http.ListenAndServe("0.0.0.0:10000", nil)
}
