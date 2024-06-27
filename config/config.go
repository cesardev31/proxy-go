package config

import (
	"os"

	"github.com/pelletier/go-toml"
)

type GlobalConfig struct {
	CheckNewVersion    bool `toml:"checkNewVersion"`
	SendAnonymousUsage bool `toml:"sendAnonymousUsage"`
}

type EntryPoint struct {
	Address string `toml:"address"`
}

type EntryPointsConfig struct {
	Web       EntryPoint `toml:"web"`
	WebSecure EntryPoint `toml:"websecure"`
}

type LogConfig struct {
	Level    string `toml:"level"`
	FilePath string `toml:"filePath"`
	Format   string `toml:"format"`
}

type APIConfig struct {
	Insecure  bool `toml:"insecure"`
	Dashboard bool `toml:"dashboard"`
}

type PingConfig struct {
	EntryPoint string `toml:"entryPoint"`
}

type DockerConfig struct {
	Endpoint         string `toml:"endpoint"`
	DefaultRule      string `toml:"defaultRule"`
	ExposedByDefault bool   `toml:"exposedByDefault"`
}

type ProvidersConfig struct {
	Docker DockerConfig `toml:"docker"`
}

type TCPBackendConfig struct {
	Name string `toml:"name"`
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type BackendConfig struct {
	Name string `toml:"name"`
	URL  string `toml:"url"`
}
type AccessLogConfig struct {
	FilePath string `toml:"filePath"`
}

type Config struct {
	Global      GlobalConfig       `toml:"global"`
	EntryPoints EntryPointsConfig  `toml:"entryPoints"`
	Log         LogConfig          `toml:"log"`
	API         APIConfig          `toml:"api"`
	Ping        PingConfig         `toml:"ping"`
	Providers   ProvidersConfig    `toml:"providers"`
	Backends    []BackendConfig    `toml:"backends"`
	AccessLog   AccessLogConfig    `toml:"accessLog"`
	TCPBackends []TCPBackendConfig `toml:"tcpBackends"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = toml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
