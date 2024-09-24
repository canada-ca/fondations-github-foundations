package gen

import (
	repositoryset "gh_foundations/cmd/gen/repository_set"
	teamset "gh_foundations/cmd/gen/team_set"

	"github.com/spf13/cobra"
)

var GenCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate HCL input for GitHub Foundations.",
	Long:  `Generate HCL input for GitHub Foundations. This tool is used to generate HCL input for GitHub Foundations from state files output by terraformer.`,
}

func init() {
	GenCmd.AddCommand(repositoryset.GenRepositorySetCmd)
	GenCmd.AddCommand(teamset.GenTeamSetCmd)
}
