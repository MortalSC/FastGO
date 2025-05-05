package app

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// defaultHomeDir : default home directory
	// The defualt directory for placing the fastgo configuration file
	defaultHomeDir = ".fastgo"

	// defaultConfigNmae : default config file name
	// The default name of the fastgo configuration file
	defaultConfigName = "fg-apiserver.yaml"
)

// onInitialize : initialize the config file
// Set the configuration flie names and environment variables that need to be read, and read their contents into viper
func onInitialize() {
	if configFile != "" {
		// Read the content of the configuration file specified in the command line
		viper.SetConfigFile(configFile)
	} else {
		// Use the default configuration file name and path
		for _, dir := range searchDirs() {
			// Add the 'dir' directory to the search path of the configuration file
			viper.AddConfigPath(dir)
		}

		// Set the file reading format to yaml
		viper.SetConfigType("yaml")

		// Set the name of the configuration file
		viper.SetConfigName(defaultConfigName)
	}

	// Readn the environment variables and set the prefix
	setupEnvironmentVariables()

	// Read the configuration file.
	// If a file name is specified, use the specified configuration file;
	// otherwise, search int the registered search path
	_ = viper.ReadInConfig()
}

// setupEnvironmentVariables : set up the environment variables
func setupEnvironmentVariables() {
	// Allow viper to automatically read environment variables
	viper.AutomaticEnv()
	// set the prefix for environment variables
	viper.SetEnvPrefix("fastgo")

	replace := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replace)
}

// searchDirs returns the default configuration file search path
func searchDirs() []string {
	// Get the home directory of the current user
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return []string{filepath.Join(homeDir, defaultHomeDir), "."}
}

func filePath() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return filepath.Join(home, defaultHomeDir, defaultConfigName)
}
