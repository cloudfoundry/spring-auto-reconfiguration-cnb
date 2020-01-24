/*
 * Copyright 2018-2020 the original author or authors.
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

package autoreconfiguration_test

import (
	"path/filepath"
	"testing"

	"github.com/cloudfoundry/libcfbuildpack/v2/buildpackplan"
	"github.com/cloudfoundry/libcfbuildpack/v2/test"
	"github.com/cloudfoundry/spring-auto-reconfiguration-cnb/autoreconfiguration"
	"github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestAutoReconfiguration(t *testing.T) {
	spec.Run(t, "AutoReconfiguration", func(t *testing.T, _ spec.G, it spec.S) {

		g := gomega.NewWithT(t)

		var f *test.BuildFactory

		it.Before(func() {
			f = test.NewBuildFactory(t)
		})

		it("returns false if build plan does not exist", func() {
			test.TouchFile(t, filepath.Join(f.Build.Application.Root, "spring-core-1.2.3.RELEASE.jar"))

			_, ok, err := autoreconfiguration.NewAutoReconfiguration(f.Build)
			g.Expect(err).NotTo(gomega.HaveOccurred())
			g.Expect(ok).To(gomega.BeFalse())
		})

		it("returns false if spring core jar file does not exist", func() {
			f.AddPlan(buildpackplan.Plan{Name: autoreconfiguration.Dependency})

			_, ok, err := autoreconfiguration.NewAutoReconfiguration(f.Build)
			g.Expect(err).NotTo(gomega.HaveOccurred())
			g.Expect(ok).To(gomega.BeFalse())
		})

		it("returns true if build plan and spring core jar file both exist", func() {
			f.AddPlan(buildpackplan.Plan{Name: autoreconfiguration.Dependency})
			f.AddDependency(autoreconfiguration.Dependency, filepath.Join("testdata", "stub-auto-reconfiguration.jar"))
			test.TouchFile(t, filepath.Join(f.Build.Application.Root, "spring-core-1.2.3.RELEASE.jar"))

			_, ok, err := autoreconfiguration.NewAutoReconfiguration(f.Build)
			g.Expect(err).NotTo(gomega.HaveOccurred())
			g.Expect(ok).To(gomega.BeTrue())
		})

		it("contributes jar", func() {
			f.AddPlan(buildpackplan.Plan{Name: autoreconfiguration.Dependency})
			f.AddDependency(autoreconfiguration.Dependency, filepath.Join("testdata", "stub-auto-reconfiguration.jar"))
			test.TouchFile(t, filepath.Join(f.Build.Application.Root, "test", "spring-core-1.2.3.RELEASE.jar"))

			y, ok, err := autoreconfiguration.NewAutoReconfiguration(f.Build)
			g.Expect(err).NotTo(gomega.HaveOccurred())
			g.Expect(ok).To(gomega.BeTrue())

			g.Expect(y.Contribute()).To(gomega.Succeed())

			layer := f.Build.Layers.Layer("auto-reconfiguration")
			g.Expect(layer).To(test.HaveLayerMetadata(false, false, true))
			g.Expect(filepath.Join(layer.Root, "stub-auto-reconfiguration.jar")).To(gomega.BeARegularFile())
			g.Expect(layer).To(test.HavePrependPathLaunchEnvironment("CLASSPATH", "%s",
				filepath.Join(layer.Root, "stub-auto-reconfiguration.jar")))
		})
	}, spec.Report(report.Terminal{}))
}
