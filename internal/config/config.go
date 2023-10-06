// Package config agent's setting parser. Applies flags and environments. Environments are prioritized.
package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	ServerAddr       string
	ConnectionsCount int
	GeneratorPath    string
}

// Parse main func to parse variables.
func Parse() Config {
	var config = Config{}
	checkFlags(&config)
	checkEnvironments(&config)
	return config
}

// FLAGS PARSING.
const (
	flagServerAddress    = "a"
	flagConnectionsCount = "n"
	flagGeneratorPath    = "p"
)

// checkFlags checks flags of app's launch.
func checkFlags(config *Config) {
	flag.StringVar(&config.ServerAddr, flagServerAddress, "localhost:8080", "server's address")
	flag.IntVar(&config.ConnectionsCount, flagConnectionsCount, 4, "parallel connection count")
	flag.StringVar(&config.GeneratorPath, flagGeneratorPath, "C:\\Users\\", "path for tasks generator")
	flag.Parse()
}

// ENVIRONMENTS PARSING.
// envConfig struct of environments suitable for agent.
type envConfig struct {
	ServerAddr       string `env:"HOST"`
	ConnectionsCount string `env:"CONN_COUNT"`
	GeneratorPath    string `env:"GENERATOR_PATH"`
}

// checkEnvironments checks environments suitable for agent.
func checkEnvironments(config *Config) {
	var envs = envConfig{}
	err := env.Parse(&envs)
	if err != nil {
		log.Fatal(err)
	}

	_ = setEnvToParamIfNeed(&config.ServerAddr, envs.ServerAddr)
	_ = setEnvToParamIfNeed(&config.ConnectionsCount, envs.ConnectionsCount)
	_ = setEnvToParamIfNeed(&config.GeneratorPath, envs.GeneratorPath)
}
