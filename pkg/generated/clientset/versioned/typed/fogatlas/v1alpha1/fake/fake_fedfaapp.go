/*
Copyright FBK.

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

	v1alpha1 "github.com/fogatlas/crd-client-go/pkg/apis/fogatlas/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFedFAApps implements FedFAAppInterface
type FakeFedFAApps struct {
	Fake *FakeFogatlasV1alpha1
	ns   string
}

var fedfaappsResource = schema.GroupVersionResource{Group: "fogatlas.fbk.eu", Version: "v1alpha1", Resource: "fedfaapps"}

var fedfaappsKind = schema.GroupVersionKind{Group: "fogatlas.fbk.eu", Version: "v1alpha1", Kind: "FedFAApp"}

// Get takes name of the fedFAApp, and returns the corresponding fedFAApp object, and an error if there is any.
func (c *FakeFedFAApps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FedFAApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fedfaappsResource, c.ns, name), &v1alpha1.FedFAApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FedFAApp), err
}

// List takes label and field selectors, and returns the list of FedFAApps that match those selectors.
func (c *FakeFedFAApps) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FedFAAppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fedfaappsResource, fedfaappsKind, c.ns, opts), &v1alpha1.FedFAAppList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FedFAAppList{ListMeta: obj.(*v1alpha1.FedFAAppList).ListMeta}
	for _, item := range obj.(*v1alpha1.FedFAAppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fedFAApps.
func (c *FakeFedFAApps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fedfaappsResource, c.ns, opts))

}

// Create takes the representation of a fedFAApp and creates it.  Returns the server's representation of the fedFAApp, and an error, if there is any.
func (c *FakeFedFAApps) Create(ctx context.Context, fedFAApp *v1alpha1.FedFAApp, opts v1.CreateOptions) (result *v1alpha1.FedFAApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fedfaappsResource, c.ns, fedFAApp), &v1alpha1.FedFAApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FedFAApp), err
}

// Update takes the representation of a fedFAApp and updates it. Returns the server's representation of the fedFAApp, and an error, if there is any.
func (c *FakeFedFAApps) Update(ctx context.Context, fedFAApp *v1alpha1.FedFAApp, opts v1.UpdateOptions) (result *v1alpha1.FedFAApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fedfaappsResource, c.ns, fedFAApp), &v1alpha1.FedFAApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FedFAApp), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFedFAApps) UpdateStatus(ctx context.Context, fedFAApp *v1alpha1.FedFAApp, opts v1.UpdateOptions) (*v1alpha1.FedFAApp, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fedfaappsResource, "status", c.ns, fedFAApp), &v1alpha1.FedFAApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FedFAApp), err
}

// Delete takes name of the fedFAApp and deletes it. Returns an error if one occurs.
func (c *FakeFedFAApps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(fedfaappsResource, c.ns, name), &v1alpha1.FedFAApp{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFedFAApps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fedfaappsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FedFAAppList{})
	return err
}

// Patch applies the patch and returns the patched fedFAApp.
func (c *FakeFedFAApps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FedFAApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fedfaappsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FedFAApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FedFAApp), err
}
