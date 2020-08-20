// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "k8s.io/cnat-client-go/cnat/pkg/apis/cnat/v1alpha1"
)

// FakeAts implements AtInterface
type FakeAts struct {
	Fake *FakeCnatV1alpha1
	ns   string
}

var atsResource = schema.GroupVersionResource{Group: "cnat.programming-kubernetes.info", Version: "v1alpha1", Resource: "ats"}

var atsKind = schema.GroupVersionKind{Group: "cnat.programming-kubernetes.info", Version: "v1alpha1", Kind: "At"}

// Get takes name of the at, and returns the corresponding at object, and an error if there is any.
func (c *FakeAts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.At, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(atsResource, c.ns, name), &v1alpha1.At{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.At), err
}

// List takes label and field selectors, and returns the list of Ats that match those selectors.
func (c *FakeAts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AtList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(atsResource, atsKind, c.ns, opts), &v1alpha1.AtList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AtList{ListMeta: obj.(*v1alpha1.AtList).ListMeta}
	for _, item := range obj.(*v1alpha1.AtList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ats.
func (c *FakeAts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(atsResource, c.ns, opts))

}

// Create takes the representation of a at and creates it.  Returns the server's representation of the at, and an error, if there is any.
func (c *FakeAts) Create(ctx context.Context, at *v1alpha1.At, opts v1.CreateOptions) (result *v1alpha1.At, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(atsResource, c.ns, at), &v1alpha1.At{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.At), err
}

// Update takes the representation of a at and updates it. Returns the server's representation of the at, and an error, if there is any.
func (c *FakeAts) Update(ctx context.Context, at *v1alpha1.At, opts v1.UpdateOptions) (result *v1alpha1.At, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(atsResource, c.ns, at), &v1alpha1.At{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.At), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAts) UpdateStatus(ctx context.Context, at *v1alpha1.At, opts v1.UpdateOptions) (*v1alpha1.At, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(atsResource, "status", c.ns, at), &v1alpha1.At{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.At), err
}

// Delete takes name of the at and deletes it. Returns an error if one occurs.
func (c *FakeAts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(atsResource, c.ns, name), &v1alpha1.At{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(atsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AtList{})
	return err
}

// Patch applies the patch and returns the patched at.
func (c *FakeAts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.At, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(atsResource, c.ns, name, pt, data, subresources...), &v1alpha1.At{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.At), err
}