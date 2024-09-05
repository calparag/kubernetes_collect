package collectors

import (
	"context"
	"log"

	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextv1client "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

// CRDInfo represents the structure of CRD data to be collected
type CRDInfo struct {
	Name             string                             `json:"name"`
	GroupVersion     string                             `json:"group_version"`
	Scope            string                             `json:"scope"`
	ValidationSchema *apiextv1.CustomResourceValidation `json:"validation_schema,omitempty"`
}

// CollectCrds collects CRD information and saves it to a JSON file
func CollectCrds(config *rest.Config, resourceType string) {
	clientset, err := apiextv1client.NewForConfig(config) // Use the correct clientset
	if err != nil {
		log.Printf("Error creating API extensions client: %v", err)
		return
	}

	crds, err := clientset.ApiextensionsV1().CustomResourceDefinitions().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing CRDs: %v", err)
		return
	}

	var crdInfos []CRDInfo
	for _, crd := range crds.Items {
		// Iterate over versions to find the validation schema
		var validationSchema *apiextv1.CustomResourceValidation
		if len(crd.Spec.Versions) > 0 {
			// Assuming we take the first version's validation schema
			validationSchema = crd.Spec.Versions[0].Schema
		}

		crdInfo := CRDInfo{
			Name:             crd.Name,
			GroupVersion:     crd.Spec.Group + "/" + crd.Spec.Versions[0].Name,
			Scope:            string(crd.Spec.Scope),
			ValidationSchema: validationSchema,
		}
		crdInfos = append(crdInfos, crdInfo)
	}

	// Save the CRD information to a JSON file
	SaveData(resourceType, crdInfos)
}
