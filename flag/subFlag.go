package main

import (
	"flag"
	"log"
)

var name string

func main() {
	flag.Parse()//将命令行解析为定义的标志
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)//返回带有指定名称和错误处理属性的空命令集，创建新的命令集支持子命令
	goCmd.StringVar(&name, "name", "Go语言-Flag", "学习")
	javaCmd := flag.NewFlagSet("java", flag.ExitOnError)
	javaCmd.StringVar(&name, "j", "java语言", "java-cli-help")

	args := flag.Args()
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "java":
		_ = javaCmd.Parse(args[1:])
	}
	log.Printf("name: %s", name)
}
