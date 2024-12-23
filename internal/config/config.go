package config

import (
	"fmt"
	"os"
)

type Config struct {
	db_url            string
	current_user_name string
}

func Read() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Errorf("home directory could not be identified", err)
	}

	fmt.Printf(homeDir)
}

func (c Config) SetUser() {
	// sets the current_user_name to struct
	// then writes the Config struct back to JSON ~/.gatorconfig.json
}

func GetConfigPath() (string, error) {
	return "Hello from internal", nil
}

func write(cfg Config) error {
	return nil
}
