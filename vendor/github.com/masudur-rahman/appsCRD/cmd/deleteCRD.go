package cmd

import (
	"github.com/masudur-rahman/appsCRD/appsclient"
	"github.com/spf13/cobra"
)

var deleteCRD = &cobra.Command{
	Use: 	"delete",
	Short: 	"Delete CustomResourceDefinition",
	Run: func(cmd *cobra.Command, args []string) {
		appsclient.DeleteCRD()
	},
}

func init() {
	rootCmd.AddCommand(deleteCRD)
}

