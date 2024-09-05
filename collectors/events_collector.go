package collectors

import (
	"context"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// EventInfo represents the structure of Event data to be collected
type EventInfo struct {
	EventType      string                 `json:"event_type"`
	Reason         string                 `json:"reason"`
	Message        string                 `json:"message"`
	InvolvedObject corev1.ObjectReference `json:"involved_object"`
	Timestamp      metav1.Time            `json:"timestamp"`
}

// CollectData collects Event information and saves it to a JSON file
func CollectEvents(clientset *kubernetes.Clientset, resourceType string) {
	events, err := clientset.CoreV1().Events("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing Events: %v", err)
		return
	}

	var eventInfos []EventInfo
	for _, event := range events.Items {
		eventInfo := EventInfo{
			EventType:      event.Type,
			Reason:         event.Reason,
			Message:        event.Message,
			InvolvedObject: event.InvolvedObject,
			Timestamp:      event.LastTimestamp,
		}
		eventInfos = append(eventInfos, eventInfo)
	}

	// Save the Event information to a JSON file
	SaveData(resourceType, eventInfos)
}
