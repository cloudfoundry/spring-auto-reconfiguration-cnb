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

package main

import (
	"path/filepath"
	"testing"

	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/buildpack/libbuildpack/detect"
	"github.com/cloudfoundry/libcfbuildpack/test"
	"github.com/cloudfoundry/spring-auto-reconfiguration-cnb/autoreconfiguration"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestDetect(t *testing.T) {
	spec.Run(t, "Detect", func(t *testing.T, _ spec.G, it spec.S) {

		g := NewGomegaWithT(t)

		var f *test.DetectFactory

		it.Before(func() {
			f = test.NewDetectFactory(t)
		})

		it("passes with a spring core jar present", func() {
			test.TouchFile(t, filepath.Join(f.Detect.Application.Root, "spring-core-1.2.3.RELEASE.jar"))

			g.Expect(d(f.Detect)).To(Equal(detect.PassStatusCode))
			g.Expect(f.Output).To(Equal(buildplan.BuildPlan{
				autoreconfiguration.Dependency: buildplan.Dependency{},
			}))
		})

		it("fails with no spring core jar present", func() {
			g.Expect(d(f.Detect)).To(Equal(detect.FailStatusCode))
		})
	}, spec.Report(report.Terminal{}))
}
