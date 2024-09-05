package collectors

import (
	"context"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// DaemonSetInfo represents the structure of DaemonSet data to be collected
type DaemonSetInfo struct {
	Name           string                         `json:"name"`
	Namespace      string                         `json:"namespace"`
	Labels         map[string]string              `json:"labels"`
	Annotations    map[string]string              `json:"annotations"`
	Selector       metav1.LabelSelector           `json:"selector"`
	NodeSelector   map[string]string              `json:"node_selector"`
	UpdateStrategy appsv1.DaemonSetUpdateStrategy `json:"update_strategy"`
}

// CollectData collects DaemonSet information and saves it to a JSON file
func CollectDaemonSets(clientset *kubernetes.Clientset, resourceType string) {
	daemonSets, err := clientset.AppsV1().DaemonSets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing DaemonSets: %v", err)
		return
	}

	var daemonSetInfos []DaemonSetInfo
	for _, ds := range daemonSets.Items {
		daemonSetInfo := DaemonSetInfo{
			Name:           ds.Name,
			Namespace:      ds.Namespace,
			Labels:         ds.Labels,
			Annotations:    ds.Annotations,
			Selector:       *ds.Spec.Selector,
			NodeSelector:   ds.Spec.Template.Spec.NodeSelector,
			UpdateStrategy: ds.Spec.UpdateStrategy,
		}
		daemonSetInfos = append(daemonSetInfos, daemonSetInfo)
	}

	// Save the DaemonSet information to a JSON file
	SaveData(resourceType, daemonSetInfos)
}
