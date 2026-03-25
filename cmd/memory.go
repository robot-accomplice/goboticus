package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var memoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "Query and manage memory tiers",
}

var memoryWorkingCmd = &cobra.Command{
	Use:   "working",
	Short: "List working memory entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiGet("/api/memory/working")
		if err != nil {
			return err
		}
		printJSON(data)
		return nil
	},
}

var memoryEpisodicCmd = &cobra.Command{
	Use:   "episodic",
	Short: "List episodic memory entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiGet("/api/memory/episodic")
		if err != nil {
			return err
		}
		printJSON(data)
		return nil
	},
}

var memorySemanticCmd = &cobra.Command{
	Use:   "semantic",
	Short: "List semantic memory entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiGet("/api/memory/semantic")
		if err != nil {
			return err
		}
		printJSON(data)
		return nil
	},
}

var memorySearchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search across all memory tiers",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiGet(fmt.Sprintf("/api/memory/search?q=%s", args[0]))
		if err != nil {
			return err
		}
		printJSON(data)
		return nil
	},
}

var memoryStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show memory tier statistics",
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, tier := range []string{"working", "episodic", "semantic"} {
			data, err := apiGet("/api/memory/" + tier)
			if err != nil {
				continue
			}
			if entries, ok := data["entries"].([]any); ok {
				fmt.Printf("%-10s %d entries\n", tier, len(entries))
			} else if entries, ok := data["memories"].([]any); ok {
				fmt.Printf("%-10s %d entries\n", tier, len(entries))
			} else {
				fmt.Printf("%-10s (no entries key)\n", tier)
			}
		}
		return nil
	},
}

func init() {
	memoryCmd.AddCommand(memoryWorkingCmd, memoryEpisodicCmd, memorySemanticCmd, memorySearchCmd, memoryStatsCmd)
	rootCmd.AddCommand(memoryCmd)
}
