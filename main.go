package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/akosej/agaKube/kubefunc"
	"github.com/akosej/agaKube/pfsenseapi"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "./config", "ruta al kubeconfig")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Falta el comando")
		os.Exit(1)
	}

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

	switch args[0] {
	case "ls":
		if len(args) == 1 {
			fmt.Println("Falta el argumento para listar (nodes, srv o pods)")
			os.Exit(1)
		}
		switch args[1] {
		case "nodes":
			kubefunc.Nodes(clientset)
		case "srv":
			kubefunc.Services(clientset)
		case "pods":
			if len(args) > 2 && args[2] == "-w" {
				kubefunc.WatchPods(clientset)
			} else {
				kubefunc.Pods(clientset)
			}
		default:
			fmt.Println("Comando inválido")
			os.Exit(1)
		}
	case "pf":
		switch args[1] {
		case "dhcp":
			PfSenseDHCPList()
		default:
			fmt.Println("Comando inválido")
			os.Exit(1)
		}
	default:
		fmt.Println("Comando inválido")
		os.Exit(1)
	}
}

func PfSenseDHCPList() {
	ctx := context.Background()
	// client := pfsenseapi.NewClientWithLocalAuth(
	// 	"https://192.168.10.1",
	// 	"admin",
	// 	"adminpassword",
	// )
	client := pfsenseapi.NewClientWithTokenAuth(
		"https://192.168.10.1",
		"id",
		"token",
	)

	leases, err := client.DHCP.ListLeases(ctx)
	if err != nil {
		panic(err)
	}

	for _, lease := range leases {
		fmt.Println(lease.Ip)
	}
}
