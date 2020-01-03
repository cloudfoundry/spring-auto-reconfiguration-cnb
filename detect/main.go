/*
 * Copyright 2019-2020 the original author or authors.
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
	"fmt"
	"os"
	"strconv"

	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/cloudfoundry/jvm-application-cnb/jvmapplication"
	"github.com/cloudfoundry/libcfbuildpack/detect"
	"github.com/cloudfoundry/spring-auto-reconfiguration-cnb/autoreconfiguration"
)

func main() {
	detect, err := detect.DefaultDetect()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialize Detect: %s\n", err)
		os.Exit(101)
	}

	if code, err := d(detect); err != nil {
		detect.Logger.TerminalError(detect.Buildpack, err.Error())
		os.Exit(code)
	} else {
		os.Exit(code)
	}
}

func d(detect detect.Detect) (int, error) {
	e := true

	if v, ok := os.LookupEnv("BP_AUTO_RECONFIGURATION_ENABLED"); ok {
		if f, err := strconv.ParseBool(v); err != nil {
			return detect.Error(102), err
		} else {
			e = f
		}
	}

	if !e {
		return detect.Fail(), nil
	}

	return detect.Pass(buildplan.Plan{
		Provides: []buildplan.Provided{
			{Name: autoreconfiguration.Dependency},
		},
		Requires: []buildplan.Required{
			{Name: autoreconfiguration.Dependency},
			{Name: jvmapplication.Dependency},
		},
	})
}
