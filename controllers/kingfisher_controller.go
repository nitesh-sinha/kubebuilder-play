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
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	birdsv1beta1 "github.com/nitesh-sinha/kubebuilder-play/api/v1beta1"
)

// KingfisherReconciler reconciles a Kingfisher object
type KingfisherReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	// After we log something, we will flip this bit
	logged bool
}

//+kubebuilder:rbac:groups=birds.github.com,resources=kingfishers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=birds.github.com,resources=kingfishers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=birds.github.com,resources=kingfishers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Kingfisher object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *KingfisherReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logr := log.FromContext(ctx)

	// TODO(user): your logic here
	if !r.logged {
		logr.Info("Alright here we go!!!")
		r.logged = true
	}
	var kf birdsv1beta1.Kingfisher
	if err := r.Get(ctx, req.NamespacedName, &kf); err != nil {
		logr.Error(err, "Error fetching Kingfisher\n")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	kf.Status.Message = fmt.Sprintf("We'll let this %s Kingfisher fly higher!!", kf.Spec.Colour)
	//logr.Infof("%v", kf)
	if err := r.Status().Update(ctx, &kf); err != nil {
		logr.Error(err, "Error updating Kingfisher \n")
		return ctrl.Result{}, err
	}
	// if err := r.Update(ctx, &kf); err != nil {
	// 	logr.Error(err, "Error updating Kingfisher \n")
	// 	return ctrl.Result{}, err
	// }
	logr.Info("Successfully updated the Kingfisher Status field")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KingfisherReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&birdsv1beta1.Kingfisher{}).
		Complete(r)
}
