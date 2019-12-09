// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/silverbackhq/sloth/internal/app/agent"
)

// Agent struct
type Agent struct {
	Config *agent.Config
}

// NewAgent create a new instance
func NewAgent(config *agent.Config) *Agent {
	return &Agent{
		Config: config,
	}
}

// Run runs the agent
func (a *Agent) Run() {
}
