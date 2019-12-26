// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/silverbackhq/sloth/internal/app/api"
	"github.com/silverbackhq/sloth/internal/app/api/controller"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

	r := gin.Default()

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusNoContent, "")
	})

	r.GET("/api/_health", controller.HealthCheck)
	r.POST("/api/agents", controller.CreateAgent)
	r.GET("/api/agents", controller.GetAgents)
	r.GET("/api/agents/:id", controller.GetAgent)
	r.PUT("/api/agents/:id", controller.UpdateAgent)
	r.DELETE("/api/agents/:id", controller.DeleteAgent)

	if viper.GetBool("api.tls.status") {
		r.RunTLS(
			fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("api.port"))),
			viper.GetString("api.tls.pemPath"),
			viper.GetString("api.tls.keyPath"),
		)
	} else {
		r.Run(
			fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("api.port"))),
		)
	}
}
