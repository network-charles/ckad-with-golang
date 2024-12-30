package main

import (
	"lab-02/pkg/kindcluster"
	"log"
)

func main() {
	// First, check if the cluster exists
	if kindcluster.Cluster_exists() {
		log.Println("Cluster exists. Deleting the cluster...")
		kindcluster.Delete_cluster()
	} else {
		log.Println("Cluster does not exist. Creating the cluster...")
		kindcluster.Create_cluster()
	}
}
