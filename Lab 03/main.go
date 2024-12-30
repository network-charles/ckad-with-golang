package main

import (
	"lab-03/pkg/kindcluster"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Authenticate to cluster
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
	// Uses the "lab-03/pkg/kindcluster" and "log" package
	if kindcluster.Cluster_exists() {
		log.Println("Cluster exists. Deleting the cluster...")
		kindcluster.Delete_cluster()
	} else {
		log.Println("Cluster does not exist. Creating the cluster...")
		kindcluster.Create_cluster()
	}

	// Authenticate to cluster
	authenticateToCluster()
}
