/*
Package cmd is all the available commands for the CLI application
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/mrz1836/paymail-inspector/chalker"
	"github.com/mrz1836/paymail-inspector/paymail"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// Default flag values for various commands
var (
	amount             uint64
	brfcAuthor         string
	brfcTitle          string
	brfcVersion        string
	configFile         string
	generateDocs       bool
	nameServer         string
	port               int
	priority           int
	protocol           string
	purpose            string
	satoshis           uint64
	serviceName        string
	signature          string
	skipBrfcValidation bool
	skipDnsCheck       bool
	skipPki            bool
	skipPublicProfile  bool
	skipSrvCheck       bool
	skipSSLCheck       bool
	skipTracing        bool
	weight             int
)

// Defaults for the application
const (
	configDefault     = "paymail-inspector" // Config file and application name
	defaultDomainName = "moneybutton.com"   // Used in examples
	defaultNameServer = "8.8.8.8"           // Default DNS NameServer
	docsLocation      = "docs/commands"     // Default location for command documentation
)

// These are keys for known flags that are used in the configuration
const (
	flagBsvAlias     = "bsvalias"
	flagSenderHandle = "sender-handle"
	flagSenderName   = "sender-name"
)

// Version is set manually (also make:build overwrites this value from Github's latest tag)
var Version = "v0.0.20"

// rootCmd represents the base command when called without any sub-commands
var rootCmd = &cobra.Command{
	DisableAutoGenTag: true,
	Use:               configDefault,
	Short:             "Inspect, validate or resolve paymail domains and addresses",
	Example:           configDefault + " -h",
	Long: chalk.Green.Color(`
__________                             .__.__    .___                                     __                
\______   \_____  ___.__. _____ _____  |__|  |   |   | ____   ____________   ____   _____/  |_  ___________ 
 |     ___/\__  \<   |  |/     \\__  \ |  |  |   |   |/    \ /  ___/\____ \_/ __ \_/ ___\   __\/  _ \_  __ \
 |    |     / __ \\___  |  Y Y  \/ __ \|  |  |__ |   |   |  \\___ \ |  |_> >  ___/\  \___|  | (  <_> )  | \/
 |____|    (____  / ____|__|_|  (____  /__|____/ |___|___|  /____  >|   __/ \___  >\___  >__|  \____/|__|   
                \/\/          \/     \/                   \/     \/ |__|        \/     \/     `+Version) + `
` + chalk.Yellow.Color("Author: MrZ © 2020 github.com/mrz1836/"+configDefault) + `

This CLI tool can help you inspect, validate or resolve a paymail domain/address.

Help contribute via Github!
`,
	Version: Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	// Run root command
	er(rootCmd.Execute())

	// Generate docs from all commands
	if generateDocs {

		// Replace the colorful logs in terminal (displays in Cobra docs) (color numbers generated)
		replacer := strings.NewReplacer("[32m", "```", "[33m", "```\n", "[39m", "", "[36m", "", "\u001B", "")
		rootCmd.Long = replacer.Replace(rootCmd.Long)

		// Loop all command, adjust the Long description, re-add command
		for _, command := range rootCmd.Commands() {
			rootCmd.RemoveCommand(command)
			command.Long = replacer.Replace(command.Long)
			rootCmd.AddCommand(command)
		}

		// Generate the markdown docs
		if err := doc.GenMarkdownTree(rootCmd, docsLocation); err != nil {
			chalker.Log(chalker.ERROR, fmt.Sprintf("error generating docs: %s", err.Error()))
			return
		}
		chalker.Log(chalker.SUCCESS, fmt.Sprintf("successfully generated documentation for %d commands", len(rootCmd.Commands())))
	}
}

func init() {

	// Load the configuration
	cobra.OnInitialize(initConfig)

	// Add config option
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/."+configDefault+".yaml)")

	// Add document generation for all commands
	rootCmd.PersistentFlags().BoolVar(&generateDocs, "docs", false, "Generate docs from all commands (./"+docsLocation+")")

	// Add a toggle for request tracing
	rootCmd.PersistentFlags().BoolVarP(&skipTracing, "skip-tracing", "t", false, "Turn off request tracing information")

	// Add a bsvalias version to target
	rootCmd.PersistentFlags().String(flagBsvAlias, paymail.DefaultBsvAliasVersion, fmt.Sprintf("The %s version", flagBsvAlias))
	er(viper.BindPFlag(flagBsvAlias, rootCmd.PersistentFlags().Lookup(flagBsvAlias)))
}

// er is a basic helper method to catch errors loading the application
func er(err error) {
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if configFile != "" {

		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {

		// Find home directory.
		home, err := homedir.Dir()
		er(err)

		// Search config in home directory with name "."+configDefault (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("." + configDefault)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		chalker.Log(chalker.INFO, fmt.Sprintf("...loaded config file: %s", viper.ConfigFileUsed()))
	}
}
