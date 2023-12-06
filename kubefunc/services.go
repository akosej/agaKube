package kubefunc

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Services(clientset *kubernetes.Clientset) {
	services, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error al listar servicios:", err.Error())
		return
	}

	fmt.Println("Servicios en el clúster:")
	for _, service := range services.Items {
		fmt.Printf("- %s\n", service.Name)
	}
}
