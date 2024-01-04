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

package v1

import (
	"context"
	"strings"

	admissionv1 "k8s.io/api/admission/v1"
	admissionv1beta1 "k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
)

// Adapted from https://github.com/openshift/generic-admission-server/blob/master/pkg/registry/admissionreview/admission_review.go

type AdmissionHookFunc func(req *admissionv1.AdmissionRequest) *admissionv1.AdmissionResponse

type REST struct {
	hookFn AdmissionHookFunc
}

var (
	_ rest.Creater                  = &REST{}
	_ rest.Scoper                   = &REST{}
	_ rest.GroupVersionKindProvider = &REST{}
	_ rest.Storage                  = &REST{}
	_ rest.SingularNameProvider     = &REST{}
)

func NewREST(hookFn AdmissionHookFunc) *REST {
	return &REST{
		hookFn: hookFn,
	}
}

func (r *REST) New() runtime.Object {
	return &admissionv1beta1.AdmissionReview{}
}

func (r *REST) GroupVersionKind(_ schema.GroupVersion) schema.GroupVersionKind {
	return admissionv1beta1.SchemeGroupVersion.WithKind("AdmissionReview")
}

func (r *REST) GetSingularName() string {
	return strings.ToLower("AdmissionReview")
}

func (r *REST) NamespaceScoped() bool {
	return false
}

func (r *REST) Create(ctx context.Context, obj runtime.Object, _ rest.ValidateObjectFunc, _ *metav1.CreateOptions) (runtime.Object, error) {
	admissionReview := obj.(*admissionv1beta1.AdmissionReview)
	req := admissionReview.Request
	resp := r.hookFn(&admissionv1.AdmissionRequest{
		UID:                req.UID,
		Kind:               req.Kind,
		Resource:           req.Resource,
		SubResource:        req.SubResource,
		RequestKind:        req.RequestKind,
		RequestResource:    req.RequestResource,
		RequestSubResource: req.RequestSubResource,
		Name:               req.Name,
		Namespace:          req.Namespace,
		Operation:          admissionv1.Operation(req.Operation),
		UserInfo:           req.UserInfo,
		Object:             req.Object,
		OldObject:          req.OldObject,
		DryRun:             req.DryRun,
		Options:            req.Options,
	})
	respv1beta1 := &admissionv1beta1.AdmissionResponse{
		UID:              req.UID,
		Allowed:          resp.Allowed,
		Result:           resp.Result,
		Patch:            resp.Patch,
		PatchType:        nil,
		AuditAnnotations: resp.AuditAnnotations,
		Warnings:         resp.Warnings,
	}
	if resp.PatchType != nil {
		pt := admissionv1beta1.PatchType(*resp.PatchType)
		respv1beta1.PatchType = &pt
	}
	admissionReview.Response = respv1beta1
	return admissionReview, nil
}

func (r *REST) Destroy() {
}
