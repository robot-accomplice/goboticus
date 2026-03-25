package cmd

import (
	"github.com/spf13/cobra"
)

var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Wallet balance and address management",
}

var walletBalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Show wallet balance",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiGet("/api/wallet/balance")
		if err != nil {
			return err
		}
		printJSON(data)
		return nil
	},
}

var walletAddressCmd = &cobra.Command{
	Use:   "address",
	Short: "Show wallet address",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiGet("/api/wallet/address")
		if err != nil {
			return err
		}
		printJSON(data)
		return nil
	},
}

func init() {
	walletCmd.AddCommand(walletBalanceCmd, walletAddressCmd)
	rootCmd.AddCommand(walletCmd)
}
