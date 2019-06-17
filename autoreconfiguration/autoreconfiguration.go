/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package autoreconfiguration

import (
	"path/filepath"
	"regexp"

	"github.com/cloudfoundry/libcfbuildpack/build"
	"github.com/cloudfoundry/libcfbuildpack/helper"
	"github.com/cloudfoundry/libcfbuildpack/layers"
)

const Dependency = "auto-reconfiguration"

// AutoReconfiguration represents the Spring Auto-reconfiguration for a JVM application.
type AutoReconfiguration struct {
	layer layers.DependencyLayer
}

// Contribute makes the contribution to launch.
func (c AutoReconfiguration) Contribute() error {
	return c.layer.Contribute(func(artifact string, layer layers.DependencyLayer) error {
		layer.Logger.SubsequentLine("Copying to %s", layer.Root)

		destination := filepath.Join(layer.Root, layer.ArtifactName())

		if err := helper.CopyFile(artifact, destination); err != nil {
			return err
		}

		return layer.AppendPathLaunchEnv("CLASSPATH", "%s", destination)
	}, layers.Launch)
}

// NewAutoReconfiguration creates a new AutoReconfiguration instance. OK is true if
// a spring core jar is found in the application.
func NewAutoReconfiguration(build build.Build) (AutoReconfiguration, bool, error) {
	bp, ok := build.BuildPlan[Dependency]
	if !ok {
		return AutoReconfiguration{}, false, nil
	}

	if exist, err := helper.HasFile(build.Application.Root, regexp.MustCompile(filepath.Join(".*", ".*spring-core.*\\.jar"))); err != nil {
		return AutoReconfiguration{}, false, err
	} else if !exist {
		return AutoReconfiguration{}, false, nil
	}

	deps, err := build.Buildpack.Dependencies()
	if err != nil {
		return AutoReconfiguration{}, false, err
	}

	dep, err := deps.Best(Dependency, bp.Version, build.Stack)
	if err != nil {
		return AutoReconfiguration{}, false, err
	}

	return AutoReconfiguration{build.Layers.DependencyLayer(dep)}, true, nil
}
