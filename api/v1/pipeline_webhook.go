/*

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
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var pipelinelog = logf.Log.WithName("pipeline-resource")

func (r *Pipeline) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-app-brainest-ml-v1-pipeline,mutating=true,failurePolicy=fail,groups=app.brainest.ml,resources=pipelines,verbs=create;update,versions=v1,name=mpipeline.kb.io

var _ webhook.Defaulter = &Pipeline{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Pipeline) Default() {
	pipelinelog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-app-brainest-ml-v1-pipeline,mutating=false,failurePolicy=fail,groups=app.brainest.ml,resources=pipelines,versions=v1,name=vpipeline.kb.io

var _ webhook.Validator = &Pipeline{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Pipeline) ValidateCreate() error {
	pipelinelog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Pipeline) ValidateUpdate(old runtime.Object) error {
	pipelinelog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Pipeline) ValidateDelete() error {
	pipelinelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
