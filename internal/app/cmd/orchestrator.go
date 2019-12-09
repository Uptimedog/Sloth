// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/silverbackhq/sloth/internal/app/orchestrator"
)

// Orchestrator struct
type Orchestrator struct {
	Config *orchestrator.Config
}

// NewOrchestrator create a new instance
func NewOrchestrator(config *orchestrator.Config) *Orchestrator {
	return &Orchestrator{
		Config: config,
	}
}

// Run runs the orchestrator
func (o *Orchestrator) Run() {
}
