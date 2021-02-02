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
	"net/rpc"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	goplugin "github.com/hashicorp/go-plugin"
)

// Client is an RPC client for the host.
type Client struct {
	rpcClient *rpc.Client
	broker    *goplugin.MuxBroker
}

// ClientOpts is an option for initializing a Client.
type ClientOpts struct {
	Cmd *exec.Cmd
}

// NewClient is a wrapper of plugin.NewClient.
func NewClient(opts *ClientOpts) *goplugin.Client {
	return goplugin.NewClient(&goplugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins: map[string]goplugin.Plugin{
			"formatter": &formatterPlugin{},
		},
		Cmd: opts.Cmd,
		Logger: hclog.New(&hclog.LoggerOptions{
			Name:   "plugin",
			Output: os.Stderr,
			Level:  hclog.LevelFromString(os.Getenv("TFDOCS_LOG")),
		}),
	})
}

// Name calls the server-side Name method and returns its version.
func (c *Client) Name() (string, error) {
	var resp string
	err := c.rpcClient.Call("Plugin.Name", new(interface{}), &resp)
	return resp, err
}

// Version calls the server-side Version method and returns its version.
func (c *Client) Version() (string, error) {
	var resp string
	err := c.rpcClient.Call("Plugin.Version", new(interface{}), &resp)
	return resp, err
}

// Execute calls the server-side Execute method and returns generated output.
func (c *Client) Execute(args ExecuteArgs) (string, error) {
	var resp string
	err := c.rpcClient.Call("Plugin.Execute", args, &resp)
	return resp, err
}
