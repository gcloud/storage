// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package storage

import (
	"github.com/gcloud/identity"
)

// The objects available from the storage service.
type Objects struct {
	identity.Account
	Containers
}

// List objects available to the account.
func (s *Objects) List() {}

// Show object information for a given id.
func (s *Objects) Show(id string) {}

// Create an object.
func (s *Objects) Create() {}

// Destroy an object.
func (s *Objects) Destroy() {}
