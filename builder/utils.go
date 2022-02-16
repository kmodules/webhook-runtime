/*
Copyright AppsCode Inc. and Contributors

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
