// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAgent controller
func CreateAgent(c *gin.Context) {
	//rawBody, _ := c.GetRawData()
	//body := string(rawBody)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetAgents controller
func GetAgents(c *gin.Context) {
	//rawBody, _ := c.GetRawData()
	//body := string(rawBody)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetAgent controller
func GetAgent(c *gin.Context) {
	//rawBody, _ := c.GetRawData()
	//body := string(rawBody)
	//ID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// UpdateAgent controller
func UpdateAgent(c *gin.Context) {
	//rawBody, _ := c.GetRawData()
	//body := string(rawBody)
	//ID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// DeleteAgent controller
func DeleteAgent(c *gin.Context) {
	//rawBody, _ := c.GetRawData()
	//body := string(rawBody)
	//ID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
