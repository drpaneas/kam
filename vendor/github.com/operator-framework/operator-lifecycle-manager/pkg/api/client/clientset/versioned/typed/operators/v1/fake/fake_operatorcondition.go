/*
Copyright Red Hat, Inc.

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

	operatorsv1 "github.com/operator-framework/api/pkg/operators/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOperatorConditions implements OperatorConditionInterface
type FakeOperatorConditions struct {
	Fake *FakeOperatorsV1
	ns   string
}

var operatorconditionsResource = schema.GroupVersionResource{Group: "operators.coreos.com", Version: "v1", Resource: "operatorconditions"}

var operatorconditionsKind = schema.GroupVersionKind{Group: "operators.coreos.com", Version: "v1", Kind: "OperatorCondition"}

// Get takes name of the operatorCondition, and returns the corresponding operatorCondition object, and an error if there is any.
func (c *FakeOperatorConditions) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorsv1.OperatorCondition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(operatorconditionsResource, c.ns, name), &operatorsv1.OperatorCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OperatorCondition), err
}

// List takes label and field selectors, and returns the list of OperatorConditions that match those selectors.
func (c *FakeOperatorConditions) List(ctx context.Context, opts v1.ListOptions) (result *operatorsv1.OperatorConditionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(operatorconditionsResource, operatorconditionsKind, c.ns, opts), &operatorsv1.OperatorConditionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorsv1.OperatorConditionList{ListMeta: obj.(*operatorsv1.OperatorConditionList).ListMeta}
	for _, item := range obj.(*operatorsv1.OperatorConditionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested operatorConditions.
func (c *FakeOperatorConditions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(operatorconditionsResource, c.ns, opts))

}

// Create takes the representation of a operatorCondition and creates it.  Returns the server's representation of the operatorCondition, and an error, if there is any.
func (c *FakeOperatorConditions) Create(ctx context.Context, operatorCondition *operatorsv1.OperatorCondition, opts v1.CreateOptions) (result *operatorsv1.OperatorCondition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(operatorconditionsResource, c.ns, operatorCondition), &operatorsv1.OperatorCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OperatorCondition), err
}

// Update takes the representation of a operatorCondition and updates it. Returns the server's representation of the operatorCondition, and an error, if there is any.
func (c *FakeOperatorConditions) Update(ctx context.Context, operatorCondition *operatorsv1.OperatorCondition, opts v1.UpdateOptions) (result *operatorsv1.OperatorCondition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(operatorconditionsResource, c.ns, operatorCondition), &operatorsv1.OperatorCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OperatorCondition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOperatorConditions) UpdateStatus(ctx context.Context, operatorCondition *operatorsv1.OperatorCondition, opts v1.UpdateOptions) (*operatorsv1.OperatorCondition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(operatorconditionsResource, "status", c.ns, operatorCondition), &operatorsv1.OperatorCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OperatorCondition), err
}

// Delete takes name of the operatorCondition and deletes it. Returns an error if one occurs.
func (c *FakeOperatorConditions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(operatorconditionsResource, c.ns, name, opts), &operatorsv1.OperatorCondition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOperatorConditions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(operatorconditionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &operatorsv1.OperatorConditionList{})
	return err
}

// Patch applies the patch and returns the patched operatorCondition.
func (c *FakeOperatorConditions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorsv1.OperatorCondition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(operatorconditionsResource, c.ns, name, pt, data, subresources...), &operatorsv1.OperatorCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OperatorCondition), err
}
