// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/mneverov/cluster-cidr-controller/pkg/apis/clustercidr/v1"
	scheme "github.com/mneverov/cluster-cidr-controller/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterCIDRsGetter has a method to return a ClusterCIDRInterface.
// A group's client should implement this interface.
type ClusterCIDRsGetter interface {
	ClusterCIDRs() ClusterCIDRInterface
}

// ClusterCIDRInterface has methods to work with ClusterCIDR resources.
type ClusterCIDRInterface interface {
	Create(ctx context.Context, clusterCIDR *v1.ClusterCIDR, opts metav1.CreateOptions) (*v1.ClusterCIDR, error)
	Update(ctx context.Context, clusterCIDR *v1.ClusterCIDR, opts metav1.UpdateOptions) (*v1.ClusterCIDR, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ClusterCIDR, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ClusterCIDRList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterCIDR, err error)
	ClusterCIDRExpansion
}

// clusterCIDRs implements ClusterCIDRInterface
type clusterCIDRs struct {
	client rest.Interface
}

// newClusterCIDRs returns a ClusterCIDRs
func newClusterCIDRs(c *NetworkingV1Client) *clusterCIDRs {
	return &clusterCIDRs{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterCIDR, and returns the corresponding clusterCIDR object, and an error if there is any.
func (c *clusterCIDRs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterCIDR, err error) {
	result = &v1.ClusterCIDR{}
	err = c.client.Get().
		Resource("clustercidrs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterCIDRs that match those selectors.
func (c *clusterCIDRs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterCIDRList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterCIDRList{}
	err = c.client.Get().
		Resource("clustercidrs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterCIDRs.
func (c *clusterCIDRs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clustercidrs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterCIDR and creates it.  Returns the server's representation of the clusterCIDR, and an error, if there is any.
func (c *clusterCIDRs) Create(ctx context.Context, clusterCIDR *v1.ClusterCIDR, opts metav1.CreateOptions) (result *v1.ClusterCIDR, err error) {
	result = &v1.ClusterCIDR{}
	err = c.client.Post().
		Resource("clustercidrs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterCIDR).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterCIDR and updates it. Returns the server's representation of the clusterCIDR, and an error, if there is any.
func (c *clusterCIDRs) Update(ctx context.Context, clusterCIDR *v1.ClusterCIDR, opts metav1.UpdateOptions) (result *v1.ClusterCIDR, err error) {
	result = &v1.ClusterCIDR{}
	err = c.client.Put().
		Resource("clustercidrs").
		Name(clusterCIDR.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterCIDR).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterCIDR and deletes it. Returns an error if one occurs.
func (c *clusterCIDRs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clustercidrs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterCIDRs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clustercidrs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterCIDR.
func (c *clusterCIDRs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterCIDR, err error) {
	result = &v1.ClusterCIDR{}
	err = c.client.Patch(pt).
		Resource("clustercidrs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
