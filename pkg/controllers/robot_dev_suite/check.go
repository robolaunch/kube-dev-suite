package robot_dev_suite

import (
	"context"
	"reflect"

	robotv1alpha1 "github.com/robolaunch/kube-dev-suite/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
)

func (r *RobotDevSuiteReconciler) reconcileCheckRobotVDI(ctx context.Context, instance *robotv1alpha1.RobotDevSuite) error {

	robotVDIQuery := &robotv1alpha1.RobotVDI{}
	err := r.Get(ctx, *instance.GetRobotVDIMetadata(), robotVDIQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.RobotVDIStatus = robotv1alpha1.RobotVDIInstanceStatus{}
		} else {
			return err
		}
	} else {

		if instance.Spec.VDIEnabled {

			if !reflect.DeepEqual(instance.Spec.RobotVDITemplate, robotVDIQuery.Spec) {
				robotVDIQuery.Spec = instance.Spec.RobotVDITemplate
				err = r.Update(ctx, robotVDIQuery)
				if err != nil {
					return err
				}
			}

			instance.Status.RobotVDIStatus.Created = true
			instance.Status.RobotVDIStatus.Phase = robotVDIQuery.Status.Phase

		} else {

			err := r.Delete(ctx, robotVDIQuery)
			if err != nil {
				return err
			}

		}

	}

	return nil
}

func (r *RobotDevSuiteReconciler) reconcileCheckRobotIDE(ctx context.Context, instance *robotv1alpha1.RobotDevSuite) error {

	robotIDEQuery := &robotv1alpha1.RobotIDE{}
	err := r.Get(ctx, *instance.GetRobotIDEMetadata(), robotIDEQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.RobotIDEStatus = robotv1alpha1.RobotIDEInstanceStatus{}
		} else {
			return err
		}
	} else {

		if instance.Spec.IDEEnabled {

			if !reflect.DeepEqual(instance.Spec.RobotIDETemplate, robotIDEQuery.Spec) {
				robotIDEQuery.Spec = instance.Spec.RobotIDETemplate
				err = r.Update(ctx, robotIDEQuery)
				if err != nil {
					return err
				}
			}

			instance.Status.RobotIDEStatus.Created = true
			instance.Status.RobotIDEStatus.Phase = robotIDEQuery.Status.Phase

		} else {

			err := r.Delete(ctx, robotIDEQuery)
			if err != nil {
				return err
			}

		}

	}

	return nil
}
