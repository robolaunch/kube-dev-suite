package launch_manager

import (
	"context"

	"github.com/robolaunch/robot-operator/internal/label"
	"github.com/robolaunch/robot-operator/internal/resources"
	robotv1alpha1 "github.com/robolaunch/robot-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *LaunchManagerReconciler) createLaunchPod(ctx context.Context, instance *robotv1alpha1.LaunchManager) error {

	robot, err := r.reconcileGetTargetRobot(ctx, instance)
	if err != nil {
		return err
	}

	robotVDI := &robotv1alpha1.RobotVDI{}
	if label.GetTargetRobotVDI(instance) != "" {
		robotVDI, err = r.reconcileGetTargetRobotVDI(ctx, instance)
		if err != nil {
			return err
		}
	}

	buildManager, err := r.reconcileGetCurrentBuildManager(ctx, instance)
	if err != nil {
		return err
	}

	activeNode, err := r.reconcileCheckNode(ctx, robot)
	if err != nil {
		return err
	}

	launchPod := resources.GetLaunchPod(instance, instance.GetLaunchPodMetadata(), *robot, *buildManager, *robotVDI, *activeNode)
	if err != nil {
		return err
	}

	err = ctrl.SetControllerReference(instance, launchPod, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, launchPod)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: Launch pod is created.")
	return nil
}
