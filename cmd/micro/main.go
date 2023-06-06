package main

import (
	"github.com/devexps/go-micro/cmd/micro/v2/internal/project"
	"github.com/devexps/go-micro/cmd/micro/v2/internal/run"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:     "micro",
	Short:   "Micro: An elegant toolkit for Go microservices.",
	Long:    `Micro: An elegant toolkit for Go microservices.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(run.CmdRun)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
