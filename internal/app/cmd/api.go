// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/silverbackhq/sloth/internal/app/api"
)

// API struct
type API struct {
	Config *api.Config
}

// NewAPI create a new instance
func NewAPI(config *api.Config) *API {
	return &API{
		Config: config,
	}
}

// Run runs the api
func (w *API) Run() {
	fmt.Println("API server started .....")
}
