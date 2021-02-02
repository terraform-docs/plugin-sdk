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
	"github.com/terraform-docs/plugin-sdk/print"
)

type plugin interface {
	Name() string
	Version() string
	Execute(ExecuteArgs) (string, error)
}

type base struct {
	name    string
	version string
	engine  print.Engine
}

func (b *base) Name() string {
	return b.name
}

func (b *base) Version() string {
	return b.version
}

func (b *base) Execute(args ExecuteArgs) (string, error) {
	return b.engine.Print(args.Module, args.Settings)
}
