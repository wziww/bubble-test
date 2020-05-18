package config

import (
	"log"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
)

// Config bubble-test config
type Config struct {
	init   bool
	Docker Docker `toml:"docker"`
	Server Server `toml:"server"`
}

// Server bubble-test config
type Server struct {
	Host string `toml:"host"`
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
	if c.Server.Host == "" {
		c.Server.Host = "0.0.0.0:8080"
	}
}

func getDefaultConfig() *Config {
	if cfg.init {
		return &cfg
	}
	c := &Config{
		Docker: Docker{},
	}
	setDefaultConfig(c)
	return c
}

// Get ...
func Get() *Config {
	return &cfg
}

// LoadConfig ...
func LoadConfig(file string) (*Config, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if cfg.init {
		return &cfg, nil
	}
	cfg = Config{}
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
			return nil, err
		}
	}
	if _, err = toml.DecodeFile(file, &cfg); err != nil {
		log.Fatal(err)
		return nil, err
	}
	setDefaultConfig(&cfg)
	cfg.init = true
	return &cfg, nil
}
