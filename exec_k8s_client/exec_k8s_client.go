// $ go get k8s.io/client-go/...
// $ go run exec_k8s_client.go -k '/home/USERNAME/.kube/confg' -n 'kube-system'

package main

import (
	"flag"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig string
	namespace  string
)

func init() {
	flag.StringVar(&kubeconfig, "k", "", "Absolute path to the kubeconfig file, e.g. '/home/USERNAME/.kube/config'.")
	flag.StringVar(&namespace, "n", "default", "Kubernetes namespace in which the client should operate.")
}

func main() {
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("There are %d pods in the cluster:\n", len(pods.Items))
	for idx, pod := range(pods.Items) {
		fmt.Printf("%d: %s\t%s\t%s\t%s\n", idx, pod.Name, pod.Status.HostIP, pod.Status.Phase, pod.CreationTimestamp)
	}
}
