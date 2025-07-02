package main

import (
	"fmt"
	"os"
)

func echo2() {
	var s, sep string // 修正后的变量声明
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
