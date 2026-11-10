package cmd

import (
	"fmt"
	"os"

	"github.com/GandrothuNagaSaiRishitha/OllaClip/utils"
	"github.com/pelletier/go-toml/v2"
)

type ConfigCmd struct {
	Init InitConfigCmd `cmd:"" help:"Create default config at $HOME/.OllaClip.toml"`
	List ListConfigCmd `cmd:"" help:"List contents of config"`
}

// if user uses -f then it is forced to create a new config file.
type InitConfigCmd struct {
	Force bool `short:"f" help:"If true, will replace current config with default config, if file exists."`
}

type ListConfigCmd struct {
}

// --- Plain config command
//do not do anything if user just types in cmd command(by mistake, probably)

func (c *ConfigCmd) Run() error {
	return nil
}

// --- config init command
//if the user types in init command, then execute this function.

func (c *InitConfigCmd) Run() error {
	configPath, err := utils.GetConfigPath()
	if err != nil {
		return err
	}
	// If we are not forcing a write, check if file already exists(its ricky to force)
	if !c.Force {
		_, err = os.Stat(configPath)
		if os.IsNotExist(err) {
			utils.Logger.Infof("Config file does not exist yet, will create default one now...")
		} else {
			return fmt.Errorf("config already exists at %s", configPath)
		}
	}

	err = writeConfigToFile(configPath)
	if err != nil {
		return err
	}

	return nil
}

func writeConfigToFile(filename string) error {
	defaultFlags := utils.Flags{Model: "qwen3:0.6b", OllamaHost: "127.0.0.1:11434"}
	defaultPrompts := utils.Prompts{Entries: map[string]string{"example": "Using the '--expand' flag with the word 'example' as the prompt will expand into this string."}}
	defaultConfig := utils.ConfigToml{Flags: defaultFlags, Prompts: defaultPrompts}
	content, err := toml.Marshal(defaultConfig)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

// --- config list command

func (c *ListConfigCmd) Run() error {
	configPath, err := utils.GetConfigPath()
	if err != nil {
		return err
	}
	config, err := utils.ReadConfigAsBytes()
	if err != nil {
		return err
	}
	utils.Logger.Infof("Content of default file %s\n", configPath)

	fmt.Println(string(config))
	return nil
}
