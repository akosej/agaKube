package kubefunc

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Pods(clientset *kubernetes.Clientset) {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error al listar pods:", err.Error())
		return
	}

	fmt.Println("Pods en el clúster:")
	for _, pod := range pods.Items {
		fmt.Printf("- %s\n", pod.Name)
	}
}

func WatchPods(clientset *kubernetes.Clientset) {
	watcher, err := clientset.CoreV1().Pods("").Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error al iniciar la vigilancia de los pods:", err.Error())
		return
	}
	fmt.Println("Vigilando cambios en los Pods del clúster...")

	for event := range watcher.ResultChan() {
		pod, ok := event.Object.(*corev1.Pod)
		if !ok {
			fmt.Println("Error al obtener el objeto Pod")
			continue
		}

		fmt.Printf("Evento: %s Pod: %s en estado: %s\n", event.Type, pod.Name, pod.Status.Phase)
	}
}
