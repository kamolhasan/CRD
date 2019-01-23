package cmd

import (
	"github.com/masudur-rahman/appsCRD/appsclient"
	"github.com/spf13/cobra"
)

var getDeploy = &cobra.Command{
	Use: 	"getDeploy",
	Short: 	"Get all CustomDeployments",
	Run: func(cmd *cobra.Command, args []string) {
		appsclient.GetCustomDeployment()
	},
}

func init() {
	rootCmd.AddCommand(getDeploy)
}
