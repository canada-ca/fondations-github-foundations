package list

import (
	orgs "gh_foundations/cmd/list/orgs"
	repos "gh_foundations/cmd/list/repos"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list managed resources",
	Long: `list various resources managed by the tool.\n
	Currently supported resources are:\n\n

	- repos\n
	- orgs\n\n`,
}

func init() {
	ListCmd.AddCommand(orgs.OrgsCmd)
	ListCmd.AddCommand(repos.ReposCmd)

}
