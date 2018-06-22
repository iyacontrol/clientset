/*
Copyright The Kubernetes Authors.

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
	canary_v1 "github.com/iyacontrol/clientset/pkg/apis/canary/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCanaries implements CanaryInterface
type FakeCanaries struct {
	Fake *FakeCanaryV1
	ns   string
}

var canariesResource = schema.GroupVersionResource{Group: "canary.k8s.io", Version: "v1", Resource: "canaries"}

var canariesKind = schema.GroupVersionKind{Group: "canary.k8s.io", Version: "v1", Kind: "Canary"}

// Get takes name of the canary, and returns the corresponding canary object, and an error if there is any.
func (c *FakeCanaries) Get(name string, options v1.GetOptions) (result *canary_v1.Canary, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(canariesResource, c.ns, name), &canary_v1.Canary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*canary_v1.Canary), err
}

// List takes label and field selectors, and returns the list of Canaries that match those selectors.
func (c *FakeCanaries) List(opts v1.ListOptions) (result *canary_v1.CanaryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(canariesResource, canariesKind, c.ns, opts), &canary_v1.CanaryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &canary_v1.CanaryList{ListMeta: obj.(*canary_v1.CanaryList).ListMeta}
	for _, item := range obj.(*canary_v1.CanaryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested canaries.
func (c *FakeCanaries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(canariesResource, c.ns, opts))

}

// Create takes the representation of a canary and creates it.  Returns the server's representation of the canary, and an error, if there is any.
func (c *FakeCanaries) Create(canary *canary_v1.Canary) (result *canary_v1.Canary, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(canariesResource, c.ns, canary), &canary_v1.Canary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*canary_v1.Canary), err
}

// Update takes the representation of a canary and updates it. Returns the server's representation of the canary, and an error, if there is any.
func (c *FakeCanaries) Update(canary *canary_v1.Canary) (result *canary_v1.Canary, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(canariesResource, c.ns, canary), &canary_v1.Canary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*canary_v1.Canary), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCanaries) UpdateStatus(canary *canary_v1.Canary) (*canary_v1.Canary, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(canariesResource, "status", c.ns, canary), &canary_v1.Canary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*canary_v1.Canary), err
}

// Delete takes name of the canary and deletes it. Returns an error if one occurs.
func (c *FakeCanaries) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(canariesResource, c.ns, name), &canary_v1.Canary{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCanaries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(canariesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &canary_v1.CanaryList{})
	return err
}

// Patch applies the patch and returns the patched canary.
func (c *FakeCanaries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *canary_v1.Canary, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(canariesResource, c.ns, name, data, subresources...), &canary_v1.Canary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*canary_v1.Canary), err
}
