package config

import (
	"os"
	"path/filepath"
	"runtime"
)

const ConfigDir = "nbspify"

func GetConfigHome() string {
	configHome := ""

	// TODO - other operating systems
	if runtime.GOOS == "linux" {
		xdgConfigDir := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfigDir != "" {
			configHome = xdgConfigDir
		}
	}

	if configHome == "" {
		configHome = filepath.Join(os.Getenv("HOME"), ".config")
	}

	return configHome
}

func GetConfigDir() string {
	return filepath.Join(GetConfigHome(), ConfigDir)
}
