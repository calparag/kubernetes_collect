package collectors

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// SecretInfo represents the structure of Secret data to be collected
type SecretInfo struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Type      string            `json:"type"`
	Data      map[string]string `json:"data"` // Base64 encoded data
}

// CollectData collects Secret information and saves it to a JSON file
func CollectSecret(clientset *kubernetes.Clientset, resourceType string) {
	secrets, err := clientset.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing Secrets: %v", err)
		return
	}

	var secretInfos []SecretInfo
	for _, secret := range secrets.Items {
		secretInfo := SecretInfo{
			Name:      secret.Name,
			Namespace: secret.Namespace,
			Type:      string(secret.Type),
			Data:      make(map[string]string),
		}
		for key, value := range secret.Data {
			secretInfo.Data[key] = string(value) // Note: data is base64-encoded
		}
		secretInfos = append(secretInfos, secretInfo)
	}

	// Save the Secret information to a JSON file
	SaveData(resourceType, secretInfos)
}
