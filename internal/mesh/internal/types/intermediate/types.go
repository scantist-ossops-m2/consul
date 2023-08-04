package intermediate

import (
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v1alpha1"
	pbmesh "github.com/hashicorp/consul/proto-public/pbmesh/v1alpha1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

// CombinedDestinationRef contains all references we need for a specific
// destination on the mesh.
type CombinedDestinationRef struct {
	// ServiceRef is the reference to the destination service.
	ServiceRef *pbresource.Reference

	// Port is the port name for this destination.
	Port string

	// SourceProxies are the IDs of source proxy state template resources.
	// The keys are a string representation of *pbresource.ID.
	SourceProxies map[string]*pbresource.ID

	// ExplicitDestinationsID is the id of the pbmesh.Upstreams resource. For implicit destinations,
	// this should be nil.
	ExplicitDestinationsID *pbresource.ID
}

type ServiceEndpoints struct {
	Resource  *pbresource.Resource
	Endpoints *pbcatalog.ServiceEndpoints
}

type Destinations struct {
	Resource     *pbresource.Resource
	Destinations *pbmesh.Upstreams
}

type Workload struct {
	Resource *pbresource.Resource
	Workload *pbcatalog.Workload
}

type ProxyStateTemplate struct {
	Resource *pbresource.Resource
	Tmpl     *pbmesh.ProxyStateTemplate
}

type ProxyConfiguration struct {
	Resource *pbresource.Resource
	Cfg      *pbmesh.ProxyConfiguration
}

type Destination struct {
	Explicit         *pbmesh.Upstream
	ServiceEndpoints *ServiceEndpoints
	Identities       []*pbresource.Reference
}

type Status struct {
	ID         *pbresource.ID
	Generation string
	Conditions []*pbresource.Condition
	OldStatus  map[string]*pbresource.Status
}
