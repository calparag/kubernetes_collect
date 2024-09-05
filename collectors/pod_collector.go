package collectors

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// PodInfo represents the structure of pod data to be collected
type PodInfo struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	Status      string            `json:"status"`
	Node        string            `json:"node"`
	Containers  []ContainerInfo   `json:"containers"`
	Volumes     []VolumeInfo      `json:"volumes"`
}

type ContainerInfo struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type VolumeInfo struct {
	Name string `json:"name"`
}

// CollectData collects pod information and saves it to a JSON file
func CollectPod(clientset *kubernetes.Clientset, resourceType string) {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing pods: %v", err)
		return
	}

	var podInfos []PodInfo
	for _, pod := range pods.Items {
		var containers []ContainerInfo
		for _, container := range pod.Spec.Containers {
			containers = append(containers, ContainerInfo{Name: container.Name, Image: container.Image})
		}

		var volumes []VolumeInfo
		for _, volume := range pod.Spec.Volumes {
			volumes = append(volumes, VolumeInfo{Name: volume.Name})
		}

		podInfo := PodInfo{
			Name:        pod.Name,
			Namespace:   pod.Namespace,
			Labels:      pod.Labels,
			Annotations: pod.Annotations,
			Status:      string(pod.Status.Phase),
			Node:        pod.Spec.NodeName,
			Containers:  containers,
			Volumes:     volumes,
		}
		podInfos = append(podInfos, podInfo)
	}

	// Save the pod information to a JSON file
	SaveData(resourceType, podInfos)
}
