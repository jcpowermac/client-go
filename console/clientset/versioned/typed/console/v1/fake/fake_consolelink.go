// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	consolev1 "github.com/openshift/api/console/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConsoleLinks implements ConsoleLinkInterface
type FakeConsoleLinks struct {
	Fake *FakeConsoleV1
}

var consolelinksResource = schema.GroupVersionResource{Group: "console.openshift.io", Version: "v1", Resource: "consolelinks"}

var consolelinksKind = schema.GroupVersionKind{Group: "console.openshift.io", Version: "v1", Kind: "ConsoleLink"}

// Get takes name of the consoleLink, and returns the corresponding consoleLink object, and an error if there is any.
func (c *FakeConsoleLinks) Get(ctx context.Context, name string, options v1.GetOptions) (result *consolev1.ConsoleLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(consolelinksResource, name), &consolev1.ConsoleLink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*consolev1.ConsoleLink), err
}

// List takes label and field selectors, and returns the list of ConsoleLinks that match those selectors.
func (c *FakeConsoleLinks) List(ctx context.Context, opts v1.ListOptions) (result *consolev1.ConsoleLinkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(consolelinksResource, consolelinksKind, opts), &consolev1.ConsoleLinkList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &consolev1.ConsoleLinkList{ListMeta: obj.(*consolev1.ConsoleLinkList).ListMeta}
	for _, item := range obj.(*consolev1.ConsoleLinkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested consoleLinks.
func (c *FakeConsoleLinks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(consolelinksResource, opts))
}

// Create takes the representation of a consoleLink and creates it.  Returns the server's representation of the consoleLink, and an error, if there is any.
func (c *FakeConsoleLinks) Create(ctx context.Context, consoleLink *consolev1.ConsoleLink, opts v1.CreateOptions) (result *consolev1.ConsoleLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(consolelinksResource, consoleLink), &consolev1.ConsoleLink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*consolev1.ConsoleLink), err
}

// Update takes the representation of a consoleLink and updates it. Returns the server's representation of the consoleLink, and an error, if there is any.
func (c *FakeConsoleLinks) Update(ctx context.Context, consoleLink *consolev1.ConsoleLink, opts v1.UpdateOptions) (result *consolev1.ConsoleLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(consolelinksResource, consoleLink), &consolev1.ConsoleLink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*consolev1.ConsoleLink), err
}

// Delete takes name of the consoleLink and deletes it. Returns an error if one occurs.
func (c *FakeConsoleLinks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(consolelinksResource, name, opts), &consolev1.ConsoleLink{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConsoleLinks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(consolelinksResource, listOpts)

	_, err := c.Fake.Invokes(action, &consolev1.ConsoleLinkList{})
	return err
}

// Patch applies the patch and returns the patched consoleLink.
func (c *FakeConsoleLinks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *consolev1.ConsoleLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(consolelinksResource, name, pt, data, subresources...), &consolev1.ConsoleLink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*consolev1.ConsoleLink), err
}
