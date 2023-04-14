// 练习 1.4： 修改dup2，出现重复的行时打印文件名称。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "dup.1.4: 必须输入文件")
	} else {
		for _, arg := range files {
			counts := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup1.4: %v\n", err)
				continue
			}
			if hasDup(f, counts) {
				fmt.Println("文件名: " + f.Name() + "有重复行")
			}
			f.Close()
		}
	}
}

func hasDup(f *os.File, counts map[string]int) bool {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("line: %s 重复, 次数: %d, 结束\n", line, n)
			return true
		}
	}
	return false
}
