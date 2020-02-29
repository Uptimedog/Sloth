// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/clivern/sloth/internal/app/worker"
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
	fmt.Println("Worker started .....")
}
