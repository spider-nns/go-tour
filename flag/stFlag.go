package main

import (
	"flag"
	"log"
)

func main() {
	//go run stFlag.go -name=x -n=xx
	//out name:xx
	var name string
	flag.StringVar(&name, "name", "Go编程之旅", "帮助信息")
	flag.StringVar(&name, "n", "Go编程之旅", "帮助信息")
	flag.Parse()
	log.Printf("name: %s", name)
}

//标准库 flag
//StringVar 方法实现对命令行参数解析和绑定
//参数  命令行标示位，默认值，帮助信息
//-flag 	仅支持布尔类型
//-flag x   仅支持非布尔类型
//-flag=x   都支持
