package main

import (
	"fmt"
	"k8s.io/client-go/tools/cache"
	"reflect"
)

func main() {
	informerMap := make(map[reflect.Type]cache.SharedIndexInformer)
	fmt.Println(len(informerMap))
	for range informerMap {
		fmt.Println("sss")
	}
}

type A interface {
	sing()
	say()
}

type a struct {
	A
}

func test(d A) {

}
