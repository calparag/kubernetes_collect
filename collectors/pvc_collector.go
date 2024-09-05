package collectors

import (
	"context"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// PVCInfo represents the structure of PVC data to be collected
type PVCInfo struct {
	Name         string                              `json:"name"`
	Namespace    string                              `json:"namespace"`
	StorageClass string                              `json:"storage_class"`
	AccessModes  []corev1.PersistentVolumeAccessMode `json:"access_modes"`
	Capacity     string                              `json:"capacity"`
	Status       string                              `json:"status"`
	BoundPV      string                              `json:"bound_pv"`
}

// CollectData collects PVC information and saves it to a JSON file
func CollectPvc(clientset *kubernetes.Clientset, resourceType string) {
	pvcs, err := clientset.CoreV1().PersistentVolumeClaims("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing PVCs: %v", err)
		return
	}

	var pvcInfos []PVCInfo
	for _, pvc := range pvcs.Items {
		storageClass := ""
		if pvc.Spec.StorageClassName != nil {
			storageClass = *pvc.Spec.StorageClassName
		}

		capacity := pvc.Status.Capacity.Storage().String()

		pvcInfo := PVCInfo{
			Name:         pvc.Name,
			Namespace:    pvc.Namespace,
			StorageClass: storageClass,
			AccessModes:  pvc.Spec.AccessModes,
			Capacity:     capacity,
			Status:       string(pvc.Status.Phase),
			BoundPV:      pvc.Spec.VolumeName,
		}
		pvcInfos = append(pvcInfos, pvcInfo)
	}

	// Save the PVC information to a JSON file
	SaveData(resourceType, pvcInfos)
}
