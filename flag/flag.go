package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

func main() {
	var name Name
	flag.Var(&name, "name", "go run subFlag.go -name=yourCode")
	flag.Parse()
	log.Printf("记录你的每一个nb时刻-> %s", name)
}

type Value interface {
	String() string
	Set(string) error
}
type Name string

func (i *Name) String() string {
	return fmt.Sprint(*i)
}

func (i *Name) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("name flag already set")
	}
	*i = Name("Get command args is: " + value)
	return nil
}
