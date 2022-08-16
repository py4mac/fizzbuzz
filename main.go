package main

import (
	"os"

	"github.com/py4mac/fizzbuzz/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	var rootCmd = &cobra.Command{
		Use: "fizzbuzz",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Error(err, "cannot show help")
			}
		},
	}

	rootCmd.AddCommand(cmd.MakeFizzBuzz())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
