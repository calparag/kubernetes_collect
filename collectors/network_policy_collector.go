package collectors

import (
	"context"
	"log"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// NetworkPolicyInfo represents the structure of NetworkPolicy data to be collected
type NetworkPolicyInfo struct {
	Name        string                                  `json:"name"`
	Namespace   string                                  `json:"namespace"`
	PodSelector map[string]string                       `json:"pod_selector"`
	Ingress     []networkingv1.NetworkPolicyIngressRule `json:"ingress"`
	Egress      []networkingv1.NetworkPolicyEgressRule  `json:"egress"`
}

// CollectData collects NetworkPolicy information and saves it to a JSON file
func CollectNetworkPolicy(clientset *kubernetes.Clientset, resourceType string) {
	networkPolicies, err := clientset.NetworkingV1().NetworkPolicies("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing NetworkPolicies: %v", err)
		return
	}

	var networkPolicyInfos []NetworkPolicyInfo
	for _, networkPolicy := range networkPolicies.Items {
		networkPolicyInfo := NetworkPolicyInfo{
			Name:        networkPolicy.Name,
			Namespace:   networkPolicy.Namespace,
			PodSelector: networkPolicy.Spec.PodSelector.MatchLabels,
			Ingress:     networkPolicy.Spec.Ingress,
			Egress:      networkPolicy.Spec.Egress,
		}
		networkPolicyInfos = append(networkPolicyInfos, networkPolicyInfo)
	}

	// Save the NetworkPolicy information to a JSON file
	SaveData(resourceType, networkPolicyInfos)
}
