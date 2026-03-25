package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

const currentVersion = "0.1.0"

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for updates",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("goboticus %s (%s/%s)\n", currentVersion, runtime.GOOS, runtime.GOARCH)
		fmt.Println("Auto-update not yet implemented.")
		fmt.Println("Check https://github.com/goboticus/goboticus/releases for new versions.")
		return nil
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("goboticus %s\n", currentVersion)
		fmt.Printf("go:        %s\n", runtime.Version())
		fmt.Printf("os/arch:   %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd, versionCmd)
}
