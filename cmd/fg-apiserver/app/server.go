package app

import (
	"github.com/MortalSC/FastGO/cmd/fg-apiserver/app/options"
	"github.com/MortalSC/FastGO/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string // Path to configuration file loaded from command-line flag

// NewFastG0Command creates the root cobra.Command for the application
// This constructs the main command with configuration, arguments, and execution logic
// The command hierarchy and execution flow are defined here
func NewFastG0Command() *cobra.Command {
	// Initialize server options with default values
	opts := options.NewServerOptions()

	cmd := &cobra.Command{
		Use:   "fg-apiserver",
		Short: "A very lightweight ful1 go project",
		Long:  "A very lightweight full go project,designed to help beginners quickly learn Go project development.",

		// Silence usage display when errors occur
		SilenceUsage: true,

		// Execution entrypoint with proper error handling
		RunE: func(cmd *cobra.Command, args []string) error {
			// Main execution sequence: config -> validation -> server init -> run
			return run(opts)
		},

		// Strict argument validation
		Args: cobra.NoArgs,
	}

	// Register configuration initialization hook
	cobra.OnInitialize(onInitialize)

	// Persistent flag available to all subcommands
	cmd.PersistentFlags().StringVarP(&configFile,
		"config",   // Flag name
		"c",        // Shorthand
		filePath(), // Default value (implementation not shown)
		"Path to config file in YAML/JSON format", // Description
	)

	// Add --version flag to display version information
	version.AddFlags(cmd.PersistentFlags())

	return cmd
}

// run executes the main application workflow:
// Config unmarshalling
// Config validation
// Server construction
// Server execution
func run(opts *options.ServerOptions) error {
	// if has --version flag, print version and exit
	version.PrintAndExitIfRequested()

	if err := viper.Unmarshal(opts); err != nil {
		return err
	}

	if err := opts.Validate(); err != nil {
		return err
	}

	cfg, err := opts.Config()
	if err != nil {
		return err
	}

	server, err := cfg.NewServer()
	if err != nil {
		return err
	}

	return server.Run()
}
