package collectors

import (
	"context"
	"log"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// RoleBindingInfo represents the structure of RoleBinding data to be collected
type RoleBindingInfo struct {
	Name      string           `json:"name"`
	Namespace string           `json:"namespace"`
	RoleRef   rbacv1.RoleRef   `json:"role_ref"`
	Subjects  []rbacv1.Subject `json:"subjects"`
}

// CollectData collects RoleBinding information and saves it to a JSON file
func CollectRoleBindings(clientset *kubernetes.Clientset, resourceType string) {
	roleBindings, err := clientset.RbacV1().RoleBindings("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing RoleBindings: %v", err)
		return
	}

	var roleBindingInfos []RoleBindingInfo
	for _, roleBinding := range roleBindings.Items {
		roleBindingInfo := RoleBindingInfo{
			Name:      roleBinding.Name,
			Namespace: roleBinding.Namespace,
			RoleRef:   roleBinding.RoleRef,
			Subjects:  roleBinding.Subjects,
		}
		roleBindingInfos = append(roleBindingInfos, roleBindingInfo)
	}

	// Save the RoleBinding information to a JSON file
	SaveData(resourceType, roleBindingInfos)
}
