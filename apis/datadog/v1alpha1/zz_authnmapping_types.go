/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AuthnMappingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Identity provider key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The ID of a role to attach to all users with the corresponding key and value.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Identity provider value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AuthnMappingParameters struct {

	// Identity provider key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The ID of a role to attach to all users with the corresponding key and value.
	// +crossplane:generate:reference:type=Role
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a Role to populate role.
	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// Selector for a Role to populate role.
	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`

	// Identity provider value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// AuthnMappingSpec defines the desired state of AuthnMapping
type AuthnMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthnMappingParameters `json:"forProvider"`
}

// AuthnMappingStatus defines the observed state of AuthnMapping.
type AuthnMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthnMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AuthnMapping is the Schema for the AuthnMappings API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type AuthnMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.key)",message="key is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.value)",message="value is a required parameter"
	Spec   AuthnMappingSpec   `json:"spec"`
	Status AuthnMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthnMappingList contains a list of AuthnMappings
type AuthnMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthnMapping `json:"items"`
}

// Repository type metadata.
var (
	AuthnMapping_Kind             = "AuthnMapping"
	AuthnMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthnMapping_Kind}.String()
	AuthnMapping_KindAPIVersion   = AuthnMapping_Kind + "." + CRDGroupVersion.String()
	AuthnMapping_GroupVersionKind = CRDGroupVersion.WithKind(AuthnMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthnMapping{}, &AuthnMappingList{})
}
