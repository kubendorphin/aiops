package main

import {
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
}

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/keone/.kube/config", "")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("解析kubeconfig文件错误：%s", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	for _, pod := range pod.Items {
		fmt.Println(pod.Name)
	}
	// deployment
	deployments, _ := clientset.AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{})
	for _, deployment := range deployments.Items {
		fmt.Println(deployment.Name)
	}
}