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

	goplugin "github.com/hashicorp/go-plugin"
)

// handShakeConfig is used for UX. ProcotolVersion will be updated by incompatible changes.
var handshakeConfig = goplugin.HandshakeConfig{
	ProtocolVersion:  7,
	MagicCookieKey:   "TFDOCS_PLUGIN",
	MagicCookieValue: "A7U5oTDDJwdL6UKOw6RXATDa86NEo4xLK3rz7QqegT1N4EY66qb6UeAJDSxLwtXH",
}

// formatterPlugin is a wrapper to satisfy the interface of go-plugin.
type formatterPlugin struct {
	impl plugin
}

// Server returns an RPC server acting as a plugin.
func (p *formatterPlugin) Server(b *goplugin.MuxBroker) (interface{}, error) {
	return &server{impl: p.impl, broker: b}, nil
}

// Client returns an RPC client for the host.
func (formatterPlugin) Client(b *goplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &Client{rpcClient: c, broker: b}, nil
}
