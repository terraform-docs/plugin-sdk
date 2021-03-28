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

	// IndentLevel control the indentation of headers [available: 1, 2, 3, 4, 5]
	//
	// default: 2
	// scope: Asciidoc, Markdown
	IndentLevel int

	// OutputValues extract and show Output values from Terraform module output
	//
	// default: false
	// scope: Global
	OutputValues bool

	// ShowColor print "colorized" version of result in the terminal
	//
	// default: true
	// scope: Pretty
	ShowColor bool

	// ShowDataSources shows the data sources in the "Resources" section
	//
	// default: true
	// scope: Global
	ShowDataSources bool

	// ShowDefault show "Default" column
	//
	// default: true
	// scope: Asciidoc, Markdown
	ShowDefault bool

	// ShowFooter show "footer" module information
	//
	// default: false
	// scope: Global
	ShowFooter bool

	// ShowHeader show "Header" module information
	//
	// default: true
	// scope: Global
	ShowHeader bool

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

	// ShowRequired show "Required" column
	//
	// default: true
	// scope: Asciidoc, Markdown
	ShowRequired bool

	// ShowSensitivity show "Sensitive" column
	//
	// default: true
	// scope: Asciidoc, Markdown
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

	// ShowType show "Type" column
	//
	// default: true
	// scope: Asciidoc, Markdown
	ShowType bool
}
