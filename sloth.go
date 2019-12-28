// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/silverbackhq/sloth/internal/app/module"
	"github.com/silverbackhq/sloth/internal/app/util"

	"github.com/drone/envsubst"
	"github.com/spf13/viper"
)

const (
	// AgentRole var
	AgentRole = "agent"
	// WorkerRole var
	WorkerRole = "worker"
	// APIRole var
	APIRole = "api"
)

func main() {
	var configFile string
	var role string

	flag.StringVar(&configFile, "config", "config.prod.yml", "config")
	flag.StringVar(&role, "role", "", "role")
	flag.Parse()

	configUnparsed, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(fmt.Sprintf(
			"Error while reading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	configParsed, err := envsubst.EvalEnv(string(configUnparsed))

	if err != nil {
		panic(fmt.Sprintf(
			"Error while parsing config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer([]byte(configParsed)))

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading configs [%s]: %s",
			configFile,
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

	if role == "" {
		role = strings.ToLower(viper.GetString("role"))
	}

	if !util.InArray(role, []string{AgentRole, WorkerRole, APIRole}) {
		panic(fmt.Sprintf(
			"Error! Invalid role [%s]",
			role,
		))
	}

	if role == strings.ToLower(AgentRole) {
		agent := InitializeNewAgent()
		err := agent.Register()
		if err != nil {
			panic(fmt.Sprintf("Error! Unable to register agent: %s", err.Error()))
		}
		agent.Run()
	}

	if role == strings.ToLower(WorkerRole) {
		InitializeNewWorker().Run()
	}

	if role == strings.ToLower(APIRole) {
		InitializeNewAPI().Run()
	}
}
