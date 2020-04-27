// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package agent

// Log struct
type Log struct {
	Level  string `mapstructure:"level"`
	Output string `mapstructure:"output"`
	Format string `mapstructure:"format"`
}

// CheckConfig struct
type CheckConfig struct {
	Name  string `mapstructure:"name"`
	Value string `mapstructure:"value"`
}

// Check struct
type Check struct {
	Type    string        `mapstructure:"type"`
	Configs []CheckConfig `mapstructure:"configs"`
}

// Agent struct
type Agent struct {
	Hostname      string  `mapstructure:"hostname"`
	ValidationKey string  `mapstructure:"validation_key"`
	ClientKey     string  `mapstructure:"client_key"`
	Checks        []Check `mapstructure:"checks"`
}

// Config struct
type Config struct {
	Role  string `mapstructure:"role"`
	Log   *Log   `mapstructure:"log"`
	Agent *Agent `mapstructure:"agent"`
}

// NewConfig create a new instance
func NewConfig() *Config {
	return &Config{}
}
