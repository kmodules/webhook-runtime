package builder

import (
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

type GroupPrefix string

const (
	MutatorGroupPrefix   GroupPrefix = "mutators."
	ValidatorGroupPrefix GroupPrefix = "validators."
)
const (
	ResourceSuffixPlural   = "webhooks"
	ResourceSuffixSingular = "webhook"
)

func MutatorResource(gvk schema.GroupVersionKind) (plural schema.GroupVersionResource, singular string) {
	return schema.GroupVersionResource{
		Group:    string(MutatorGroupPrefix) + gvk.Group,
		Version:  "v1alpha1",
		Resource: strings.ToLower(gvk.Kind) + ResourceSuffixPlural,
	}, strings.ToLower(gvk.Kind) + ResourceSuffixSingular
}

func ValidatorResource(gvk schema.GroupVersionKind) (plural schema.GroupVersionResource, singular string) {
	return schema.GroupVersionResource{
		Group:    string(ValidatorGroupPrefix) + gvk.Group,
		Version:  "v1alpha1",
		Resource: strings.ToLower(gvk.Kind) + ResourceSuffixPlural,
	}, strings.ToLower(gvk.Kind) + ResourceSuffixSingular
}

func resource(prefix GroupPrefix, gvk schema.GroupVersionKind) (plural schema.GroupVersionResource, singular string) {
	return schema.GroupVersionResource{
		Group:    string(prefix) + gvk.Group,
		Version:  "v1alpha1",
		Resource: strings.ToLower(gvk.Kind) + ResourceSuffixPlural,
	}, strings.ToLower(gvk.Kind) + ResourceSuffixSingular
}
