package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display agent health and status",
	RunE: func(cmd *cobra.Command, args []string) error {
		health, err := apiGet("/api/health")
		if err != nil {
			return fmt.Errorf("health check failed: %w", err)
		}

		fmt.Printf("Status:    %v\n", health["status"])
		fmt.Printf("Uptime:    %v\n", health["uptime"])
		fmt.Printf("Go:        %v\n", health["go"])

		if providers, ok := health["providers"].([]any); ok {
			fmt.Printf("Providers: %d\n", len(providers))
			for _, p := range providers {
				pm, _ := p.(map[string]any)
				fmt.Printf("  - %v (%v) state=%v\n", pm["name"], pm["format"], pm["state"])
			}
		}

		agent, err := apiGet("/api/agent/status")
		if err == nil {
			fmt.Printf("\nAgent:     %v\n", agent["status"])
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
