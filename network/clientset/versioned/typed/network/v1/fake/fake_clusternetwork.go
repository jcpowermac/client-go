// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	networkv1 "github.com/openshift/api/network/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterNetworks implements ClusterNetworkInterface
type FakeClusterNetworks struct {
	Fake *FakeNetworkV1
}

var clusternetworksResource = schema.GroupVersionResource{Group: "network.openshift.io", Version: "v1", Resource: "clusternetworks"}

var clusternetworksKind = schema.GroupVersionKind{Group: "network.openshift.io", Version: "v1", Kind: "ClusterNetwork"}

// Get takes name of the clusterNetwork, and returns the corresponding clusterNetwork object, and an error if there is any.
func (c *FakeClusterNetworks) Get(ctx context.Context, name string, options v1.GetOptions) (result *networkv1.ClusterNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusternetworksResource, name), &networkv1.ClusterNetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkv1.ClusterNetwork), err
}

// List takes label and field selectors, and returns the list of ClusterNetworks that match those selectors.
func (c *FakeClusterNetworks) List(ctx context.Context, opts v1.ListOptions) (result *networkv1.ClusterNetworkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusternetworksResource, clusternetworksKind, opts), &networkv1.ClusterNetworkList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkv1.ClusterNetworkList{ListMeta: obj.(*networkv1.ClusterNetworkList).ListMeta}
	for _, item := range obj.(*networkv1.ClusterNetworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterNetworks.
func (c *FakeClusterNetworks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusternetworksResource, opts))
}

// Create takes the representation of a clusterNetwork and creates it.  Returns the server's representation of the clusterNetwork, and an error, if there is any.
func (c *FakeClusterNetworks) Create(ctx context.Context, clusterNetwork *networkv1.ClusterNetwork, opts v1.CreateOptions) (result *networkv1.ClusterNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusternetworksResource, clusterNetwork), &networkv1.ClusterNetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkv1.ClusterNetwork), err
}

// Update takes the representation of a clusterNetwork and updates it. Returns the server's representation of the clusterNetwork, and an error, if there is any.
func (c *FakeClusterNetworks) Update(ctx context.Context, clusterNetwork *networkv1.ClusterNetwork, opts v1.UpdateOptions) (result *networkv1.ClusterNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusternetworksResource, clusterNetwork), &networkv1.ClusterNetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkv1.ClusterNetwork), err
}

// Delete takes name of the clusterNetwork and deletes it. Returns an error if one occurs.
func (c *FakeClusterNetworks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusternetworksResource, name, opts), &networkv1.ClusterNetwork{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterNetworks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusternetworksResource, listOpts)

	_, err := c.Fake.Invokes(action, &networkv1.ClusterNetworkList{})
	return err
}

// Patch applies the patch and returns the patched clusterNetwork.
func (c *FakeClusterNetworks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *networkv1.ClusterNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusternetworksResource, name, pt, data, subresources...), &networkv1.ClusterNetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkv1.ClusterNetwork), err
}
