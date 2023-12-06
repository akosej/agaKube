package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "./config", "ruta al kubeconfig")
	listPods := flag.Bool("list-pods", false, "lista los pods del clúster")
	listServices := flag.Bool("list-services", false, "lista los servicios del clúster")
	listNodes := flag.Bool("list-nodes", false, "lista los nodos del clúster")
	watchPods := flag.Bool("watch-pods", false, "vigila los pods del clúster")
	// nameSp := flag.Bool("name-space", false, "Crear nameSpace")
	// deploy := flag.Bool("deploy", false, "Crear deploy")
	// newNamespace := flag.String("namespace", "new-namespace", "nombre del nuevo Namespace")
	// deploymentFile := flag.String("file", "deployment.yaml", "archivo YAML con la definición del despliegue")

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Error al crear la configuración:", err.Error())
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Error al crear el cliente:", err.Error())
		os.Exit(1)
	}

	if *listPods {
		ListPods(clientset)
	}

	if *listServices {
		ListServices(clientset)
	}

	if *listNodes {
		ListNodes(clientset)
	}

	if *watchPods {
		WatchPods(clientset)
	}
	// Llama a la función DeployFromFile para crear un Deployment desde un archivo YAML
	// if *deploy {
	// 	err = DeployFromFile(clientset, *deploymentFile)
	// 	if err != nil {
	// 		fmt.Println("Error al crear el Deployment:", err.Error())
	// 	}
	// }
	// if *nameSp {
	// 	err = CreateNamespace(clientset, *newNamespace)
	// 	if err != nil {
	// 		fmt.Println("Error al crear el Namespace:", err.Error())
	// 	}
	// }
}

func ListPods(clientset *kubernetes.Clientset) {
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

func ListServices(clientset *kubernetes.Clientset) {
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

func ListNodes(clientset *kubernetes.Clientset) {
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

// func DeployFromFile(clientset *kubernetes.Clientset, filename string) error {
// 	yamlFile, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return fmt.Errorf("error al leer el archivo YAML: %v", err)
// 	}

// 	// Serializa el YAML en un objeto Unstructured
// 	decode := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
// 	obj, _, err := decode.Decode(yamlFile, nil, nil)
// 	if err != nil {
// 		return fmt.Errorf("error al decodificar el YAML: %v", err)
// 	}

// 	// Convierte el objeto Unstructured a un objeto Deployment
// 	deployment, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&appsv1.Deployment{})
// 	if err != nil {
// 		return fmt.Errorf("error al convertir a objeto Deployment: %v", err)
// 	}

// 	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(*unstructured.Unstructured).Object, deployment)
// 	if err != nil {
// 		return fmt.Errorf("error al convertir a objeto Deployment: %v", err)
// 	}

// 	// Crea el despliegue
// 	_, err = clientset.AppsV1().Deployments("default").Create(context.Background(), &appsv1.Deployment{
// 		ObjectMeta: deployment["metadata"].(appsv1.ObjectMeta),
// 		Spec:       deployment["spec"].(appsv1.DeploymentSpec),
// 	}, metav1.CreateOptions{})
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Deployment creado exitosamente.")
// 	return nil
// }

// func CreateNamespace(clientset *kubernetes.Clientset, namespaceName string) error {
// 	newNamespace := &metav1.Namespace{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: namespaceName,
// 		},
// 	}

// 	_, err := clientset.CoreV1().Namespaces().Create(context.Background(), newNamespace, metav1.CreateOptions{})
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Namespace creado exitosamente.")
// 	return nil
// }
