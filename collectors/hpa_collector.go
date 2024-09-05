package collectors

import (
	"context"
	"log"

	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// HPAInfo represents the structure of HPA data to be collected
type HPAInfo struct {
	Name        string                                    `json:"name"`
	Namespace   string                                    `json:"namespace"`
	Target      autoscalingv1.CrossVersionObjectReference `json:"target"`
	MinReplicas *int32                                    `json:"min_replicas"`
	MaxReplicas int32                                     `json:"max_replicas"`
	Metrics     []autoscalingv1.MetricSpec                `json:"metrics"`
}

// CollectData collects HPA information and saves it to a JSON file
func CollectHpa(clientset *kubernetes.Clientset, resourceType string) {
	hpas, err := clientset.AutoscalingV1().HorizontalPodAutoscalers("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing HPAs: %v", err)
		return
	}

	var hpaInfos []HPAInfo
	for _, hpa := range hpas.Items {
		hpaInfo := HPAInfo{
			Name:        hpa.Name,
			Namespace:   hpa.Namespace,
			Target:      hpa.Spec.ScaleTargetRef,
			MinReplicas: hpa.Spec.MinReplicas,
			MaxReplicas: hpa.Spec.MaxReplicas,
		}
		hpaInfos = append(hpaInfos, hpaInfo)
	}

	// Save the HPA information to a JSON file
	SaveData(resourceType, hpaInfos)
}
