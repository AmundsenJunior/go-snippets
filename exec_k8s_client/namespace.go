// Get pods of namespace on cluster, using k9s config file
// $ go get k8s.io/client-go/...
// $ go run namespace.go -k /home/USERNAME/.kube/confg -n pfs-user [delete|create]

package main

import (
	"flag"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig string
	namespace  string
	command    string
)

func init() {
	flag.StringVar(&kubeconfig, "k", "", "Absolute path to the kubeconfig file, e.g. '/home/USERNAME/.kube/config'.")
	flag.StringVar(&namespace, "n", "default", "Kubernetes namespace to create or delete.")
	// non-flag arg is command [create|delete]
}

func main() {
	flag.Parse()
	args := flag.Args()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	command = args[0]
	switch command {
	case "create":
		_, err = clientset.CoreV1().Namespaces().Create(corev1.Namespace{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Namespace %s has been created.\n", namespace)
	case "delete":
		err = clientset.CoreV1().Namespaces().Delete(namespace, metav1.DeleteOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Namespace %s has been deleted.\n", namespace)
	}
}
