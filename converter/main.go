package main

import (
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type sdf struct {
}

func main() {
	schema := runtime.NewScheme()
	schema.AddKnownTypes(appsv1.SchemeGroupVersion, &appsv1.Deployment{})
	metav1.AddToGroupVersion(schema, appsv1.SchemeGroupVersion)
}
