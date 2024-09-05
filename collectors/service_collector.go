package collectors

import (
	"context"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ServiceInfo represents the structure of service data to be collected
type ServiceInfo struct {
	Name        string               `json:"name"`
	Namespace   string               `json:"namespace"`
	Labels      map[string]string    `json:"labels"`
	Annotations map[string]string    `json:"annotations"`
	Type        string               `json:"type"`
	Selector    map[string]string    `json:"selector"`
	Ports       []corev1.ServicePort `json:"ports"`
	ClusterIP   string               `json:"cluster_ip"`
}

// CollectData collects service information and saves it to a JSON file
func CollectService(clientset *kubernetes.Clientset, resourceType string) {
	services, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing services: %v", err)
		return
	}

	var serviceInfos []ServiceInfo
	for _, service := range services.Items {
		serviceInfo := ServiceInfo{
			Name:        service.Name,
			Namespace:   service.Namespace,
			Labels:      service.Labels,
			Annotations: service.Annotations,
			Type:        string(service.Spec.Type),
			Selector:    service.Spec.Selector,
			Ports:       service.Spec.Ports,
			ClusterIP:   service.Spec.ClusterIP,
		}
		serviceInfos = append(serviceInfos, serviceInfo)
	}

	// Save the service information to a JSON file
	SaveData(resourceType, serviceInfos)
}
