package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"strings"
)

func UserIndexFunc(obj interface{}) ([]string, error) {
	pod := obj.(*v1.Pod)
	usersString := pod.Annotations["users"]
	return strings.Split(usersString, ","), nil
}

func main() {
	index := cache.NewIndexer(cache.MetaNamespaceKeyFunc,
		cache.Indexers{"byUser": UserIndexFunc})

	pod1 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "one",
		Annotations: map[string]string{"users": "ernie,bert"}}}
	pod2 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "two",
		Annotations: map[string]string{"users": "bert,oscar"}}}
	pod3 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "tre",
		Annotations: map[string]string{"users": "ernie,elmo"}}}

	index.Add(pod1)
	index.Add(pod2)
	index.Add(pod3)

	erniePods, err := index.ByIndex("byUser", "ernie")
	if err != nil {
		panic(err)
	}

	for _, erniePods := range erniePods {
		fmt.Println(erniePods.(*v1.Pod).Name)
	}

	fmt.Println(index.ListKeys())
	fmt.Println(index.IndexKeys("byUser", "ernie"))
}
