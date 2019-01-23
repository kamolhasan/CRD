package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)


var rootCmd = &cobra.Command{
	Use:	"appsCRD",
	Short:	"Short description of appsCRD",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Welcome from AppsCRD...!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
