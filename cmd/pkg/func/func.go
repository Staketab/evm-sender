package ff

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/staketab/evm-sender/cmd/internal/controller"
	vars "github.com/staketab/evm-sender/cmd/internal/var"
	"os"
	"os/user"
	"path/filepath"
)

func CreateDirectory(path string) error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to controller current user: %s", err)
	}

	fullPath := filepath.Join(usr.HomeDir, path)

	err = os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to cr directory: %s", err)
	}

	vars.InfoLog.Printf("Directory created: %s\n", fullPath)
	return nil
}

func CreateConfigFile(path string) error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to controller current user: %s", err)
	}
	filePath := filepath.Join(usr.HomeDir, path)
	content := []byte(`[DEFAULT]
rpc = ""
private_key = ""
recipient = ""
value = 1000000
gas_limit = 22000
memo = "From Staketab with LOVE!"
txcount = 3
min = 1000000000000000000
max = 9000000000000000000
`)

	err2 := os.WriteFile(filePath, content, 0644)
	if err2 != nil {
		return fmt.Errorf("failed to cr config file: %s", err)
	}

	vars.InfoLog.Printf("Config file created: %s\n", filePath)
	return nil
}

func ReadConfigs() (controller.Config, error) {
	usr, err := user.Current()
	if err != nil {
		return controller.Config{}, fmt.Errorf("Failed to controller current user: %s", err)
	}
	filePath := filepath.Join(usr.HomeDir, vars.ConfigFilePath)
	config := controller.Config{}
	tomlFile, err := toml.LoadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("failed to load config file: %s", err)
	}

	err = tomlFile.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal config: %s", err)
	}

	return config, nil
}
