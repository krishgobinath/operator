// Copyright (c) 2019 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// AddToManagerFuncs is a list of functions to add all Controllers to the Manager
// Each func takes the manager as well as a boolean indicating if we're running on openshift,
// as well as a boolean indicating whether we need to start TSEE controllers.
var AddToManagerFuncs []func(manager.Manager, bool, bool) error

// AddToManager adds all Controllers to the Manager
func AddToManager(m manager.Manager, openshift bool, tsee bool) error {
	for _, f := range AddToManagerFuncs {
		if err := f(m, openshift, tsee); err != nil {
			return err
		}
	}
	return nil
}
