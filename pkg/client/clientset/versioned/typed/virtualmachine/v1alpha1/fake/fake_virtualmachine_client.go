/*
Copyright 2018 Rancher Labs, Inc.

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
package fake

import (
	v1alpha1 "github.com/rancher/vm/pkg/client/clientset/versioned/typed/virtualmachine/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeVirtualmachineV1alpha1 struct {
	*testing.Fake
}

func (c *FakeVirtualmachineV1alpha1) ARPTables() v1alpha1.ARPTableInterface {
	return &FakeARPTables{c}
}

func (c *FakeVirtualmachineV1alpha1) Credentials() v1alpha1.CredentialInterface {
	return &FakeCredentials{c}
}

func (c *FakeVirtualmachineV1alpha1) Settings() v1alpha1.SettingInterface {
	return &FakeSettings{c}
}

func (c *FakeVirtualmachineV1alpha1) VirtualMachines() v1alpha1.VirtualMachineInterface {
	return &FakeVirtualMachines{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeVirtualmachineV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
