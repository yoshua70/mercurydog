package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mercurydog",
	Short: "A task scheduler in Go",
	Long: `Mercurydog lets you create tasks and pipelines by leveraging the power of RabbitMQ and SQLite.

Create a new task by providing the command to execute and the RabbitMQ queue to send
message to when the task is completed. You can then add another task to be run after it.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
