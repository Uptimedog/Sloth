// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/clivern/sloth/core/module"

	"github.com/drone/envsubst"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start Sloth Worker",
	Run: func(cmd *cobra.Command, args []string) {
		configUnparsed, err := ioutil.ReadFile(config)

		if err != nil {
			panic(fmt.Sprintf(
				"Error while reading config file [%s]: %s",
				config,
				err.Error(),
			))
		}

		configParsed, err := envsubst.EvalEnv(string(configUnparsed))

		if err != nil {
			panic(fmt.Sprintf(
				"Error while parsing config file [%s]: %s",
				config,
				err.Error(),
			))
		}

		viper.SetConfigType("yaml")
		err = viper.ReadConfig(bytes.NewBuffer([]byte(configParsed)))

		if err != nil {
			panic(fmt.Sprintf(
				"Error while loading configs [%s]: %s",
				config,
				err.Error(),
			))
		}

		if viper.GetString("log.output") != "stdout" {
			fs := module.FileSystem{}
			dir, _ := filepath.Split(viper.GetString("log.output"))

			if !fs.DirExists(dir) {
				if _, err := fs.EnsureDir(dir, 777); err != nil {
					panic(fmt.Sprintf(
						"Directory [%s] creation failed with error: %s",
						dir,
						err.Error(),
					))
				}
			}

			if !fs.FileExists(viper.GetString("log.output")) {
				f, err := os.Create(viper.GetString("log.output"))
				if err != nil {
					panic(fmt.Sprintf(
						"Error while creating log file [%s]: %s",
						viper.GetString("log.output"),
						err.Error(),
					))
				}
				defer f.Close()
			}
		}

		if viper.GetString("log.output") == "stdout" {
			log.SetOutput(os.Stdout)
		} else {
			f, _ := os.Create(viper.GetString("log.output"))
			log.SetOutput(f)
		}

		lvl := strings.ToLower(viper.GetString("log.level"))
		level, err := log.ParseLevel(lvl)

		if err != nil {
			level = log.InfoLevel
		}

		log.SetLevel(level)

		if viper.GetString("log.format") == "json" {
			log.SetFormatter(&log.JSONFormatter{})
		} else {
			log.SetFormatter(&log.TextFormatter{})
		}

		fmt.Println("Worker Started ...")
	},
}

func init() {
	workerCmd.Flags().StringVarP(&config, "config", "c", "config.prod.yml", "Absolute path to config file (required)")
	workerCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(workerCmd)
}
