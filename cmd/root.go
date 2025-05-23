package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "githeim",
	Short: "Githeim is a CLI tool for managing Git repositories.",
	Long:  `Githeim is a CLI tool for managing Git repositories. It provides various commands to interact with Git repositories, including creating, cloning, and managing branches.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This is the default action when no subcommand is provided.
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
