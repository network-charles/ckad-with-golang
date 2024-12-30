package kindcluster

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"sigs.k8s.io/kind/pkg/cluster"
)

// Name of the cluster
var clusterName string = "kind"

// Check if the cluster exists using 'kind get clusters'
func Cluster_exists() bool {
	// Run the command: kind get clusters
	cmd := exec.Command("kind", "get", "clusters")
	output, err := cmd.Output()

	if err != nil {
		log.Fatalf("Error executing kind get clusters: %v", err)
	}

	// Check if the output contains the clusterName
	if string(output) == "" {
		// No clusters found
		return false
	}

	// If the cluster is found in the output
	return true
}

func Delete_cluster() {
	provider := cluster.NewProvider()

	// get kubeconfig
	home, _ := os.UserHomeDir()
	kubeConfigPath := filepath.Join(home, ".kube/config")

	// delete cluster
	err := provider.Delete(clusterName, kubeConfigPath)

	if err != nil {
		log.Fatalf("Failed to delete cluster: %v", err)
	}

	log.Println("Kind cluster deleted successfully!")
}

func Create_cluster() {
	// Initialize a new provider
	// This accesses the NewProvider function in the cluster package
	provider := cluster.NewProvider()

	// Name of the cluster
	clusterName := "kind"

	config_file := filepath.Join("pkg", "kindcluster", "kind.yaml")

	// Create the cluster
	err := provider.Create(clusterName, cluster.CreateWithConfigFile(config_file))

	if err != nil {
		log.Fatalf("Failed to create cluster: %v", err)
	}

	log.Println("Kind cluster created successfully!")
}
