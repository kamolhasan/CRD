package cmd

import (
	"github.com/masudur-rahman/appsCRD/appsclient"
	"github.com/spf13/cobra"
)

var createCRD= &cobra.Command{
	Use: 	"create",
	Short: 	"Create CustomResourceDefinition",
	Run: func(cmd *cobra.Command, args []string) {
		appsclient.CreateCRD()
	},
}

func init() {
	rootCmd.AddCommand(createCRD)
}
