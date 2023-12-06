package kubefunc

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Nodes(clientset *kubernetes.Clientset) {
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error al listar nodos:", err.Error())
		return
	}

	fmt.Println("Nodos en el clúster:")
	for _, node := range nodes.Items {
		fmt.Printf("- %s\n", node.Name)
	}
}
