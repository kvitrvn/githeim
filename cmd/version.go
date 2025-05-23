package cmd

import (
	"github.com/kvitrvn/githeim/internal/build"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Githeim",
	Long:  `All software has versions. This is Githeim's`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Githeim version %s, build %s \n", build.Version, build.Commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
