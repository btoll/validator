package validators

import (
	"context"
	"log"
	"os"
	"path/filepath"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
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

func GetDeploymentClient(kind string) (*v1.Deployment, error) {
	clientset, err := GetClient()
	if err != nil {
		return nil, err
	}
	deploymentsClient := clientset.AppsV1().Deployments("default")
	return deploymentsClient.Get(context.TODO(), kind, metav1.GetOptions{})
}

func GetIngressClient(kind string) (*networkingv1.Ingress, error) {
	clientset, err := GetClient()
	if err != nil {
		return nil, err
	}
	ingressClient := clientset.NetworkingV1().Ingresses("default")
	return ingressClient.Get(context.TODO(), kind, metav1.GetOptions{})
}

func GetServiceClient(kind string) (*corev1.Service, error) {
	clientset, err := GetClient()
	if err != nil {
		return nil, err
	}
	servicesClient := clientset.CoreV1().Services("default")
	return servicesClient.Get(context.TODO(), kind, metav1.GetOptions{})
}
