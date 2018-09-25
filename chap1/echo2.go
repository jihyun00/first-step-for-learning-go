package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// range는 반복할 때 인덱스와 인덱스에 있는 원소의 값을 생성함
	for _, arg := range os.Args[1:] { // _ 요런 빈 식별자를 사용할 수도 있음. 변수명이 필요한데 로직에서 불필요할 때 쓰는 용도임
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
