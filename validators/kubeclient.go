package validators

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var reServiceName = regexp.MustCompile(`^([a-z\-]*)-[a-z]*$`)

var clientset *kubernetes.Clientset
var done bool

// Let's only get the client once.
func GetClient() (*kubernetes.Clientset, error) {
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
	substring := reServiceName.FindStringSubmatch(deploymentName)
	deployment, err := deploymentsClient.Get(context.TODO(), substring[1], metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}

	//	apiVersion := fmt.Sprintf("      APIVersion: %s\n", m.APIVersion) // empty value in struct
	//	kind := fmt.Sprintf("           Kind: %s\n", m.Kind)              // empty value in struct
	apiVersion := fmt.Sprintf("      APIVersion: %s\n", "apps/v1") // empty value in struct
	kind := fmt.Sprintf("           Kind: %s\n", "Deployment")
	//	metadata := fmt.Sprintf("       Metadata: %+v\n", m.Metadata)
	name := fmt.Sprintf("           Name: %+v\n", deployment.Name)
	namespace := fmt.Sprintf("      Namespace: %+v\n", deployment.Namespace)
	labels := fmt.Sprintf("         Labels:")
	fmt.Println(apiVersion, kind, name, namespace, labels)
	for k, v := range deployment.Labels {
		fmt.Printf("                  %s: %s\n", k, v)
	}

	spec := deployment.Spec
	replicas := fmt.Sprintf("        Replicas: %d\n", int(*spec.Replicas))
	selector := fmt.Sprintf("       Selector: %+v\n", spec.Selector.MatchLabels)
	fmt.Println(replicas, selector)

	container := spec.Template.Spec.Containers[0]
	name = fmt.Sprintf("            Name: %s\n", container.Name)
	image := fmt.Sprintf("          Image: %s\n", container.Image)
	imagePullPolicy := fmt.Sprintf("ImagePullPolicy: %s\n", container.ImagePullPolicy)
	fmt.Println(fmt.Sprintf("             Env:"))
	for _, s := range container.Env {
		fmt.Printf("                  %s: %s\n", s.Name, s.Value)
	}
	//	envFrom := fmt.Sprintf("        EnvFrom: %+v\n", container.EnvFrom[0].ConfigMapRef.LocalObjectReference)
	envFrom := fmt.Sprintf("        EnvFrom: %+v\n", container.EnvFrom)
	ports := fmt.Sprintf("          Ports: %+v\n", container.Ports)
	resources := fmt.Sprintf("      Resources: %+v\n", container.Resources)
	//	fmt.Println(name, image, imagePullPolicy, envVars, envFrom, ports, resources)
	fmt.Println(name, image, imagePullPolicy, envFrom, ports, resources)
}

func PrintIngress(ingressName string) {
	clientset, err := GetClient()
	ingressClient := clientset.NetworkingV1().Ingresses("default")
	substring := reServiceName.FindStringSubmatch(ingressName)
	ingress, err := ingressClient.Get(context.TODO(), substring[1], metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}

	//	apiVersion := fmt.Sprintf("      APIVersion: %s\n", ingress.APIVersion) // empty value in struct
	//	kind := fmt.Sprintf("           Kind: %s\n", ingress.Kind) // empty value in struct
	apiVersion := fmt.Sprintf("      APIVersion: %s\n", "networking.k8s.io/v1")
	kind := fmt.Sprintf("           Kind: %s\n", "Ingress")
	name := fmt.Sprintf("           Name: %+v\n", ingress.Name)
	namespace := fmt.Sprintf("      Namespace: %+v\n", ingress.Namespace)
	annotations := fmt.Sprintf("    Annotations:")
	fmt.Println(apiVersion, kind, name, namespace, annotations)
	for k, v := range ingress.Annotations {
		fmt.Printf("                  %s: %s\n", k, v)
	}

	for _, rule := range ingress.Spec.Rules {
		for _, path := range rule.HTTP.Paths {
			fmt.Println("\n            Rule:")
			fmt.Printf("                     Path: %s\n", path.Path)
			fmt.Printf("                 PathType: %s\n", *path.PathType)
			fmt.Printf("                  Backend: %+v\n\n", path.Backend)
		}
	}
}

func PrintService(serviceName string) {
	clientset, err := GetClient()
	servicesClient := clientset.CoreV1().Services("default")
	substring := reServiceName.FindStringSubmatch(serviceName)
	service, err := servicesClient.Get(context.TODO(), substring[1], metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}

	//	apiVersion := fmt.Sprintf("      APIVersion: %s\n", service.APIVersion) // empty value in struct
	//	kind := fmt.Sprintf("           Kind: %s\n", service.Kind)              // empty value in struct
	apiVersion := fmt.Sprintf("      APIVersion: %s\n", "v1") // empty value in struct
	kind := fmt.Sprintf("           Kind: %s\n", "Service")
	name := fmt.Sprintf("           Name: %+v\n", service.Name)
	namespace := fmt.Sprintf("      Namespace: %+v\n", service.Namespace)
	labels := fmt.Sprintf("         Labels:")
	fmt.Println(apiVersion, kind, name, namespace, labels)
	for k, v := range service.Labels {
		fmt.Printf("                  %s: %s\n", k, v)
	}

	spec := service.Spec
	ports := fmt.Sprintf("           Ports: %+v\n", spec.Ports)
	selector := fmt.Sprintf("       Selector: %+v\n", spec.Selector)
	_type := fmt.Sprintf("           Type: %s\n", spec.Type)
	fmt.Println(ports, selector, _type)
}
