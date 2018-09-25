package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ { // 0부터 실행하면 실행파일까지 같이 호출, go는 전치 연산자 허용하지 않음, while 문 없음.
		s+= sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
