package cmd

import (
	"github.com/masudur-rahman/appsCRD/appsclient"
	"github.com/spf13/cobra"
)

var deploy = &cobra.Command{
	Use: 	"deploy",
	Short: 	"Create a CustomDeployment",
	Run: func(cmd *cobra.Command, args []string) {
		appsclient.CreateCustomDeployment()
	},
}

func init() {
	rootCmd.AddCommand(deploy)
}
