package main

import (
	"context"
	"lab-04/pkg/kindcluster"
	"log"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Uses the "k8s.io/api/core/v1", and "k8s.io/apimachinery/pkg/apis/meta/v1" package
func createPod(clientset *kubernetes.Clientset) {
	// Define a Pod
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx-pod",
			Namespace: "default",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx-container",
					Image: "nginx:latest",
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 80,
						},
					},
				},
			},
		},
	}

	// Create the Pod in the "default" namespace
	_, err := clientset.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create Pod: %v", err)
	}

	log.Println("Pod created successfully!")
}

// Function to authenticate and create a Kubernetes client
func authenticateToCluster() (*kubernetes.Clientset, error) {
	// Load kubeconfig to connect to the cluster
	home, _ := os.UserHomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")

	// Use the "k8s.io/client-go/tools/clientcmd" package
	// Provides the "BuildConfigFromFlags" function to load the
	// cluster configuration from the kubeconfig file.
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	// Create a Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func main() {
	// Check if the Kind cluster exists
	if kindcluster.Cluster_exists() {
		log.Println("Cluster exists. Deleting the cluster...")
		kindcluster.Delete_cluster()
	} else {
		log.Println("Cluster does not exist. Creating the cluster...")
		kindcluster.Create_cluster()
	}

	// Authenticate to cluster
	clientset, err := authenticateToCluster()
	if err != nil {
		log.Fatalf("Failed to authenticate to the cluster: %v", err)
	}

	// Create a Pod
	createPod(clientset)
}
