/*
Copyright 2020 The KubeVirt Authors.

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

package v1beta1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	scheme "kubevirt.io/client-go/generated/containerized-data-importer/clientset/versioned/scheme"
	v1beta1 "kubevirt.io/containerized-data-importer/pkg/apis/core/v1beta1"
)

// DataVolumesGetter has a method to return a DataVolumeInterface.
// A group's client should implement this interface.
type DataVolumesGetter interface {
	DataVolumes(namespace string) DataVolumeInterface
}

// DataVolumeInterface has methods to work with DataVolume resources.
type DataVolumeInterface interface {
	Create(*v1beta1.DataVolume) (*v1beta1.DataVolume, error)
	Update(*v1beta1.DataVolume) (*v1beta1.DataVolume, error)
	UpdateStatus(*v1beta1.DataVolume) (*v1beta1.DataVolume, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.DataVolume, error)
	List(opts v1.ListOptions) (*v1beta1.DataVolumeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.DataVolume, err error)
	DataVolumeExpansion
}

// dataVolumes implements DataVolumeInterface
type dataVolumes struct {
	client rest.Interface
	ns     string
}

// newDataVolumes returns a DataVolumes
func newDataVolumes(c *CdiV1beta1Client, namespace string) *dataVolumes {
	return &dataVolumes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dataVolume, and returns the corresponding dataVolume object, and an error if there is any.
func (c *dataVolumes) Get(name string, options v1.GetOptions) (result *v1beta1.DataVolume, err error) {
	result = &v1beta1.DataVolume{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datavolumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DataVolumes that match those selectors.
func (c *dataVolumes) List(opts v1.ListOptions) (result *v1beta1.DataVolumeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.DataVolumeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datavolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dataVolumes.
func (c *dataVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("datavolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dataVolume and creates it.  Returns the server's representation of the dataVolume, and an error, if there is any.
func (c *dataVolumes) Create(dataVolume *v1beta1.DataVolume) (result *v1beta1.DataVolume, err error) {
	result = &v1beta1.DataVolume{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("datavolumes").
		Body(dataVolume).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dataVolume and updates it. Returns the server's representation of the dataVolume, and an error, if there is any.
func (c *dataVolumes) Update(dataVolume *v1beta1.DataVolume) (result *v1beta1.DataVolume, err error) {
	result = &v1beta1.DataVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datavolumes").
		Name(dataVolume.Name).
		Body(dataVolume).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dataVolumes) UpdateStatus(dataVolume *v1beta1.DataVolume) (result *v1beta1.DataVolume, err error) {
	result = &v1beta1.DataVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datavolumes").
		Name(dataVolume.Name).
		SubResource("status").
		Body(dataVolume).
		Do().
		Into(result)
	return
}

// Delete takes name of the dataVolume and deletes it. Returns an error if one occurs.
func (c *dataVolumes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datavolumes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dataVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datavolumes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dataVolume.
func (c *dataVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.DataVolume, err error) {
	result = &v1beta1.DataVolume{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("datavolumes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
