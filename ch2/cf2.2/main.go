/*
练习 2.2： 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。
*/
package main

import (
	"bufio"
	"fmt"
	"gopl.io/ch2/lengconv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			s := input.Text()
			printCF(s)
		}
	} else {
		for _, arg := range os.Args[1:] {
			printCF(arg)
		}
	}
}

func printCF(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf2.2: %v\n", err)
		os.Exit(1)
	}
	m := lengconv.Metre(t)
	f := lengconv.Feet(t)
	fmt.Printf("%s = %s, %s = %s\n",
		m, lengconv.MToF(m),
		f, lengconv.FToM(f),
	)
}
