package collectors

import (
	"context"
	"log"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ClusterRoleInfo represents the structure of ClusterRole data to be collected
type ClusterRoleInfo struct {
	Name  string              `json:"name"`
	Rules []rbacv1.PolicyRule `json:"rules"`
}

// CollectData collects ClusterRole information and saves it to a JSON file
func CollectClusterRoles(clientset *kubernetes.Clientset, resourceType string) {
	clusterRoles, err := clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing ClusterRoles: %v", err)
		return
	}

	var clusterRoleInfos []ClusterRoleInfo
	for _, clusterRole := range clusterRoles.Items {
		clusterRoleInfo := ClusterRoleInfo{
			Name:  clusterRole.Name,
			Rules: clusterRole.Rules,
		}
		clusterRoleInfos = append(clusterRoleInfos, clusterRoleInfo)
	}

	// Save the ClusterRole information to a JSON file
	SaveData(resourceType, clusterRoleInfos)
}
