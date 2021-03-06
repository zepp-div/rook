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
	"time"

	v1 "github.com/rook/rook/pkg/apis/edgefs.rook.io/v1"
	scheme "github.com/rook/rook/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SMBsGetter has a method to return a SMBInterface.
// A group's client should implement this interface.
type SMBsGetter interface {
	SMBs(namespace string) SMBInterface
}

// SMBInterface has methods to work with SMB resources.
type SMBInterface interface {
	Create(ctx context.Context, sMB *v1.SMB, opts metav1.CreateOptions) (*v1.SMB, error)
	Update(ctx context.Context, sMB *v1.SMB, opts metav1.UpdateOptions) (*v1.SMB, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SMB, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SMBList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SMB, err error)
	SMBExpansion
}

// sMBs implements SMBInterface
type sMBs struct {
	client rest.Interface
	ns     string
}

// newSMBs returns a SMBs
func newSMBs(c *EdgefsV1Client, namespace string) *sMBs {
	return &sMBs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sMB, and returns the corresponding sMB object, and an error if there is any.
func (c *sMBs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SMB, err error) {
	result = &v1.SMB{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("smbs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SMBs that match those selectors.
func (c *sMBs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SMBList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SMBList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("smbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sMBs.
func (c *sMBs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("smbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sMB and creates it.  Returns the server's representation of the sMB, and an error, if there is any.
func (c *sMBs) Create(ctx context.Context, sMB *v1.SMB, opts metav1.CreateOptions) (result *v1.SMB, err error) {
	result = &v1.SMB{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("smbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sMB).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sMB and updates it. Returns the server's representation of the sMB, and an error, if there is any.
func (c *sMBs) Update(ctx context.Context, sMB *v1.SMB, opts metav1.UpdateOptions) (result *v1.SMB, err error) {
	result = &v1.SMB{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("smbs").
		Name(sMB.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sMB).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sMB and deletes it. Returns an error if one occurs.
func (c *sMBs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("smbs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sMBs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("smbs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sMB.
func (c *sMBs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SMB, err error) {
	result = &v1.SMB{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("smbs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
