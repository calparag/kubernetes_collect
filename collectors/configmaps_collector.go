package collectors

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ConfigMapInfo represents the structure of ConfigMap data to be collected
type ConfigMapInfo struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Data      map[string]string `json:"data"`
}

// CollectData collects ConfigMap information and saves it to a JSON file
func CollectConfigmaps(clientset *kubernetes.Clientset, resourceType string) {
	configmaps, err := clientset.CoreV1().ConfigMaps("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing ConfigMaps: %v", err)
		return
	}

	var configMapInfos []ConfigMapInfo
	for _, configmap := range configmaps.Items {
		configMapInfo := ConfigMapInfo{
			Name:      configmap.Name,
			Namespace: configmap.Namespace,
			Data:      configmap.Data,
		}
		configMapInfos = append(configMapInfos, configMapInfo)
	}

	// Save the ConfigMap information to a JSON file
	SaveData(resourceType, configMapInfos)
}
