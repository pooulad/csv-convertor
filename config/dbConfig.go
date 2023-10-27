package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type MysqlConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Name     string `json:"name"`
	Port     string `json:"port"`
	Table    string `json:"table_name"`
}

type PostgresConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Name     string `json:"name"`
	SSL      string `json:"ssl"`
	Table    string `json:"table_name"`
}

func ReadMysqlConfig(path string) (*MysqlConfig, error) {
	if path == "" {
		return nil, fmt.Errorf("no path provided")
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg MysqlConfig
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Host == "" {
		return nil, fmt.Errorf("no host provided(config file)")
	}
	if cfg.Port == "" {
		return nil, fmt.Errorf("no port provided(config file)")
	}
	if cfg.Name == "" {
		return nil, fmt.Errorf("no name provided(config file)")
	}
	if cfg.Password == "" {
		return nil, fmt.Errorf("no password provided(config file)")
	}
	if cfg.Username == "" {
		return nil, fmt.Errorf("no username provided(config file)")
	}
	if cfg.Table == "" {
		return nil, fmt.Errorf("no table provided(config file)")
	}

	return &cfg, nil
}

func ReadPostgresConfig(path string) (*PostgresConfig, error) {
	if path == "" {
		return nil, fmt.Errorf("no path provided")
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg PostgresConfig
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Host == "" {
		return nil, fmt.Errorf("no host provided(config file)")
	}
	if cfg.SSL == "" {
		return nil, fmt.Errorf("no ssl provided(config file)")
	}
	if cfg.Name == "" {
		return nil, fmt.Errorf("no name provided(config file)")
	}
	if cfg.Password == "" {
		return nil, fmt.Errorf("no password provided(config file)")
	}
	if cfg.Username == "" {
		return nil, fmt.Errorf("no username provided(config file)")
	}
	if cfg.Table == "" {
		return nil, fmt.Errorf("no table provided(config file)")
	}

	return &cfg, nil
}
