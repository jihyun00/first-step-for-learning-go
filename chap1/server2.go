package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock() // 여러 개의 클라이언트가 동시 호출 시, count값이 들쭉날쭉 될 수 있으니깐 mutex로 한 번에 하나만 호출하도록
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
