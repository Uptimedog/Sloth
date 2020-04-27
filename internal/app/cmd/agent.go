// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Start Sloth Agent",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Agent Started ...")
	},
}

func init() {
	agentCmd.Flags().StringVarP(&config, "config", "c", "config.prod.yml", "Absolute path to config file (required)")
	agentCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(agentCmd)
}
