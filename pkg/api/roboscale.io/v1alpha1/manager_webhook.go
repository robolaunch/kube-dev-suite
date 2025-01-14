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

package v1alpha1

import (
	"context"
	"errors"
	"net/url"
	"strings"

	"github.com/google/go-github/github"
	"github.com/robolaunch/robot-operator/internal"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// ********************************
// WorkspaceManager webhooks
// ********************************

// log is for logging in this package.
var workspacemanagerlog = logf.Log.WithName("workspacemanager-resource")

func (r *WorkspaceManager) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-robot-roboscale-io-v1alpha1-workspacemanager,mutating=true,failurePolicy=fail,sideEffects=None,groups=robot.roboscale.io,resources=workspacemanagers,verbs=create;update,versions=v1alpha1,name=mworkspacemanager.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &WorkspaceManager{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *WorkspaceManager) Default() {
	workspacemanagerlog.Info("default", "name", r.Name)
	// _ = r.setRepositoryInfo()
}

//+kubebuilder:webhook:path=/validate-robot-roboscale-io-v1alpha1-workspacemanager,mutating=false,failurePolicy=fail,sideEffects=None,groups=robot.roboscale.io,resources=workspacemanagers,verbs=create;update,versions=v1alpha1,name=vworkspacemanager.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &WorkspaceManager{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *WorkspaceManager) ValidateCreate() error {
	workspacemanagerlog.Info("validate create", "name", r.Name)

	err := r.checkTargetRobotLabel()
	if err != nil {
		return err
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *WorkspaceManager) ValidateUpdate(old runtime.Object) error {
	workspacemanagerlog.Info("validate update", "name", r.Name)

	err := r.checkTargetRobotLabel()
	if err != nil {
		return err
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *WorkspaceManager) ValidateDelete() error {
	workspacemanagerlog.Info("validate delete", "name", r.Name)
	return nil
}

func (r *WorkspaceManager) checkTargetRobotLabel() error {
	labels := r.GetLabels()

	if _, ok := labels[internal.TARGET_ROBOT_LABEL_KEY]; !ok {
		return errors.New("target robot label should be added with key " + internal.TARGET_ROBOT_LABEL_KEY)
	}

	return nil
}

func (r *WorkspaceManager) setRepositoryInfo() error {

	for k1, ws := range r.Spec.Workspaces {
		for k2, repo := range ws.Repositories {
			userOrOrg, repoName, err := getPathVariables(repo.URL)
			if err != nil {
				return err
			}

			repo.Owner = userOrOrg
			repo.Repo = repoName

			lastCommitHash, err := getLastCommitHash(repo)
			repo.Hash = lastCommitHash

			ws.Repositories[k2] = repo
		}
		r.Spec.Workspaces[k1] = ws
	}

	return nil

}

func getPathVariables(URL string) (string, string, error) {
	u, err := url.Parse(URL)
	if err != nil {
		return "", "", err
	}

	path := strings.Split(u.Path, "/")[1:]
	if len(path) != 2 {
		return "", "", errors.New("path is not right")
	}

	return path[0], path[1], nil
}

func getLastCommitHash(repository Repository) (string, error) {

	client := github.NewClient(nil)

	branch, _, err := client.Repositories.GetBranch(context.Background(), repository.Owner, repository.Repo, repository.Branch)
	if err != nil {
		return "", err
	}

	commitSHA := branch.Commit.SHA

	return *commitSHA, nil
}

// ********************************
// BuildManager webhooks
// ********************************

// log is for logging in this package.
var buildmanagerlog = logf.Log.WithName("buildmanager-resource")

func (r *BuildManager) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-robot-roboscale-io-v1alpha1-buildmanager,mutating=true,failurePolicy=fail,sideEffects=None,groups=robot.roboscale.io,resources=buildmanagers,verbs=create;update,versions=v1alpha1,name=mbuildmanager.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &BuildManager{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *BuildManager) Default() {
	buildmanagerlog.Info("default", "name", r.Name)
}

//+kubebuilder:webhook:path=/validate-robot-roboscale-io-v1alpha1-buildmanager,mutating=false,failurePolicy=fail,sideEffects=None,groups=robot.roboscale.io,resources=buildmanagers,verbs=create;update,versions=v1alpha1,name=vbuildmanager.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &BuildManager{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *BuildManager) ValidateCreate() error {
	buildmanagerlog.Info("validate create", "name", r.Name)

	err := r.checkTargetRobotLabel()
	if err != nil {
		return err
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *BuildManager) ValidateUpdate(old runtime.Object) error {
	buildmanagerlog.Info("validate update", "name", r.Name)

	err := r.checkTargetRobotLabel()
	if err != nil {
		return err
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *BuildManager) ValidateDelete() error {
	buildmanagerlog.Info("validate delete", "name", r.Name)
	return nil
}

func (r *BuildManager) checkTargetRobotLabel() error {
	labels := r.GetLabels()

	if _, ok := labels[internal.TARGET_ROBOT_LABEL_KEY]; !ok {
		return errors.New("target robot label should be added with key " + internal.TARGET_ROBOT_LABEL_KEY)
	}

	return nil
}

// ********************************
// LaunchManager webhooks
// ********************************

// log is for logging in this package.
var launchmanagerlog = logf.Log.WithName("launchmanager-resource")

func (r *LaunchManager) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-robot-roboscale-io-v1alpha1-launchmanager,mutating=true,failurePolicy=fail,sideEffects=None,groups=robot.roboscale.io,resources=launchmanagers,verbs=create;update,versions=v1alpha1,name=mlaunchmanager.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &LaunchManager{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *LaunchManager) Default() {
	launchmanagerlog.Info("default", "name", r.Name)
}

//+kubebuilder:webhook:path=/validate-robot-roboscale-io-v1alpha1-launchmanager,mutating=false,failurePolicy=fail,sideEffects=None,groups=robot.roboscale.io,resources=launchmanagers,verbs=create;update,versions=v1alpha1,name=vlaunchmanager.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &LaunchManager{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *LaunchManager) ValidateCreate() error {
	launchmanagerlog.Info("validate create", "name", r.Name)

	err := r.checkTargetRobotLabel()
	if err != nil {
		return err
	}

	err = r.checkTargetRobotVDILabel()
	if err != nil {
		return err
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *LaunchManager) ValidateUpdate(old runtime.Object) error {
	launchmanagerlog.Info("validate update", "name", r.Name)

	err := r.checkTargetRobotLabel()
	if err != nil {
		return err
	}

	err = r.checkTargetRobotVDILabel()
	if err != nil {
		return err
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *LaunchManager) ValidateDelete() error {
	launchmanagerlog.Info("validate delete", "name", r.Name)
	return nil
}

func (r *LaunchManager) checkTargetRobotLabel() error {
	labels := r.GetLabels()

	if _, ok := labels[internal.TARGET_ROBOT_LABEL_KEY]; !ok {
		return errors.New("target robot label should be added with key " + internal.TARGET_ROBOT_LABEL_KEY)
	}

	return nil
}

func (r *LaunchManager) checkTargetRobotVDILabel() error {
	labels := r.GetLabels()

	if r.Spec.Display {
		if _, ok := labels[internal.TARGET_VDI_LABEL_KEY]; !ok {
			return errors.New("target robot vdi label should be added with key " + internal.TARGET_VDI_LABEL_KEY)
		}
	}

	return nil
}
