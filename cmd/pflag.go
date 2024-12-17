package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

func main() {
	// 定义一些标志（flags）
	name := pflag.String("name", "World", "your name")
	age := pflag.Int("age", 30, "your age")
	verbose := pflag.Bool("verbose", false, "enable verbose output")

	// 解析命令行标志
	pflag.Parse()

	// 使用命令行标志的值
	fmt.Printf("Hello, %s!\n", *name)
	fmt.Printf("Your age is: %d\n", *age)
	if *verbose {
		fmt.Println("Verbose output enabled.")
	}
}
