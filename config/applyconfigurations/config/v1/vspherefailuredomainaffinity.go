// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// VSphereFailureDomainAffinityApplyConfiguration represents a declarative configuration of the VSphereFailureDomainAffinity type for use
// with apply.
type VSphereFailureDomainAffinityApplyConfiguration struct {
	VMGroup    *string `json:"vmGroup,omitempty"`
	HostGroup  *string `json:"hostGroup,omitempty"`
	VMHostRule *string `json:"vmHostRule,omitempty"`
}

// VSphereFailureDomainAffinityApplyConfiguration constructs a declarative configuration of the VSphereFailureDomainAffinity type for use with
// apply.
func VSphereFailureDomainAffinity() *VSphereFailureDomainAffinityApplyConfiguration {
	return &VSphereFailureDomainAffinityApplyConfiguration{}
}

// WithVMGroup sets the VMGroup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VMGroup field is set to the value of the last call.
func (b *VSphereFailureDomainAffinityApplyConfiguration) WithVMGroup(value string) *VSphereFailureDomainAffinityApplyConfiguration {
	b.VMGroup = &value
	return b
}

// WithHostGroup sets the HostGroup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostGroup field is set to the value of the last call.
func (b *VSphereFailureDomainAffinityApplyConfiguration) WithHostGroup(value string) *VSphereFailureDomainAffinityApplyConfiguration {
	b.HostGroup = &value
	return b
}

// WithVMHostRule sets the VMHostRule field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VMHostRule field is set to the value of the last call.
func (b *VSphereFailureDomainAffinityApplyConfiguration) WithVMHostRule(value string) *VSphereFailureDomainAffinityApplyConfiguration {
	b.VMHostRule = &value
	return b
}