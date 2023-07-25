package main

import (
	"fmt"
	"reflect"
)

type ss struct {
}

func main() {
	obj := ss{}
	t := reflect.TypeOf(obj)
	fmt.Println(t.Name())
}
