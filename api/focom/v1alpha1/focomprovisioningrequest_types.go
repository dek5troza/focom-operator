/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var _ webhook.Validator = &FocomProvisioningRequest{}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FocomProvisioningRequestSpec defines the desired state of FocomProvisioningRequest
type FocomProvisioningRequestSpec struct {
	// +kubebuilder:validation:Required
	OCloudId string `json:"oCloudId"`

	// +kubebuilder:validation:Required
	OCloudNamespace string `json:"oCloudNamespace"`

	// +kubebuilder:validation:Optional
	Name string `json:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Description string `json:"description,omitempty"`

	// +kubebuilder:validation:Required
	TemplateName string `json:"templateName"`

	// +kubebuilder:validation:Required
	TemplateVersion string `json:"templateVersion"`

	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Required
	TemplateParameters runtime.RawExtension `json:"templateParameters"`
}

// FocomProvisioningRequestStatus defines the observed state of FocomProvisioningRequest
type FocomProvisioningRequestStatus struct {
	Phase       string       `json:"phase,omitempty"`
	Message     string       `json:"message,omitempty"`
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
	// The name of the remote resource in the target cluster
	RemoteName string `json:"remoteName,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:webhook:groups=focom.oss.ericsson.com,versions=v1alpha1,resources=focomprovisioningrequests,verbs=create;update,admissionReviewVersions={v1,v1beta1},name=vfocomprovisioningrequest.kb.io,path=/validate-focom-oss-ericsson-com-v1alpha1-focomprovisioningrequest,mutating=false,failurePolicy=fail,sideEffects=None

// FocomProvisioningRequest is the Schema for the focomprovisioningrequests API
type FocomProvisioningRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FocomProvisioningRequestSpec   `json:"spec,omitempty"`
	Status FocomProvisioningRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FocomProvisioningRequestList contains a list of FocomProvisioningRequest
type FocomProvisioningRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FocomProvisioningRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FocomProvisioningRequest{}, &FocomProvisioningRequestList{})
}
func (r *FocomProvisioningRequest) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// ValidateCreate satisfies admission.Validator
func (r *FocomProvisioningRequest) ValidateCreate() (admission.Warnings, error) {
	// e.g. allow creation always:
	return nil, nil
}

// ValidateUpdate checks immutability or other logic
func (r *FocomProvisioningRequest) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	oldFpr, ok := old.(*FocomProvisioningRequest)
	if !ok {
		return nil, fmt.Errorf("cannot cast old object to FocomProvisioningRequest")
	}

	// If you want to forbid changes to templateParameters:
	if !rawExtensionEqual(oldFpr.Spec.TemplateParameters.Raw, r.Spec.TemplateParameters.Raw) {
		return nil, fmt.Errorf("templateParameters is immutable and cannot be changed after creation")
	}

	return nil, nil
}

// ValidateDelete satisfies admission.Validator
func (r *FocomProvisioningRequest) ValidateDelete() (admission.Warnings, error) {
	// allow all deletes:
	return nil, nil
}

// rawExtensionEqual is your utility
func rawExtensionEqual(a, b []byte) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	if len(a) != len(b) {
		// quick check
		return false
	}
	// Compare content by unmarshaling
	var objA interface{}
	var objB interface{}
	if err := json.Unmarshal(a, &objA); err != nil {
		return false
	}
	if err := json.Unmarshal(b, &objB); err != nil {
		return false
	}
	return fmt.Sprintf("%v", objA) == fmt.Sprintf("%v", objB)
}
