package robot

import (
	"context"
	"reflect"

	"github.com/robolaunch/kube-dev-suite/internal/node"
	"github.com/robolaunch/kube-dev-suite/internal/resources"
	robotv1alpha1 "github.com/robolaunch/kube-dev-suite/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *RobotReconciler) createPVC(ctx context.Context, instance *robotv1alpha1.Robot, pvcNamespacedName *types.NamespacedName) error {

	pvc := resources.GetPersistentVolumeClaim(instance, pvcNamespacedName)

	err := ctrl.SetControllerReference(instance, pvc, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, pvc)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: PVC " + pvc.Name + " is created.")
	return nil
}

func (r *RobotReconciler) createJob(ctx context.Context, instance *robotv1alpha1.Robot, jobNamespacedName *types.NamespacedName) error {

	activeNode, err := r.reconcileCheckNode(ctx, instance)
	if err != nil {
		return err
	}

	job := resources.GetLoaderJob(instance, jobNamespacedName, node.HasGPU(*activeNode))

	err = ctrl.SetControllerReference(instance, job, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, job)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: Job " + job.Name + " is created.")
	return nil
}

func (r *RobotReconciler) createRobotDevSuite(ctx context.Context, instance *robotv1alpha1.Robot, rdsNamespacedName *types.NamespacedName) error {

	robotDevSuite := resources.GetRobotDevSuite(instance, rdsNamespacedName)

	err := ctrl.SetControllerReference(instance, robotDevSuite, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, robotDevSuite)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: Robot dev suite " + robotDevSuite.Name + " is created.")
	return nil
}

func (r *RobotReconciler) createWorkspaceManager(ctx context.Context, instance *robotv1alpha1.Robot, wsmNamespacedName *types.NamespacedName) error {

	workspaceManager := resources.GetWorkspaceManager(instance, wsmNamespacedName)

	err := ctrl.SetControllerReference(instance, workspaceManager, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, workspaceManager)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: Workspace manager " + workspaceManager.Name + " is created.")
	return nil
}

func (r *RobotReconciler) createBuildManager(ctx context.Context, instance *robotv1alpha1.Robot) error {

	if reflect.DeepEqual(instance.Status.InitialBuildManagerStatus, robotv1alpha1.ManagerStatus{}) && !reflect.DeepEqual(instance.Spec.BuildManagerTemplate, robotv1alpha1.BuildManagerSpec{}) {
		buildManager := resources.GetBuildManager(instance, &types.NamespacedName{Namespace: instance.Namespace, Name: instance.Name + "-build"})

		err := ctrl.SetControllerReference(instance, buildManager, r.Scheme)
		if err != nil {
			return err
		}

		err = r.Create(ctx, buildManager)
		if err != nil && errors.IsAlreadyExists(err) {
			return nil
		} else if err != nil {
			return err
		}

		logger.Info("STATUS: Build manager " + buildManager.Name + " is created.")

		instance.Status.InitialBuildManagerStatus.Created = true
		instance.Status.InitialBuildManagerStatus.Name = instance.Name + "-build"
	}

	return nil
}
