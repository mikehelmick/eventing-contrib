/*
Copyright 2019 The Knative Authors

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/eventing-contrib/couchdb/source/pkg/apis/sources/v1alpha1"
)

// FakeCouchDbSources implements CouchDbSourceInterface
type FakeCouchDbSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var couchdbsourcesResource = schema.GroupVersionResource{Group: "sources.eventing.knative.dev", Version: "v1alpha1", Resource: "couchdbsources"}

var couchdbsourcesKind = schema.GroupVersionKind{Group: "sources.eventing.knative.dev", Version: "v1alpha1", Kind: "CouchDbSource"}

// Get takes name of the couchDbSource, and returns the corresponding couchDbSource object, and an error if there is any.
func (c *FakeCouchDbSources) Get(name string, options v1.GetOptions) (result *v1alpha1.CouchDbSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(couchdbsourcesResource, c.ns, name), &v1alpha1.CouchDbSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CouchDbSource), err
}

// List takes label and field selectors, and returns the list of CouchDbSources that match those selectors.
func (c *FakeCouchDbSources) List(opts v1.ListOptions) (result *v1alpha1.CouchDbSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(couchdbsourcesResource, couchdbsourcesKind, c.ns, opts), &v1alpha1.CouchDbSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CouchDbSourceList{ListMeta: obj.(*v1alpha1.CouchDbSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.CouchDbSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested couchDbSources.
func (c *FakeCouchDbSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(couchdbsourcesResource, c.ns, opts))

}

// Create takes the representation of a couchDbSource and creates it.  Returns the server's representation of the couchDbSource, and an error, if there is any.
func (c *FakeCouchDbSources) Create(couchDbSource *v1alpha1.CouchDbSource) (result *v1alpha1.CouchDbSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(couchdbsourcesResource, c.ns, couchDbSource), &v1alpha1.CouchDbSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CouchDbSource), err
}

// Update takes the representation of a couchDbSource and updates it. Returns the server's representation of the couchDbSource, and an error, if there is any.
func (c *FakeCouchDbSources) Update(couchDbSource *v1alpha1.CouchDbSource) (result *v1alpha1.CouchDbSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(couchdbsourcesResource, c.ns, couchDbSource), &v1alpha1.CouchDbSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CouchDbSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCouchDbSources) UpdateStatus(couchDbSource *v1alpha1.CouchDbSource) (*v1alpha1.CouchDbSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(couchdbsourcesResource, "status", c.ns, couchDbSource), &v1alpha1.CouchDbSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CouchDbSource), err
}

// Delete takes name of the couchDbSource and deletes it. Returns an error if one occurs.
func (c *FakeCouchDbSources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(couchdbsourcesResource, c.ns, name), &v1alpha1.CouchDbSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCouchDbSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(couchdbsourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CouchDbSourceList{})
	return err
}

// Patch applies the patch and returns the patched couchDbSource.
func (c *FakeCouchDbSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CouchDbSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(couchdbsourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CouchDbSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CouchDbSource), err
}
