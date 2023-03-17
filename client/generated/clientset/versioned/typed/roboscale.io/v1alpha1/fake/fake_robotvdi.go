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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/robolaunch/kube-dev-suite/pkg/api/roboscale.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRobotVDIs implements RobotVDIInterface
type FakeRobotVDIs struct {
	Fake *FakeRoboscaleV1alpha1
	ns   string
}

var robotvdisResource = schema.GroupVersionResource{Group: "roboscale.io", Version: "v1alpha1", Resource: "robotvdis"}

var robotvdisKind = schema.GroupVersionKind{Group: "roboscale.io", Version: "v1alpha1", Kind: "RobotVDI"}

// Get takes name of the robotVDI, and returns the corresponding robotVDI object, and an error if there is any.
func (c *FakeRobotVDIs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RobotVDI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(robotvdisResource, c.ns, name), &v1alpha1.RobotVDI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RobotVDI), err
}

// List takes label and field selectors, and returns the list of RobotVDIs that match those selectors.
func (c *FakeRobotVDIs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RobotVDIList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(robotvdisResource, robotvdisKind, c.ns, opts), &v1alpha1.RobotVDIList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RobotVDIList{ListMeta: obj.(*v1alpha1.RobotVDIList).ListMeta}
	for _, item := range obj.(*v1alpha1.RobotVDIList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested robotVDIs.
func (c *FakeRobotVDIs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(robotvdisResource, c.ns, opts))

}

// Create takes the representation of a robotVDI and creates it.  Returns the server's representation of the robotVDI, and an error, if there is any.
func (c *FakeRobotVDIs) Create(ctx context.Context, robotVDI *v1alpha1.RobotVDI, opts v1.CreateOptions) (result *v1alpha1.RobotVDI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(robotvdisResource, c.ns, robotVDI), &v1alpha1.RobotVDI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RobotVDI), err
}

// Update takes the representation of a robotVDI and updates it. Returns the server's representation of the robotVDI, and an error, if there is any.
func (c *FakeRobotVDIs) Update(ctx context.Context, robotVDI *v1alpha1.RobotVDI, opts v1.UpdateOptions) (result *v1alpha1.RobotVDI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(robotvdisResource, c.ns, robotVDI), &v1alpha1.RobotVDI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RobotVDI), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRobotVDIs) UpdateStatus(ctx context.Context, robotVDI *v1alpha1.RobotVDI, opts v1.UpdateOptions) (*v1alpha1.RobotVDI, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(robotvdisResource, "status", c.ns, robotVDI), &v1alpha1.RobotVDI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RobotVDI), err
}

// Delete takes name of the robotVDI and deletes it. Returns an error if one occurs.
func (c *FakeRobotVDIs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(robotvdisResource, c.ns, name, opts), &v1alpha1.RobotVDI{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRobotVDIs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(robotvdisResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RobotVDIList{})
	return err
}

// Patch applies the patch and returns the patched robotVDI.
func (c *FakeRobotVDIs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RobotVDI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(robotvdisResource, c.ns, name, pt, data, subresources...), &v1alpha1.RobotVDI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RobotVDI), err
}
