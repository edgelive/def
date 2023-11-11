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

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/edgelive/def/apis/recorder/v1"
	recorderv1 "github.com/edgelive/def/applyconfiguration/recorder/v1"
	scheme "github.com/edgelive/def/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RecordersGetter has a method to return a RecorderInterface.
// A group's client should implement this interface.
type RecordersGetter interface {
	Recorders(namespace string) RecorderInterface
}

// RecorderInterface has methods to work with Recorder resources.
type RecorderInterface interface {
	Create(ctx context.Context, recorder *v1.Recorder, opts metav1.CreateOptions) (*v1.Recorder, error)
	Update(ctx context.Context, recorder *v1.Recorder, opts metav1.UpdateOptions) (*v1.Recorder, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Recorder, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.RecorderList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Recorder, err error)
	Apply(ctx context.Context, recorder *recorderv1.RecorderApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Recorder, err error)
	RecorderExpansion
}

// recorders implements RecorderInterface
type recorders struct {
	client rest.Interface
	ns     string
}

// newRecorders returns a Recorders
func newRecorders(c *RecorderV1Client, namespace string) *recorders {
	return &recorders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the recorder, and returns the corresponding recorder object, and an error if there is any.
func (c *recorders) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Recorder, err error) {
	result = &v1.Recorder{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("recorders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Recorders that match those selectors.
func (c *recorders) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RecorderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RecorderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("recorders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested recorders.
func (c *recorders) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("recorders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a recorder and creates it.  Returns the server's representation of the recorder, and an error, if there is any.
func (c *recorders) Create(ctx context.Context, recorder *v1.Recorder, opts metav1.CreateOptions) (result *v1.Recorder, err error) {
	result = &v1.Recorder{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("recorders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(recorder).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a recorder and updates it. Returns the server's representation of the recorder, and an error, if there is any.
func (c *recorders) Update(ctx context.Context, recorder *v1.Recorder, opts metav1.UpdateOptions) (result *v1.Recorder, err error) {
	result = &v1.Recorder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("recorders").
		Name(recorder.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(recorder).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the recorder and deletes it. Returns an error if one occurs.
func (c *recorders) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("recorders").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *recorders) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("recorders").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched recorder.
func (c *recorders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Recorder, err error) {
	result = &v1.Recorder{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("recorders").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied recorder.
func (c *recorders) Apply(ctx context.Context, recorder *recorderv1.RecorderApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Recorder, err error) {
	if recorder == nil {
		return nil, fmt.Errorf("recorder provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(recorder)
	if err != nil {
		return nil, err
	}
	name := recorder.Name
	if name == nil {
		return nil, fmt.Errorf("recorder.Name must be provided to Apply")
	}
	result = &v1.Recorder{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("recorders").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
