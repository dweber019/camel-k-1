// +build integration

// To enable compilation of this file in Goland, go to "Settings -> Go -> Vendoring & Build Tags -> Custom Tags" and add "integration"

/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"github.com/apache/camel-k/pkg/util/defaults"
	"testing"

	. "github.com/onsi/gomega"

	. "github.com/apache/camel-k/e2e/support"
)

func TestKamelCLIVersion(t *testing.T) {
	WithNewTestNamespace(t, func(ns string) {
		t.Run("check version correctness", func(t *testing.T) {
			version := GetOutputString(Kamel("version"))
			Expect(version).To(ContainSubstring(defaults.Version))
		})
	})
}
