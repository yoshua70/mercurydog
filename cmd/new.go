package cmd

import (
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new job or a new pipeline",
	Long: `A command to create a new job or a new pipeline.

The command has no use on its own but has two subcommands: 'job' and 'pipeline'
Use the subcommand 'job' or 'pipeline' to create a new job or a new pipeline respectively.

Examples:
  mercuryfox new job [flags]
  mercuryfox new pipeline [flags]
`,

	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
