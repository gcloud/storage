// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package storage

import (
	"github.com/gcloud/identity"
)

// The containers available from the storage service.
type Containers struct {
	identity.Account
}

// List containers available to the account.
func (s *Containers) List() {}

// Show container information for a given id.
func (s *Containers) Show(id string) {}

// Create a container.
func (s *Containers) Create() {}

// Destroy a container.
func (s *Containers) Destroy() {}

// Distribute containers to mutliple regions.
func (s *Containers) Distribute() {}
