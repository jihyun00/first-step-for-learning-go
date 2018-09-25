package main

import (
	"fmt"
	"log"
	"net/http"
)

// 8000번 포트로 서버 띄우고, path 호출해주는 예제
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
