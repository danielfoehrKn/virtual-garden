// Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package virtualgarden

import (
	"context"
)

// DeployKubeAPIServer deploys a kubernetes api server.
func (o *operation) DeployKubeAPIServer(ctx context.Context) error {
	return nil
}

// DeleteKubeAPIServer deletes the kube-apiserver and all related resources.
func (o *operation) DeleteKubeAPIServer(ctx context.Context) error {
	return nil
}

func kubeAPIServerLabels() map[string]string {
	return map[string]string{
		LabelKeyApp:       Prefix,
		LabelKeyComponent: "kube-apiserver",
	}
}