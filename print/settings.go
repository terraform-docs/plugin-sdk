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

package print

// Settings represents all settings.
type Settings struct {
	// EscapeCharacters escapes special characters (such as _ * in Markdown and > < in JSON)
	//
	// default: true
	// scope: Markdown
	EscapeCharacters bool

	// EscapePipe escapes pipe character in Markdown
	//
	// default: true
	// scope: Markdown
	EscapePipe bool

	// IndentLevel control the indentation of AsciiDoc and Markdown headers [available: 1, 2, 3, 4, 5]
	//
	// default: 2
	// scope: Asciidoc, Markdown
	IndentLevel int

	// OutputValues ailrghaekrgj
	//
	// default: false
	// scope: Global
	OutputValues bool

	// ShowColor print "colorized" version of result in the terminal
	//
	// default: true
	// scope: Pretty
	ShowColor bool

	// ShowHeader show "Header" module information
	//
	// default: true
	// scope: Global
	ShowHeader bool

	// ShowFooter show "footer" module information
	//
	// default: false
	// scope: Global
	ShowFooter bool

	// ShowInputs show "Inputs" information
	//
	// default: true
	// scope: Global
	ShowInputs bool

	// ShowModuleCalls show "ModuleCalls" information
	//
	// default: true
	// scope: Global
	ShowModuleCalls bool

	// ShowOutputs show "Outputs" information
	//
	// default: true
	// scope: Global
	ShowOutputs bool

	// ShowProviders show "Providers" information
	//
	// default: true
	// scope: Global
	ShowProviders bool

	// ShowRequired show "Required" column when generating Markdown
	//
	// default: true
	// scope: Markdown
	ShowRequired bool

	// ShowSensitivity show "Sensitive" column when generating Markdown
	//
	// default: true
	// scope: Markdown
	ShowSensitivity bool

	// ShowRequirements show "Requirements" section
	//
	// default: true
	// scope: Global
	ShowRequirements bool

	// ShowResources show "Resources" section
	//
	// default: true
	// scope: Global
	ShowResources bool

	// SortByName sorted rendering of inputs and outputs
	//
	// default: true
	// scope: Global
	SortByName bool

	// SortByRequired sort items (inputs, providers) by name and prints required ones first
	//
	// default: false
	// scope: Global
	SortByRequired bool

	// SortByType sort items (inputs, outputs) by type alphabetically
	//
	// default: false
	// scope: Global
	SortByType bool
}
