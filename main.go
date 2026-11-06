// This is the starting phase of the app
package main

import (
	"fmt"
	"os"
	"strings"

	cmd "OllaGoClip/cmd"
	"OllaGoClip/utils"

	"github.com/alecthomas/kong" //golang cmd line parser
	kongtoml "github.com/alecthomas/kong-toml"
	"github.com/charmbracelet/log"
)

type Cli struct {
	Version VersionFlag `help:"Show version"`
	Debug   DebugFlag   `help:"Enable debug info"`

	// Do is executed when no sucommand is given
	Do      cmd.DoCmd     `cmd:"" default:"withargs" help:"Send <prompt> and clipboard content to Ollama" aliases:"exec,ask,run"`
	Config  cmd.ConfigCmd `cmd:"" help:"Interact with default config at $HOME/.OllaClip.toml"`
	Prompts cmd.PromptCmd `cmd:"" help:"Interact with prompts"`
	Systems cmd.SystemCmd `cmd:"" help:"Interact with system prompts"`
}

type VersionFlag bool
type DebugFlag bool

var cli Cli
var version string

// function to print version of product
func (v VersionFlag) BeforeApply() error {
	if version == "" {
		version = "unversioned"
	}
	fmt.Print(version)
	os.Exit(0)
	return nil
}

// if user types --debug
func (v DebugFlag) BeforeApply() error {
	utils.Logger.SetLevel(log.DebugLevel)
	return nil
}

func main() {
	// If no args given, print main help, instead of just crashing.
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "--help")
	}

	desc := `
	Takes the clipboard content + given prompt and sends it to Ollama.

	If no subcommand is given, executes the 'do' command.
	
	Examples:
		$ OllaClip "let's talk about the clipboard content" -p

		$ OllaClip do "let's talk about the clipboard content" -p

		$ OllaClip config list

		$ OllaClip prompts add "let" "let's talk about the clipboard content"
		`
	comp := strings.ReplaceAll(desc, "\n\n", " \n")

	// Load CLI
	opt := kong.Configuration(kongtoml.Loader, []string{"~/.OllaClip.toml"}...)
	ctx := kong.Parse(&cli,
		kong.Name("OllaClip"),
		kong.Description(comp),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{Compact: true, NoExpandSubcommands: true, FlagsLast: true}),
		opt,
	)
	_, err := kong.New(&cli, opt)
	ctx.FatalIfErrorf(err)
	// Run main command
	err = ctx.Run() //assignes the task to respective handlers (routing basically)
	if err != nil {
		utils.Logger.Error(err)
		os.Exit(1)
	}

}
