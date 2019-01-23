package cmd

import (
	"github.com/masudur-rahman/appsCRD/appsclient"
	"github.com/spf13/cobra"
)

var deleteDeploy = &cobra.Command{
	Use: 	"deleteDeploy",
	Short: 	"Delete a CustomDeployment",
	Run: func(cmd *cobra.Command, args []string) {
		appsclient.DeleteCustomDeployment()
	},
}

func init() {
	rootCmd.AddCommand(deleteDeploy)
}
