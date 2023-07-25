package main

import (
	"k8s.io/client-go/kubernetes"
	typecorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/record"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/wuxiansong/.kube/config")
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	broadcster := record.NewBroadcaster()
	// 启动一个watcher，将event写入log
	broadcster.StartStructuredLogging(6)
	// 启动一个watcher，将event写入api-server
	broadcster.StartRecordingToSink(&typecorev1.EventSinkImpl{Interface: clientset.CoreV1().Events("default")})
}
