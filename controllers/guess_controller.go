/*
Copyright 2022.

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

package controllers

import (
	"context"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	jugsaxonyv1beta1 "github.com/rpahli/jugsaxony-operator/api/v1beta1"
)

// GuessReconciler reconciles a Guess object
type GuessReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demo.jugsaxony.org,resources=guesses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.jugsaxony.org,resources=guesses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.jugsaxony.org,resources=guesses/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Guess object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *GuessReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ret ctrl.Result, reterr error) {
	_ = log.FromContext(ctx)

	// your logic here

	guess := &jugsaxonyv1beta1.Guess{}

	if err := r.Client.Get(ctx, req.NamespacedName, guess); err != nil {
		if apierrors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return ctrl.Result{}, err
	}

	if guess.Labels == nil {
		guess.Labels = map[string]string{"jug-game": guess.Spec.Game}
		reterr = r.Update(ctx, guess)
	}

	return ctrl.Result{}, reterr
}

// SetupWithManager sets up the controller with the Manager.
func (r *GuessReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&jugsaxonyv1beta1.Guess{}).
		Complete(r)
}
