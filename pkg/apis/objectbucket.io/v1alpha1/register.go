// Package v1alpha1 contains API Schema definitions for the objectbucket v1alpha1 API group
// +k8s:deepcopy-gen=package,register
// +groupName=objectbucket.io
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	objectbucketio "github.com/kube-object-storage/lib-bucket-provisioner/pkg/apis/objectbucket.io"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: objectbucketio.GroupName, Version: "v1alpha1"}

const Version = "v1alpha1"

func GroupKindVersion(kind string) schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind(kind)
}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&ObjectBucketClaim{},
		&ObjectBucketClaimList{},
		&ObjectBucket{},
		&ObjectBucketList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
