package collectors

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// PVInfo represents the structure of PV data to be collected
type PVInfo struct {
	Name          string   `json:"name"`
	Capacity      string   `json:"capacity"`
	AccessModes   []string `json:"access_modes"`
	ReclaimPolicy string   `json:"reclaim_policy"`
	StorageClass  string   `json:"storage_class"`
	Status        string   `json:"status"`
	Claim         string   `json:"claim,omitempty"`
}

// CollectData collects PV information and saves it to a JSON file
func CollectPV(clientset *kubernetes.Clientset, resourceType string) {
	pvs, err := clientset.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing PVs: %v", err)
		return
	}

	var pvInfos []PVInfo
	for _, pv := range pvs.Items {
		claim := ""
		if pv.Spec.ClaimRef != nil {
			claim = pv.Spec.ClaimRef.Namespace + "/" + pv.Spec.ClaimRef.Name
		}
		pvInfo := PVInfo{
			Name:          pv.Name,
			Capacity:      pv.Spec.Capacity.Storage().String(),
			AccessModes:   make([]string, len(pv.Spec.AccessModes)),
			ReclaimPolicy: string(pv.Spec.PersistentVolumeReclaimPolicy),
			StorageClass:  pv.Spec.StorageClassName,
			Status:        string(pv.Status.Phase),
			Claim:         claim,
		}
		for i, mode := range pv.Spec.AccessModes {
			pvInfo.AccessModes[i] = string(mode)
		}
		pvInfos = append(pvInfos, pvInfo)
	}

	// Save the PV information to a JSON file
	SaveData(resourceType, pvInfos)
}
