/*
Copyright 2020 The Knative Authors

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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "knative.dev/eventing-ceph/pkg/apis/sources/v1alpha1"
	scheme "knative.dev/eventing-ceph/pkg/client/clientset/versioned/scheme"
)

// CephSourcesGetter has a method to return a CephSourceInterface.
// A group's client should implement this interface.
type CephSourcesGetter interface {
	CephSources(namespace string) CephSourceInterface
}

// CephSourceInterface has methods to work with CephSource resources.
type CephSourceInterface interface {
	Create(ctx context.Context, cephSource *v1alpha1.CephSource, opts v1.CreateOptions) (*v1alpha1.CephSource, error)
	Update(ctx context.Context, cephSource *v1alpha1.CephSource, opts v1.UpdateOptions) (*v1alpha1.CephSource, error)
	UpdateStatus(ctx context.Context, cephSource *v1alpha1.CephSource, opts v1.UpdateOptions) (*v1alpha1.CephSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CephSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CephSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CephSource, err error)
	CephSourceExpansion
}

// cephSources implements CephSourceInterface
type cephSources struct {
	client rest.Interface
	ns     string
}

// newCephSources returns a CephSources
func newCephSources(c *SourcesV1alpha1Client, namespace string) *cephSources {
	return &cephSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cephSource, and returns the corresponding cephSource object, and an error if there is any.
func (c *cephSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CephSource, err error) {
	result = &v1alpha1.CephSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cephsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CephSources that match those selectors.
func (c *cephSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CephSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CephSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cephsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cephSources.
func (c *cephSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cephsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cephSource and creates it.  Returns the server's representation of the cephSource, and an error, if there is any.
func (c *cephSources) Create(ctx context.Context, cephSource *v1alpha1.CephSource, opts v1.CreateOptions) (result *v1alpha1.CephSource, err error) {
	result = &v1alpha1.CephSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cephsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cephSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cephSource and updates it. Returns the server's representation of the cephSource, and an error, if there is any.
func (c *cephSources) Update(ctx context.Context, cephSource *v1alpha1.CephSource, opts v1.UpdateOptions) (result *v1alpha1.CephSource, err error) {
	result = &v1alpha1.CephSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cephsources").
		Name(cephSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cephSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cephSources) UpdateStatus(ctx context.Context, cephSource *v1alpha1.CephSource, opts v1.UpdateOptions) (result *v1alpha1.CephSource, err error) {
	result = &v1alpha1.CephSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cephsources").
		Name(cephSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cephSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cephSource and deletes it. Returns an error if one occurs.
func (c *cephSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cephsources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cephSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cephsources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cephSource.
func (c *cephSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CephSource, err error) {
	result = &v1alpha1.CephSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cephsources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}