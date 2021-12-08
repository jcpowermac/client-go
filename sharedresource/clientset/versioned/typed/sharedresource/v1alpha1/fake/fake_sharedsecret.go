// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/openshift/api/sharedresource/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSharedSecrets implements SharedSecretInterface
type FakeSharedSecrets struct {
	Fake *FakeSharedresourceV1alpha1
}

var sharedsecretsResource = schema.GroupVersionResource{Group: "sharedresource.openshift.io", Version: "v1alpha1", Resource: "sharedsecrets"}

var sharedsecretsKind = schema.GroupVersionKind{Group: "sharedresource.openshift.io", Version: "v1alpha1", Kind: "SharedSecret"}

// Get takes name of the sharedSecret, and returns the corresponding sharedSecret object, and an error if there is any.
func (c *FakeSharedSecrets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SharedSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(sharedsecretsResource, name), &v1alpha1.SharedSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedSecret), err
}

// List takes label and field selectors, and returns the list of SharedSecrets that match those selectors.
func (c *FakeSharedSecrets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SharedSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(sharedsecretsResource, sharedsecretsKind, opts), &v1alpha1.SharedSecretList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SharedSecretList{ListMeta: obj.(*v1alpha1.SharedSecretList).ListMeta}
	for _, item := range obj.(*v1alpha1.SharedSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sharedSecrets.
func (c *FakeSharedSecrets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(sharedsecretsResource, opts))
}

// Create takes the representation of a sharedSecret and creates it.  Returns the server's representation of the sharedSecret, and an error, if there is any.
func (c *FakeSharedSecrets) Create(ctx context.Context, sharedSecret *v1alpha1.SharedSecret, opts v1.CreateOptions) (result *v1alpha1.SharedSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(sharedsecretsResource, sharedSecret), &v1alpha1.SharedSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedSecret), err
}

// Update takes the representation of a sharedSecret and updates it. Returns the server's representation of the sharedSecret, and an error, if there is any.
func (c *FakeSharedSecrets) Update(ctx context.Context, sharedSecret *v1alpha1.SharedSecret, opts v1.UpdateOptions) (result *v1alpha1.SharedSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(sharedsecretsResource, sharedSecret), &v1alpha1.SharedSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedSecret), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSharedSecrets) UpdateStatus(ctx context.Context, sharedSecret *v1alpha1.SharedSecret, opts v1.UpdateOptions) (*v1alpha1.SharedSecret, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(sharedsecretsResource, "status", sharedSecret), &v1alpha1.SharedSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedSecret), err
}

// Delete takes name of the sharedSecret and deletes it. Returns an error if one occurs.
func (c *FakeSharedSecrets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(sharedsecretsResource, name, opts), &v1alpha1.SharedSecret{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSharedSecrets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(sharedsecretsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SharedSecretList{})
	return err
}

// Patch applies the patch and returns the patched sharedSecret.
func (c *FakeSharedSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SharedSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(sharedsecretsResource, name, pt, data, subresources...), &v1alpha1.SharedSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedSecret), err
}
