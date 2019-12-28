// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/silverbackhq/sloth/internal/app/agent"

	"github.com/spf13/viper"
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
	fmt.Println("Agent started .....")
}

// Register registers the agent
func (a *Agent) Register() error {

	err := viper.Unmarshal(a.Config)

	if err != nil {
		return fmt.Errorf(
			"configs are invalid [%s]",
			err.Error(),
		)
	}

	return nil
}
