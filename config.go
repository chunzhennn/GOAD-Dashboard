package main

import (
	"fmt"
	"os"
)

// Config holds all configuration values for the application
type Config struct {
	ProxmoxURL       string
	ProxmoxAuthToken string
	PfsenseURL       string
	PfsenseUsername  string
	PfsensePassword  string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	config := &Config{}

	config.ProxmoxURL = os.Getenv("PROXMOX_URL")
	if config.ProxmoxURL == "" {
		return nil, fmt.Errorf("PROXMOX_URL environment variable is required")
	}

	pveUsername := os.Getenv("PROXMOX_USERNAME")
	if pveUsername == "" {
		return nil, fmt.Errorf("PROXMOX_USERNAME environment variable is required")
	}

	pveRealm := os.Getenv("PROXMOX_REALM")
	if pveRealm == "" {
		return nil, fmt.Errorf("PROXMOX_REALM environment variable is required")
	}

	pveAPITokenName := os.Getenv("PROXMOX_API_TOKEN_NAME")
	if pveAPITokenName == "" {
		return nil, fmt.Errorf("PROXMOX_API_TOKEN_NAME environment variable is required")
	}

	pveAPIToken := os.Getenv("PROXMOX_API_TOKEN")
	if pveAPIToken == "" {
		return nil, fmt.Errorf("PROXMOX_API_TOKEN environment variable is required")
	}

	config.ProxmoxAuthToken = fmt.Sprintf("%s@%s!%s=%s", pveUsername, pveRealm, pveAPITokenName, pveAPIToken)

	return config, nil
}

func (c *Config) GetProxmoxURL() string {
	return c.ProxmoxURL
}

func (c *Config) GetProxmoxAuthToken() string {
	return c.ProxmoxAuthToken
}
