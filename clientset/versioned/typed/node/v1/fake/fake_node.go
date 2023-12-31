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
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/edgelive/def/apis/node/v1"
	nodev1 "github.com/edgelive/def/applyconfiguration/node/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodes implements NodeInterface
type FakeNodes struct {
	Fake *FakeNodeV1
	ns   string
}

var nodesResource = v1.SchemeGroupVersion.WithResource("nodes")

var nodesKind = v1.SchemeGroupVersion.WithKind("Node")

// Get takes name of the node, and returns the corresponding node object, and an error if there is any.
func (c *FakeNodes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Node, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodesResource, c.ns, name), &v1.Node{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Node), err
}

// List takes label and field selectors, and returns the list of Nodes that match those selectors.
func (c *FakeNodes) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodesResource, nodesKind, c.ns, opts), &v1.NodeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.NodeList{ListMeta: obj.(*v1.NodeList).ListMeta}
	for _, item := range obj.(*v1.NodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodes.
func (c *FakeNodes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodesResource, c.ns, opts))

}

// Create takes the representation of a node and creates it.  Returns the server's representation of the node, and an error, if there is any.
func (c *FakeNodes) Create(ctx context.Context, node *v1.Node, opts metav1.CreateOptions) (result *v1.Node, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodesResource, c.ns, node), &v1.Node{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Node), err
}

// Update takes the representation of a node and updates it. Returns the server's representation of the node, and an error, if there is any.
func (c *FakeNodes) Update(ctx context.Context, node *v1.Node, opts metav1.UpdateOptions) (result *v1.Node, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodesResource, c.ns, node), &v1.Node{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Node), err
}

// Delete takes name of the node and deletes it. Returns an error if one occurs.
func (c *FakeNodes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(nodesResource, c.ns, name, opts), &v1.Node{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nodesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.NodeList{})
	return err
}

// Patch applies the patch and returns the patched node.
func (c *FakeNodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Node, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodesResource, c.ns, name, pt, data, subresources...), &v1.Node{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Node), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied node.
func (c *FakeNodes) Apply(ctx context.Context, node *nodev1.NodeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Node, err error) {
	if node == nil {
		return nil, fmt.Errorf("node provided to Apply must not be nil")
	}
	data, err := json.Marshal(node)
	if err != nil {
		return nil, err
	}
	name := node.Name
	if name == nil {
		return nil, fmt.Errorf("node.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodesResource, c.ns, *name, types.ApplyPatchType, data), &v1.Node{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Node), err
}
