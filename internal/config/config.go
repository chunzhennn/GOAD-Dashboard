package config

import (
	"fmt"
	"os"
)

// Config holds all configuration values for the application
type Config struct {
	proxmoxURL       string
	proxmoxAuthToken string
	pfsenseURL       string
	pfsenseUsername  string
	pfsensePassword  string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	config := &Config{}

	config.proxmoxURL = os.Getenv("PROXMOX_URL")
	if config.proxmoxURL == "" {
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

	config.proxmoxAuthToken = fmt.Sprintf("%s@%s!%s=%s", pveUsername, pveRealm, pveAPITokenName, pveAPIToken)

	config.pfsenseURL = os.Getenv("PFSENSE_URL")
	if config.pfsenseURL == "" {
		return nil, fmt.Errorf("PFSENSE_URL environment variable is required")
	}

	config.pfsenseUsername = os.Getenv("PFSENSE_USERNAME")
	if config.pfsenseUsername == "" {
		return nil, fmt.Errorf("PFSENSE_USERNAME environment variable is required")
	}

	config.pfsensePassword = os.Getenv("PFSENSE_PASSWORD")
	if config.pfsensePassword == "" {
		return nil, fmt.Errorf("PFSENSE_PASSWORD environment variable is required")
	}

	return config, nil
}

// GetProxmoxURL returns the Proxmox URL
func (c *Config) GetProxmoxURL() string {
	return c.proxmoxURL
}

// GetProxmoxAuthToken returns the Proxmox auth token
func (c *Config) GetProxmoxAuthToken() string {
	return c.proxmoxAuthToken
}

// GetPfsenseURL returns the Pfsense URL
func (c *Config) GetPfsenseURL() string {
	return c.pfsenseURL
}

// GetPfsenseUsername returns the Pfsense username
func (c *Config) GetPfsenseUsername() string {
	return c.pfsenseUsername
}

// GetPfsensePassword returns the Pfsense password
func (c *Config) GetPfsensePassword() string {
	return c.pfsensePassword
}
