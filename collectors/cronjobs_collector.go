package collectors

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CronJobInfo represents the structure of CronJob data to be collected
type CronJobInfo struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Schedule    string `json:"schedule"`
	PodTemplate string `json:"pod_template"`
}

// CollectData collects CronJob information and saves it to a JSON file
func CollectCronJobs(clientset *kubernetes.Clientset, resourceType string) {
	cronJobs, err := clientset.BatchV1().CronJobs("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing CronJobs: %v", err)
		return
	}

	var cronJobInfos []CronJobInfo
	for _, cronJob := range cronJobs.Items {
		cronJobInfo := CronJobInfo{
			Name:        cronJob.Name,
			Namespace:   cronJob.Namespace,
			Schedule:    cronJob.Spec.Schedule,
			PodTemplate: cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Name, // Example, adjust as needed
		}
		cronJobInfos = append(cronJobInfos, cronJobInfo)
	}

	// Save the CronJob information to a JSON file
	SaveData(resourceType, cronJobInfos)
}
