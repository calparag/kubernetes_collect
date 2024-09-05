package collectors

import (
	"context"
	"log"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ClusterRoleBindingInfo represents the structure of ClusterRoleBinding data to be collected
type ClusterRoleBindingInfo struct {
	Name     string           `json:"name"`
	RoleRef  rbacv1.RoleRef   `json:"role_ref"`
	Subjects []rbacv1.Subject `json:"subjects"`
}

// CollectData collects ClusterRoleBinding information and saves it to a JSON file
func CollectClusterRoleBindings(clientset *kubernetes.Clientset, resourceType string) {
	clusterRoleBindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing ClusterRoleBindings: %v", err)
		return
	}

	var clusterRoleBindingInfos []ClusterRoleBindingInfo
	for _, clusterRoleBinding := range clusterRoleBindings.Items {
		clusterRoleBindingInfo := ClusterRoleBindingInfo{
			Name:     clusterRoleBinding.Name,
			RoleRef:  clusterRoleBinding.RoleRef,
			Subjects: clusterRoleBinding.Subjects,
		}
		clusterRoleBindingInfos = append(clusterRoleBindingInfos, clusterRoleBindingInfo)
	}

	// Save the ClusterRoleBinding information to a JSON file
	SaveData(resourceType, clusterRoleBindingInfos)
}
