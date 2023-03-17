package robot_ide

import (
	"context"

	"github.com/robolaunch/kube-dev-suite/internal/label"
	"github.com/robolaunch/kube-dev-suite/internal/resources"
	robotv1alpha1 "github.com/robolaunch/kube-dev-suite/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *RobotIDEReconciler) reconcileCreateService(ctx context.Context, instance *robotv1alpha1.RobotIDE) error {

	ideService := resources.GetRobotIDEService(instance, instance.GetRobotIDEServiceMetadata())

	err := ctrl.SetControllerReference(instance, ideService, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, ideService)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: IDE service is created.")

	return nil
}

func (r *RobotIDEReconciler) reconcileCreatePod(ctx context.Context, instance *robotv1alpha1.RobotIDE) error {

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

	activeNode, err := r.reconcileCheckNode(ctx, robot)
	if err != nil {
		return err
	}

	idePod := resources.GetRobotIDEPod(instance, instance.GetRobotIDEPodMetadata(), *robot, *robotVDI, *activeNode)

	err = ctrl.SetControllerReference(instance, idePod, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, idePod)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: IDE pod is created.")

	return nil
}

func (r *RobotIDEReconciler) reconcileCreateIngress(ctx context.Context, instance *robotv1alpha1.RobotIDE) error {

	robot, err := r.reconcileGetTargetRobot(ctx, instance)
	if err != nil {
		return err
	}

	ideIngress := resources.GetRobotIDEIngress(instance, instance.GetRobotIDEIngressMetadata(), *robot)

	err = ctrl.SetControllerReference(instance, ideIngress, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, ideIngress)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: IDE ingress is created.")

	return nil
}
