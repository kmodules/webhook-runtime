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
	hooks "kmodules.xyz/webhook-runtime/admission/v1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// WebhookBuilder builds a Webhook.
type WebhookBuilder struct {
	apiType             runtime.Object
	gk                  schema.GroupKind
	scheme              *runtime.Scheme
	customDefaulter     admission.CustomDefaulter
	customDefaulterOpts []admission.DefaulterOption
	customValidator     admission.CustomValidator
}

// WebhookManagedBy allows inform its Scheme and RESTMapper.
func WebhookManagedBy(s *runtime.Scheme) *WebhookBuilder {
	return &WebhookBuilder{scheme: s}
}

// TODO(droot): update the GoDoc for conversion.

// For takes a runtime.Object which should be a CR.
// If the given object implements the admission.Defaulter interface, a MutatingWebhook will be wired for this type.
// If the given object implements the admission.Validator interface, a ValidatingWebhook will be wired for this type.
func (blder *WebhookBuilder) For(apiType runtime.Object) *WebhookBuilder {
	blder.apiType = apiType
	return blder
}

// WithDefaulter takes an admission.CustomDefaulter interface, a MutatingWebhook with the provided opts (admission.DefaulterOption)
// will be wired for this type.
func (blder *WebhookBuilder) WithDefaulter(defaulter admission.CustomDefaulter, opts ...admission.DefaulterOption) *WebhookBuilder {
	blder.customDefaulter = defaulter
	blder.customDefaulterOpts = opts
	return blder
}

// WithValidator takes a admission.CustomValidator interface, a ValidatingWebhook will be wired for this type.
func (blder *WebhookBuilder) WithValidator(validator admission.CustomValidator) *WebhookBuilder {
	blder.customValidator = validator
	return blder
}

// Complete builds the webhook.
func (blder *WebhookBuilder) Complete() (hooks.AdmissionHook, hooks.AdmissionHook, error) {
	// Create webhook(s) for each type
	gvk, err := apiutil.GVKForObject(blder.apiType, blder.scheme)
	if err != nil {
		return nil, nil, err
	}
	blder.gk = gvk.GroupKind()

	return blder.registerDefaultingWebhook(), blder.registerValidatingWebhook(), nil
}

// registerDefaultingWebhook registers a defaulting webhook if th.
func (blder *WebhookBuilder) registerDefaultingWebhook() hooks.AdmissionHook {
	if blder.customValidator == nil {
		return nil
	}
	mwh := admission.WithCustomDefaulter(blder.scheme, blder.apiType, blder.customDefaulter, blder.customDefaulterOpts...)
	return &webhook{
		prefix: MutatorGroupPrefix,
		gk:     blder.gk,
		w:      mwh,
	}
}

func (blder *WebhookBuilder) registerValidatingWebhook() hooks.AdmissionHook {
	if blder.customValidator == nil {
		return nil
	}
	vwh := admission.WithCustomValidator(blder.scheme, blder.apiType, blder.customValidator)
	return &webhook{
		prefix: ValidatorGroupPrefix,
		gk:     blder.gk,
		w:      vwh,
	}
}
