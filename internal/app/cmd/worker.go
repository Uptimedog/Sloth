// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/silverbackhq/sloth/internal/app/worker"
)

// Worker struct
type Worker struct {
	Config *worker.Config
}

// NewWorker create a new instance
func NewWorker(config *worker.Config) *Worker {
	return &Worker{
		Config: config,
	}
}

// Run runs the worker
func (w *Worker) Run() {
}
