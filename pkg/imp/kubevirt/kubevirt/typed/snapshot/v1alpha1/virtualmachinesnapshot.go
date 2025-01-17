// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	scheme "github.com/kthcloud/go-deploy/pkg/imp/kubevirt/kubevirt/scheme"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubevirt.io/api/snapshot/v1alpha1"
)

// VirtualMachineSnapshotsGetter has a method to return a VirtualMachineSnapshotInterface.
// A group's client should implement this interface.
type VirtualMachineSnapshotsGetter interface {
	VirtualMachineSnapshots(namespace string) VirtualMachineSnapshotInterface
}

// VirtualMachineSnapshotInterface has methods to work with VirtualMachineSnapshot resources.
type VirtualMachineSnapshotInterface interface {
	Create(ctx context.Context, virtualMachineSnapshot *v1alpha1.VirtualMachineSnapshot, opts v1.CreateOptions) (*v1alpha1.VirtualMachineSnapshot, error)
	Update(ctx context.Context, virtualMachineSnapshot *v1alpha1.VirtualMachineSnapshot, opts v1.UpdateOptions) (*v1alpha1.VirtualMachineSnapshot, error)
	UpdateStatus(ctx context.Context, virtualMachineSnapshot *v1alpha1.VirtualMachineSnapshot, opts v1.UpdateOptions) (*v1alpha1.VirtualMachineSnapshot, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VirtualMachineSnapshot, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VirtualMachineSnapshotList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineSnapshot, err error)
	VirtualMachineSnapshotExpansion
}

// virtualMachineSnapshots implements VirtualMachineSnapshotInterface
type virtualMachineSnapshots struct {
	client rest.Interface
	ns     string
}

// newVirtualMachineSnapshots returns a VirtualMachineSnapshots
func newVirtualMachineSnapshots(c *SnapshotV1alpha1Client, namespace string) *virtualMachineSnapshots {
	return &virtualMachineSnapshots{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualMachineSnapshot, and returns the corresponding virtualMachineSnapshot object, and an error if there is any.
func (c *virtualMachineSnapshots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineSnapshot, err error) {
	result = &v1alpha1.VirtualMachineSnapshot{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinesnapshots").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineSnapshots that match those selectors.
func (c *virtualMachineSnapshots) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualMachineSnapshotList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualMachineSnapshotList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinesnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineSnapshots.
func (c *virtualMachineSnapshots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinesnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a virtualMachineSnapshot and creates it.  Returns the server's representation of the virtualMachineSnapshot, and an error, if there is any.
func (c *virtualMachineSnapshots) Create(ctx context.Context, virtualMachineSnapshot *v1alpha1.VirtualMachineSnapshot, opts v1.CreateOptions) (result *v1alpha1.VirtualMachineSnapshot, err error) {
	result = &v1alpha1.VirtualMachineSnapshot{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachinesnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineSnapshot).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a virtualMachineSnapshot and updates it. Returns the server's representation of the virtualMachineSnapshot, and an error, if there is any.
func (c *virtualMachineSnapshots) Update(ctx context.Context, virtualMachineSnapshot *v1alpha1.VirtualMachineSnapshot, opts v1.UpdateOptions) (result *v1alpha1.VirtualMachineSnapshot, err error) {
	result = &v1alpha1.VirtualMachineSnapshot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinesnapshots").
		Name(virtualMachineSnapshot.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineSnapshot).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *virtualMachineSnapshots) UpdateStatus(ctx context.Context, virtualMachineSnapshot *v1alpha1.VirtualMachineSnapshot, opts v1.UpdateOptions) (result *v1alpha1.VirtualMachineSnapshot, err error) {
	result = &v1alpha1.VirtualMachineSnapshot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinesnapshots").
		Name(virtualMachineSnapshot.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineSnapshot).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the virtualMachineSnapshot and deletes it. Returns an error if one occurs.
func (c *virtualMachineSnapshots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinesnapshots").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineSnapshots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinesnapshots").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched virtualMachineSnapshot.
func (c *virtualMachineSnapshots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineSnapshot, err error) {
	result = &v1alpha1.VirtualMachineSnapshot{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachinesnapshots").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
