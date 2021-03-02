/*
Copyright 2021 The terraform-docs Authors.

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

package terraform

import (
	"encoding/gob"
)

// Module represents a Terraform module. It consists of
//
// - Header       ('header' json key):        Module header found in shape of multi line '*.tf' comments or an entire file
// - Footer       ('footer' json key):        Module footer found in shape of multi line '*.tf' comments or an entire file
// - Inputs       ('inputs' json key):        List of input 'variables' extracted from the Terraform module .tf files
// - ModuleCalls  ('modules' json key):       List of 'modules' extracted from the Terraform module .tf files
// - Outputs      ('outputs' json key):       List of 'outputs' extracted from Terraform module .tf files
// - Providers    ('providers' json key):     List of 'providers' extracted from resources used in Terraform module
// - Requirements ('requirements' json key):  List of 'requirements' extracted from the Terraform module .tf files
// - Resources    ('resources' json key):     List of 'resources' extracted from the Terraform module .tf files
type Module interface {
	// HasHeader indicates if the module has header.
	HasHeader() bool

	// HasFooter indicates if the module has footer.
	// HasFooter() bool

	// HasInputs indicates if the module has inputs.
	HasInputs() bool

	// HasModuleCalls indicates if the module has inputs.
	HasModuleCalls() bool

	// HasOutputs indicates if the module has outputs.
	HasOutputs() bool

	// HasProviders indicates if the module has providers.
	HasProviders() bool

	// HasRequirements indicates if the module has requirements.
	HasRequirements() bool

	// HasResources indicates if the module has resources.
	HasResources() bool
}

// WithHeader adds header to Module.
func WithHeader(h string) SectionFn {
	return func(m *module) {
		m.Header = h
	}
}

// WithFooter adds footer to Module.
func WithFooter(h string) SectionFn {
	return func(m *module) {
		m.Footer = h
	}
}

// WithInputs adds inputs to Module.
func WithInputs(i []*Input) SectionFn {
	return func(m *module) {
		m.Inputs = i
	}
}

// WithModuleCalls adds Modulecalls to Module.
func WithModuleCalls(mc []*ModuleCall) SectionFn {
	return func(m *module) {
		m.ModuleCalls = mc
	}
}

// WithOutputs adds outputs to Module.
func WithOutputs(o []*Output) SectionFn {
	return func(m *module) {
		m.Outputs = o
	}
}

// WithProviders adds providers to Module.
func WithProviders(p []*Provider) SectionFn {
	return func(m *module) {
		m.Providers = p
	}
}

// WithRequirements adds requirements to Module.
func WithRequirements(r []*Requirement) SectionFn {
	return func(m *module) {
		m.Requirements = r
	}
}

// WithResources adds resources to Module.
func WithResources(r []*Resource) SectionFn {
	return func(m *module) {
		m.Resources = r
	}
}

// WithRequiredInputs adds required inputs to Module.
func WithRequiredInputs(ri []*Input) SectionFn {
	return func(m *module) {
		m.RequiredInputs = ri
	}
}

// WithOptionalInputs adds optional inputs to Module.
func WithOptionalInputs(oi []*Input) SectionFn {
	return func(m *module) {
		m.OptionalInputs = oi
	}
}

type module struct {
	Header       string
	Footer       string
	Inputs       []*Input
	ModuleCalls  []*ModuleCall
	Outputs      []*Output
	Providers    []*Provider
	Requirements []*Requirement
	Resources    []*Resource

	RequiredInputs []*Input
	OptionalInputs []*Input
}

// SectionFn is a callback function to add section data into module.
type SectionFn func(*module)

// NewModule creates a new module and adds underlying section data into it.
func NewModule(fns ...SectionFn) Module {
	m := &module{}
	for _, fn := range fns {
		fn(m)
	}
	return m
}

// HasHeader indicates if the module has header.
func (m *module) HasHeader() bool {
	return len(m.Header) > 0
}

// HasFoot indicates if the module has footer.
func (m *module) HasFooter() bool {
	return len(m.Footer) > 0
}

// HasInputs indicates if the module has inputs.
func (m *module) HasInputs() bool {
	return len(m.Inputs) > 0
}

// HasModuleCalls indicates if the module has modulecalls.
func (m *module) HasModuleCalls() bool {
	return len(m.ModuleCalls) > 0
}

// HasOutputs indicates if the module has outputs.
func (m *module) HasOutputs() bool {
	return len(m.Outputs) > 0
}

// HasProviders indicates if the module has providers.
func (m *module) HasProviders() bool {
	return len(m.Providers) > 0
}

// HasRequirements indicates if the module has requirements.
func (m *module) HasRequirements() bool {
	return len(m.Requirements) > 0
}

// HasResources indicates if the module has resources.
func (m *module) HasResources() bool {
	return len(m.Resources) > 0
}

func init() {
	gob.Register(&module{})
	gob.Register([]interface{}{})
	gob.Register(map[string]interface{}{})
}
