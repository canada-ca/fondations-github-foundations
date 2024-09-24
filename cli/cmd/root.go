package cmd

import (
	"gh_foundations/cmd/check"
	"gh_foundations/cmd/gen"
	import_cmd "gh_foundations/cmd/import"
	"gh_foundations/cmd/list"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh_foundations",
	Short: "The GitHub Foundations CLI tool.",
	Long: `The GitHub Foundations CLI tool is a tool to manage GitHub resources
	and to read the state of the resources managed by the tool.\n`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tg_import.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(import_cmd.ImportCmd)
	rootCmd.AddCommand(gen.GenCmd)
	rootCmd.AddCommand(check.CheckCmd)
	rootCmd.AddCommand(list.ListCmd)
}
