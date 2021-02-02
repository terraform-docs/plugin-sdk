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

package plugin

import (
	goplugin "github.com/hashicorp/go-plugin"

	"github.com/terraform-docs/plugin-sdk/print"
	"github.com/terraform-docs/plugin-sdk/terraform"
)

// server is an RPC server acting as a plugin.
type server struct {
	impl   plugin
	broker *goplugin.MuxBroker
}

// ServeOpts is an option for serving a plugin.
type ServeOpts struct {
	Name    string
	Version string
	Engine  print.Engine
}

// ExecuteArgs is the collection of arguments being sent by terraform-docs
// core while executing the plugin command.
type ExecuteArgs struct {
	Module   terraform.Module
	Settings *print.Settings
}

// Serve is a wrapper of plugin.Serve. This is entrypoint of all plugins.
func Serve(opts *ServeOpts) {
	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins: goplugin.PluginSet{
			"formatter": &formatterPlugin{
				impl: &base{
					name:    opts.Name,
					version: opts.Version,
					engine:  opts.Engine,
				},
			},
		},
	})
}

// Name returns the version of the plugin.
func (s *server) Name(args interface{}, resp *string) error {
	*resp = s.impl.Name()
	return nil
}

// Version returns the version of the plugin.
func (s *server) Version(args interface{}, resp *string) error {
	*resp = s.impl.Version()
	return nil
}

// Execute returns the generated output.
func (s *server) Execute(args ExecuteArgs, resp *string) error {
	*resp, _ = s.impl.Execute(args)
	return nil
}
