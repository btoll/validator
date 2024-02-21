package validators

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var clientset *kubernetes.Clientset
var done bool

func GetClient() (*kubernetes.Clientset, error) {
	// Let's only get the client once.
	if !done {
		done = true
		kubeconfig := filepath.Join(
			os.Getenv("HOME"), ".kube", "config",
		)
		kubeconfig = filepath.Clean(kubeconfig)
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatal(err)
		}
		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatal(err)
		}
	}
	return clientset, nil
}

func PrintDeployment(deploymentName string) {
	clientset, err := GetClient()
	deploymentsClient := clientset.AppsV1().Deployments("default")
	deployment, err := deploymentsClient.Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}

	// This values are empty in the returned struct.
	deployment.APIVersion = "apps/v1"
	deployment.Kind = "Deployment"

	f, err := os.Create(fmt.Sprintf("build/%s/deployment/remote", deploymentName))
	if err != nil {
		fmt.Println("err", err)
	}
	err = tpl.ExecuteTemplate(f, "deployment.tpl", deployment)
	if err != nil {
		fmt.Println("err", err)
	}
}

func PrintIngress(ingressName string) {
	clientset, err := GetClient()
	ingressClient := clientset.NetworkingV1().Ingresses("default")
	ingress, err := ingressClient.Get(context.TODO(), ingressName, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}

	// This values are empty in the returned struct.
	ingress.APIVersion = "networking.k8s.io/v1"
	ingress.Kind = "Ingress"

	f, err := os.Create(fmt.Sprintf("build/%s/ingress/remote", ingressName))
	if err != nil {
		fmt.Println("err", err)
	}
	err = tpl.ExecuteTemplate(f, "ingress.tpl", ingress)
	if err != nil {
		fmt.Println("err", err)
	}
}

func PrintService(serviceName string) {
	clientset, err := GetClient()
	servicesClient := clientset.CoreV1().Services("default")
	service, err := servicesClient.Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}

	// This values are empty in the returned struct.
	service.APIVersion = "v1"
	service.Kind = "Service"

	f, err := os.Create(fmt.Sprintf("build/%s/service/remote", serviceName))
	if err != nil {
		fmt.Println("err", err)
	}
	err = tpl.ExecuteTemplate(f, "service.tpl", service)
	if err != nil {
		fmt.Println("err", err)
	}
}
