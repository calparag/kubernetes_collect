package main

import (
	"encoding/json"
	"kube_collect/collectors"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type ResourceConfig struct {
	Type    string `json:"type"`
	Collect bool   `json:"collect"`
}

func main() {
	// Load the configuration file
	configPath := filepath.Join("config", "data_collect.json")
	configData, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var resourceConfigs []ResourceConfig
	err = json.Unmarshal(configData, &resourceConfigs)
	if err != nil {
		log.Fatalf("Error unmarshaling config JSON: %v", err)
	}

	// Initialize Kubernetes client
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	// Iterate through the configuration and call collectors
	for _, resourceConfig := range resourceConfigs {
		if resourceConfig.Collect {
			switch resourceConfig.Type {
			case "pod":
				collectors.CollectPod(clientset, "pods")
			case "deployment":
				collectors.CollectDeployment(clientset, "deployments")
			case "service":
				collectors.CollectService(clientset, "services")
			case "persistent_volumes":
				collectors.CollectPV(clientset, "persistent_volumes")
			case "persistent_volume_claims":
				collectors.CollectPvc(clientset, "persistent_volume_claims")
			case "secrets":
				collectors.CollectSecret(clientset, "secrets")
			case "configmaps":
				collectors.CollectConfigmaps(clientset, "configmaps")
			case "ingresses":
				collectors.CollectIngress(clientset, "ingresses")
			case "network_policies":
				collectors.CollectNetworkPolicy(clientset, "network_policies")
			case "roles":
				collectors.CollectRoles(clientset, "roles")
			case "role_bindings":
				collectors.CollectRoleBindings(clientset, "role_bindings")
			case "cluster_roles":
				collectors.CollectClusterRoles(clientset, "cluster_roles")
			case "cluster_role_bindings":
				collectors.CollectClusterRoleBindings(clientset, "cluster_role_bindings")
			case "crds":
				collectors.CollectCrds(config, "crds")
			case "daemonsets":
				collectors.CollectDaemonSets(clientset, "daemonsets")
			case "statefulsets":
				collectors.CollectStatefulSet(clientset, "statefulsets")
			case "jobs":
				collectors.CollectJobs(clientset, "jobs")
			case "cronjobs":
				collectors.CollectCronJobs(clientset, "cronjobs")
			case "hpas":
				collectors.CollectHpa(clientset, "hpas")
			case "events":
				collectors.CollectEvents(clientset, "events")
			// Add more cases for other resources
			default:
				log.Printf("Unknown resource type: %s", resourceConfig.Type)
			}

		}
	}

	log.Println("Data collection completed.")
}
