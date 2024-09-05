package collectors

import (
	"context"
	"log"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// IngressInfo represents the structure of Ingress data to be collected
type IngressInfo struct {
	Name           string                       `json:"name"`
	Namespace      string                       `json:"namespace"`
	Rules          []networkingv1.IngressRule   `json:"rules"`
	TLS            []networkingv1.IngressTLS    `json:"tls"`
	BackendService *networkingv1.IngressBackend `json:"backend_service,omitempty"`
}

// CollectIngress collects Ingress information and saves it to a JSON file
func CollectIngress(clientset *kubernetes.Clientset, resourceType string) {
	ingresses, err := clientset.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing Ingresses: %v", err)
		return
	}

	var ingressInfos []IngressInfo
	for _, ingress := range ingresses.Items {
		ingressInfo := IngressInfo{
			Name:           ingress.Name,
			Namespace:      ingress.Namespace,
			Rules:          ingress.Spec.Rules,
			TLS:            ingress.Spec.TLS,
			BackendService: ingress.Spec.DefaultBackend,
		}
		ingressInfos = append(ingressInfos, ingressInfo)
	}

	// Save the Ingress information to a JSON file
	SaveData(resourceType, ingressInfos)
}
