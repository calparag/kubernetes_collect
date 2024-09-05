package collectors

import (
	"context"
	"log"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// RoleInfo represents the structure of Role data to be collected
type RoleInfo struct {
	Name      string              `json:"name"`
	Namespace string              `json:"namespace"`
	Rules     []rbacv1.PolicyRule `json:"rules"`
}

// CollectData collects Role information and saves it to a JSON file
func CollectRoles(clientset *kubernetes.Clientset, resourceType string) {
	roles, err := clientset.RbacV1().Roles("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing Roles: %v", err)
		return
	}

	var roleInfos []RoleInfo
	for _, role := range roles.Items {
		roleInfo := RoleInfo{
			Name:      role.Name,
			Namespace: role.Namespace,
			Rules:     role.Rules,
		}
		roleInfos = append(roleInfos, roleInfo)
	}

	// Save the Role information to a JSON file
	SaveData(resourceType, roleInfos)
}
