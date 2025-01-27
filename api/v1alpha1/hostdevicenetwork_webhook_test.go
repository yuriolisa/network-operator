/*
2023 NVIDIA CORPORATION & AFFILIATES

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
package v1alpha1 //nolint:dupl

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//nolint:dupl
var _ = Describe("Validate", func() {
	Context("HostDeviceNetwork tests", func() {
		It("Valid ResourceName", func() {
			hostDeviceNetwork := HostDeviceNetwork{
				ObjectMeta: metav1.ObjectMeta{Name: "test"},
				Spec: HostDeviceNetworkSpec{
					ResourceName: "hostdev",
				},
			}
			_, err := hostDeviceNetwork.ValidateCreate()
			Expect(err).NotTo(HaveOccurred())
		})
		It("Invalid ResourceName", func() {
			hostDeviceNetwork := HostDeviceNetwork{
				ObjectMeta: metav1.ObjectMeta{Name: "test"},
				Spec: HostDeviceNetworkSpec{
					ResourceName: "hostdev!!",
				},
			}
			_, err := hostDeviceNetwork.ValidateCreate()
			Expect(err.Error()).To(ContainSubstring("Invalid Resource name"))
		})
	})
})
