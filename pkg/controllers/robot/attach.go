package robot

import (
	"context"
	"sort"

	"github.com/robolaunch/kube-dev-suite/internal"
	robotv1alpha1 "github.com/robolaunch/kube-dev-suite/pkg/api/roboscale.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *RobotReconciler) reconcileHandleAttachments(ctx context.Context, instance *robotv1alpha1.Robot) error {

	err := r.createBuildManager(ctx, instance)
	if err != nil {
		return err
	}

	switch instance.Spec.Development {
	case true:

		// attach development suite
		err := r.reconcileAttachDevObject(ctx, instance)
		if err != nil {
			return err
		}

	case false:

		// create workspace manager
		switch instance.Status.WorkspaceManagerStatus.Created {
		case true:

			switch instance.Status.WorkspaceManagerStatus.Status.Phase {
			case robotv1alpha1.WorkspaceManagerPhaseReady:

				// select attached build object
				err := r.reconcileAttachBuildObject(ctx, instance)
				if err != nil {
					return err
				}

			default:

				instance.Status.AttachedBuildObject = robotv1alpha1.AttachedBuildObject{}

			}

		case false:

			instance.Status.Phase = robotv1alpha1.RobotPhaseConfiguringWorkspaces
			err := r.createWorkspaceManager(ctx, instance, instance.GetWorkspaceManagerMetadata())
			if err != nil {
				return err
			}
			instance.Status.WorkspaceManagerStatus.Created = true

		}

	}

	return nil
}

func (r *RobotReconciler) reconcileAttachBuildObject(ctx context.Context, instance *robotv1alpha1.Robot) error {

	// Detach dev objects from robot
	instance.Status.AttachedDevObjects = []robotv1alpha1.AttachedDevObject{}

	// Get attached build objects for this robot
	requirements := []labels.Requirement{}
	newReq, err := labels.NewRequirement(internal.TARGET_ROBOT_LABEL_KEY, selection.In, []string{instance.Name})
	if err != nil {
		return err
	}
	requirements = append(requirements, *newReq)

	robotSelector := labels.NewSelector().Add(requirements...)

	buildManagerList := robotv1alpha1.BuildManagerList{}
	err = r.List(ctx, &buildManagerList, &client.ListOptions{Namespace: instance.Namespace, LabelSelector: robotSelector})
	if err != nil {
		return err
	}

	if len(buildManagerList.Items) == 0 {
		instance.Status.AttachedBuildObject.Reference = corev1.ObjectReference{}
		instance.Status.AttachedBuildObject.Status = robotv1alpha1.BuildManagerStatus{}
		return nil
	}

	// Sort attached build objects for this robot according to their creation timestamps
	sort.SliceStable(buildManagerList.Items[:], func(i, j int) bool {
		return buildManagerList.Items[i].CreationTimestamp.String() > buildManagerList.Items[j].CreationTimestamp.String()
	})

	selectedBuildManager := buildManagerList.Items[0]

	if instance.Status.AttachedBuildObject.Reference.Name != selectedBuildManager.Name {
		instance.Status.AttachedBuildObject.Status = robotv1alpha1.BuildManagerStatus{}
	}

	instance.Status.AttachedBuildObject.Reference = corev1.ObjectReference{
		Kind:            selectedBuildManager.Kind,
		Namespace:       selectedBuildManager.Namespace,
		Name:            selectedBuildManager.Name,
		UID:             selectedBuildManager.UID,
		APIVersion:      selectedBuildManager.APIVersion,
		ResourceVersion: selectedBuildManager.ResourceVersion,
	}

	return nil
}

func (r *RobotReconciler) reconcileAttachDevObject(ctx context.Context, instance *robotv1alpha1.Robot) error {

	// Detach build and launch objects from robot
	instance.Status.AttachedBuildObject = robotv1alpha1.AttachedBuildObject{}

	// Get attached dev objects for this robot
	requirements := []labels.Requirement{}
	newReq, err := labels.NewRequirement(internal.TARGET_ROBOT_LABEL_KEY, selection.In, []string{instance.Name})
	if err != nil {
		return err
	}
	requirements = append(requirements, *newReq)

	robotSelector := labels.NewSelector().Add(requirements...)

	robotDevSuiteList := robotv1alpha1.RobotDevSuiteList{}
	err = r.List(ctx, &robotDevSuiteList, &client.ListOptions{Namespace: instance.Namespace, LabelSelector: robotSelector})
	if err != nil {
		return err
	}

	if len(robotDevSuiteList.Items) == 0 {
		instance.Status.AttachedDevObjects = []robotv1alpha1.AttachedDevObject{}
		return nil
	}

	// Sort attached dev objects for this robot according to their creation timestamps
	sort.SliceStable(robotDevSuiteList.Items[:], func(i, j int) bool {
		return robotDevSuiteList.Items[i].CreationTimestamp.String() < robotDevSuiteList.Items[j].CreationTimestamp.String()
	})

	instance.Status.AttachedDevObjects = []robotv1alpha1.AttachedDevObject{}

	for _, rds := range robotDevSuiteList.Items {
		instance.Status.AttachedDevObjects = append(instance.Status.AttachedDevObjects, robotv1alpha1.AttachedDevObject{
			Reference: corev1.ObjectReference{
				Kind:            rds.Kind,
				Namespace:       rds.Namespace,
				Name:            rds.Name,
				UID:             rds.UID,
				APIVersion:      rds.APIVersion,
				ResourceVersion: rds.ResourceVersion,
			},
		})
	}

	return nil
}
