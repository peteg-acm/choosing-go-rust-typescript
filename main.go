package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)
var curCount int
var otherCount int
var statsCount int

func handleRootRequest(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path == "/goodbye" {
		os.Exit(0)
	}

	log.Printf("%s, %s", r.Method, r.URL.Path)


	if r.URL.Path == "/stats" {
		statsCount++
		fmt.Fprintf(w, "<p>Stats: <br> calls to root %d <br> calls to stats %d <br> other calls %d <br> total calls %d",
	              curCount, statsCount, otherCount, curCount+statsCount+otherCount)
		return
	}
	if r.URL.Path != "/" {
		otherCount++
		fmt.Fprintf(w, "I am sorry but I can't find %s", r.URL.Path)
		http.NotFound(w,r)
		return
	}

	curCount++
	if curCount % 2 == 0 {
     fmt.Fprintf(w, "Hello World! you seem to be at " + r.RemoteAddr)
	} else {
    fmt.Fprintf(w, "Hello again for the %d time!", curCount)
	}
}

func main() {
	http.HandleFunc("/", handleRootRequest)
	http.ListenAndServe(":8080", nil)
}
