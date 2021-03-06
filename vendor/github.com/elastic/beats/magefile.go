// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// +build mage

package main

import (
	"path/filepath"

	"github.com/elastic/beats/dev-tools/mage"
)

var (
	// Beats is a list of Beats to collect dashboards from.
	Beats = []string{
		"auditbeat",
		"filebeat",
		"heartbeat",
		"journalbeat",
		"metricbeat",
		"packetbeat",
		"winlogbeat",
		"x-pack/functionbeat",
	}
)

// PackageBeatDashboards packages the dashboards from all Beats into a zip
// file. The dashboards must be generated first.
func PackageBeatDashboards() error {
	version, err := mage.BeatVersion()
	if err != nil {
		return err
	}

	spec := mage.PackageSpec{
		Name:     "beats-dashboards",
		Version:  version,
		Snapshot: mage.Snapshot,
		Files: map[string]mage.PackageFile{
			".build_hash.txt": mage.PackageFile{
				Content: "{{ commit }}\n",
			},
		},
		OutputFile: "build/distributions/dashboards/{{.Name}}-{{.Version}}{{if .Snapshot}}-SNAPSHOT{{end}}",
	}

	for _, beat := range Beats {
		spec.Files[beat] = mage.PackageFile{
			Source: filepath.Join(beat, "_meta/kibana.generated"),
		}
	}

	return mage.PackageZip(spec.Evaluate())
}

// DumpVariables writes the template variables and values to stdout.
func DumpVariables() error {
	return mage.DumpVariables()
}
