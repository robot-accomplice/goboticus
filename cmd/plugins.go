package cmd

import (
	"github.com/spf13/cobra"
)

var pluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "Manage plugins",
}

var pluginsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed plugins",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiGet("/api/plugins/catalog/install")
		if err != nil {
			// Fallback: try the main plugins endpoint if it exists.
			data, err = apiGet("/api/skills") // Plugins share catalog
			if err != nil {
				return err
			}
		}
		printJSON(data)
		return nil
	},
}

func init() {
	pluginsCmd.AddCommand(pluginsListCmd)
	rootCmd.AddCommand(pluginsCmd)
}
