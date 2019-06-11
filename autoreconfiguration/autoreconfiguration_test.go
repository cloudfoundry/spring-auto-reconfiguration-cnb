/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
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

	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/cloudfoundry/libcfbuildpack/test"
	"github.com/cloudfoundry/spring-auto-reconfiguration-cnb/autoreconfiguration"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestAutoReconfiguration(t *testing.T) {
	spec.Run(t, "AutoReconfiguration", func(t *testing.T, _ spec.G, it spec.S) {

		g := NewGomegaWithT(t)

		var f *test.BuildFactory

		it.Before(func() {
			f = test.NewBuildFactory(t)
		})

		it("returns true if build plan does exist", func() {
			f.AddBuildPlan(autoreconfiguration.Dependency, buildplan.Dependency{})
			f.AddDependency(autoreconfiguration.Dependency, filepath.Join("testdata", "stub-auto-reconfiguration.jar"))

			_, ok, err := autoreconfiguration.NewAutoReconfiguration(f.Build)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(ok).To(BeTrue())
		})

		it("returns false if build plan does not exist", func() {
			_, ok, err := autoreconfiguration.NewAutoReconfiguration(f.Build)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(ok).To(BeFalse())
		})

		it("contributes jar", func() {
			f.AddBuildPlan(autoreconfiguration.Dependency, buildplan.Dependency{})
			f.AddDependency(autoreconfiguration.Dependency, filepath.Join("testdata", "stub-auto-reconfiguration.jar"))

			y, ok, err := autoreconfiguration.NewAutoReconfiguration(f.Build)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(ok).To(BeTrue())

			g.Expect(y.Contribute()).To(Succeed())

			layer := f.Build.Layers.Layer("auto-reconfiguration")
			g.Expect(layer).To(test.HaveLayerMetadata(false, false, true))
			g.Expect(filepath.Join(layer.Root, "stub-auto-reconfiguration.jar")).To(BeARegularFile())
			g.Expect(layer).To(test.HaveAppendPathLaunchEnvironment("CLASSPATH", "%s",
				filepath.Join(layer.Root, "stub-auto-reconfiguration.jar")))
		})
	}, spec.Report(report.Terminal{}))
}
