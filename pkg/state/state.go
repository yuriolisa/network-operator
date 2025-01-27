/*
Copyright 2020 NVIDIA

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

package state

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SyncState string

// Represents the Sync state of a specific State or a collection of States
const (
	SyncStateReady    = "ready"
	SyncStateNotReady = "notReady"
	SyncStateIgnore   = "ignore"
	SyncStateReset    = "reset"
	SyncStateError    = "error"
)

// State Represents a single State that requires a set of k8s API operations to be performed.
// A state is associated with a set of resources, it checks the system state against the given set of resources
// and reconciles accordingly. It basically reconciles the system to the given state.
type State interface {
	// Name provides the State name
	Name() string
	// Description provides the State description
	Description() string
	// Sync attempt to get the system to match the desired state as depicted in the custom resource
	// for the bits related to the specific state, State represents.
	// a sync operation must be relatively short and must not block the execution thread.
	// InfoCatalog is provided to optionally provide a State additional infoSources required for it to perform
	// the Sync operation.
	Sync(ctx context.Context, customResource interface{}, infoCatalog InfoCatalog) (SyncState, error)
	// Get a map of source kinds that should be watched for the state keyed by the source kind name
	GetWatchSources() map[string]client.Object
}
