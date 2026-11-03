// Package ee is the entry point for the Enterprise edition add-on.
//
// The Enterprise backend is not implemented yet. This package defines the
// contract and build-tag wiring so future paid modules can be added without
// changing the public core or the SaaS add-on.
package ee

import (
	"google.golang.org/grpc"

	"github.com/werbot/werbot/api"
	"github.com/werbot/werbot/internal/addon"
)

// Addon implements the addon.Addon contract for the Enterprise edition.
type Addon struct{}

// New creates the Enterprise add-on.
func New() *Addon {
	return &Addon{}
}

// Name returns the add-on identifier.
func (a *Addon) Name() string {
	return "ee"
}

// Routes mounts the Enterprise HTTP routes. No routes exist yet.
func (a *Addon) Routes(_ *api.Handler) {}

// RegisterGRPC registers the Enterprise gRPC services. No services exist yet.
func (a *Addon) RegisterGRPC(_ grpc.ServiceRegistrar, _ addon.Deps) {}
