package collectors

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// DeploymentInfo represents the structure of deployment data to be collected
type DeploymentInfo struct {
	Name            string            `json:"name"`
	Namespace       string            `json:"namespace"`
	Labels          map[string]string `json:"labels"`
	Annotations     map[string]string `json:"annotations"`
	Replicas        int32             `json:"replicas"`
	Selector        map[string]string `json:"selector"`
	Strategy        string            `json:"strategy"`
	RevisionHistory int32             `json:"revision_history"`
}

// CollectData collects deployment information and saves it to a JSON file
func CollectDeployment(clientset *kubernetes.Clientset, resourceType string) {
	deployments, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing deployments: %v", err)
		return
	}

	var deploymentInfos []DeploymentInfo
	for _, deployment := range deployments.Items {
		replicas := int32(1) // Default value if nil
		if deployment.Spec.Replicas != nil {
			replicas = *deployment.Spec.Replicas
		}

		revisionHistory := int32(10) // Default value if nil
		if deployment.Spec.RevisionHistoryLimit != nil {
			revisionHistory = *deployment.Spec.RevisionHistoryLimit
		}

		deploymentInfo := DeploymentInfo{
			Name:            deployment.Name,
			Namespace:       deployment.Namespace,
			Labels:          deployment.Labels,
			Annotations:     deployment.Annotations,
			Replicas:        replicas,
			Selector:        deployment.Spec.Selector.MatchLabels,
			Strategy:        string(deployment.Spec.Strategy.Type),
			RevisionHistory: revisionHistory,
		}
		deploymentInfos = append(deploymentInfos, deploymentInfo)
	}

	// Save the deployment information to a JSON file
	SaveData(resourceType, deploymentInfos)
}
