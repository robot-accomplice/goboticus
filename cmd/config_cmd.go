package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View and manage configuration",
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Display current configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiGet("/api/config")
		if err != nil {
			return err
		}
		printJSON(data)
		return nil
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Get a specific config value",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		val := viper.Get(args[0])
		if val == nil {
			return fmt.Errorf("key %q not found", args[0])
		}
		fmt.Printf("%s = %v\n", args[0], val)
		return nil
	},
}

var configValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := loadConfig()
		if err != nil {
			return fmt.Errorf("validation failed: %w", err)
		}
		fmt.Println("Configuration is valid.")
		return nil
	},
}

var configCapabilitiesCmd = &cobra.Command{
	Use:   "capabilities",
	Short: "List available capabilities",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiGet("/api/config/capabilities")
		if err != nil {
			return err
		}
		printJSON(data)
		return nil
	},
}

func init() {
	configCmd.AddCommand(configShowCmd, configGetCmd, configValidateCmd, configCapabilitiesCmd)
	rootCmd.AddCommand(configCmd)
}
