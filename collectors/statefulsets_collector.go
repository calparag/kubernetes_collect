package collectors

import (
	"context"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// StatefulSetInfo represents the structure of StatefulSet data to be collected
type StatefulSetInfo struct {
	Name         string                         `json:"name"`
	Namespace    string                         `json:"namespace"`
	Labels       map[string]string              `json:"labels"`
	Annotations  map[string]string              `json:"annotations"`
	Replicas     int32                          `json:"replicas"`
	ServiceName  string                         `json:"service_name"`
	VolumeClaims []corev1.PersistentVolumeClaim `json:"volume_claims"`
}

// CollectData collects StatefulSet information and saves it to a JSON file
func CollectStatefulSet(clientset *kubernetes.Clientset, resourceType string) {
	statefulSets, err := clientset.AppsV1().StatefulSets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing StatefulSets: %v", err)
		return
	}

	var statefulSetInfos []StatefulSetInfo
	for _, ss := range statefulSets.Items {
		statefulSetInfo := StatefulSetInfo{
			Name:         ss.Name,
			Namespace:    ss.Namespace,
			Labels:       ss.Labels,
			Annotations:  ss.Annotations,
			Replicas:     *ss.Spec.Replicas,
			ServiceName:  ss.Spec.ServiceName,
			VolumeClaims: ss.Spec.VolumeClaimTemplates,
		}
		statefulSetInfos = append(statefulSetInfos, statefulSetInfo)
	}

	// Save the StatefulSet information to a JSON file
	SaveData(resourceType, statefulSetInfos)
}
