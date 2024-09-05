package collectors

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// JobInfo represents the structure of Job data to be collected
type JobInfo struct {
	Name                 string `json:"name"`
	Namespace            string `json:"namespace"`
	SuccessfulExecutions int32  `json:"successful_executions"`
	FailedExecutions     int32  `json:"failed_executions"`
	PodTemplate          string `json:"pod_template"`
}

// CollectData collects Job information and saves it to a JSON file
func CollectJobs(clientset *kubernetes.Clientset, resourceType string) {
	jobs, err := clientset.BatchV1().Jobs("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing Jobs: %v", err)
		return
	}

	var jobInfos []JobInfo
	for _, job := range jobs.Items {
		jobInfo := JobInfo{
			Name:                 job.Name,
			Namespace:            job.Namespace,
			SuccessfulExecutions: job.Status.Succeeded,
			FailedExecutions:     job.Status.Failed,
			PodTemplate:          job.Spec.Template.Spec.Containers[0].Name, // Example, adjust as needed
		}
		jobInfos = append(jobInfos, jobInfo)
	}

	// Save the Job information to a JSON file
	SaveData(resourceType, jobInfos)
}
