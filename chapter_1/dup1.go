package main

import (
	"bufio"
	"fmt"
	"os"
)

func dup_1() {
	counts := make(map[string]int)

	// 使用Scanner读取输入（无需创建ReadWriter）
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		counts[line]++
	}

	// 处理可能的错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "读取输入错误: %v\n", err)
		return
	}

	// 输出重复行
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
