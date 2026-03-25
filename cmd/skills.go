package cmd

import (
	"github.com/spf13/cobra"
)

var skillsCmd = &cobra.Command{
	Use:   "skills",
	Short: "Manage agent skills",
}

var skillsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List loaded skills",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiGet("/api/skills")
		if err != nil {
			return err
		}
		printJSON(data)
		return nil
	},
}

var skillsReloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload skills from disk",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiPost("/api/skills/reload", nil)
		if err != nil {
			return err
		}
		printJSON(data)
		return nil
	},
}

var skillsCatalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "Browse skill catalog",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiGet("/api/skills/catalog")
		if err != nil {
			return err
		}
		printJSON(data)
		return nil
	},
}

func init() {
	skillsCmd.AddCommand(skillsListCmd, skillsReloadCmd, skillsCatalogCmd)
	rootCmd.AddCommand(skillsCmd)
}
