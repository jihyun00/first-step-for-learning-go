package main

import (
	"flag"
	"fmt"
	"strings"
)

// Go에서 flag 만들기가 이렇게 쉽다니...
var n = flag.Bool("n", false, "omit trailing newline") // n이나 sep는 플래그 변수 포인터로 간접적으로만 접근해야 됨
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse() // 플래그 변수들의 기본값을 갱신하고
	fmt.Print(strings.Join(flag.Args(), *sep)) // flag.Args() 요녀석을 통해 문자열 슬라이스 형태로 얻을 수도 있음
	if !*n {
		fmt.Println()
	}
}
