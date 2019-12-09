// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/silverbackhq/sloth/internal/app/agent"
	"github.com/silverbackhq/sloth/internal/app/cmd"
	"github.com/silverbackhq/sloth/internal/app/orchestrator"
	"github.com/silverbackhq/sloth/internal/app/worker"
)

func InitializeNewAgent() *cmd.Agent {
	wire.Build(cmd.NewAgent, agent.NewConfig)
	return &cmd.Agent{}
}

func InitializeNewOrchestrator() *cmd.Orchestrator {
	wire.Build(cmd.NewOrchestrator, orchestrator.NewConfig)
	return &cmd.Orchestrator{}
}

func InitializeNewWorker() *cmd.Worker {
	wire.Build(cmd.NewWorker, worker.NewConfig)
	return &cmd.Worker{}
}
