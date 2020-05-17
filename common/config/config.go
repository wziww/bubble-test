package config

import (
	"log"
	"os"
	"sync"
)

// Config bubble-test config
type Config struct {
	init   bool
	Docker Docker `toml:"docker"`
}

// Docker config set to connect to docker
type Docker struct {
	Host string `toml:"host"`
	//  defaultAPIVersion = "1.39"
	APIVersion string `toml:"apiVersion"`
}

var (
	cfg   Config
	mutex sync.Mutex
)

func setDefaultConfig(c *Config) {
	if c.Docker.Host == "" {
		c.Docker.Host = DefaultDockerHost
	}
	if c.Docker.APIVersion == "" {
		c.Docker.APIVersion = DefaultAPIVersion
	}
}

func getDefaultConfig() *Config {
	mutex.Lock()
	defer mutex.Unlock()
	if cfg.init {
		return &cfg
	}
	c := &Config{
		Docker: Docker{},
	}
	setDefaultConfig(c)
	return c
}

// LoadConfig ...
func LoadConfig(file string) *Config {
	mutex.Lock()
	defer mutex.Unlock()
	if cfg.init {
		return &cfg
	}
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
			return nil
		}
	}
	return nil
}
