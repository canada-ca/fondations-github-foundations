package teamset

import (
	"fmt"
	"os"

	"gh_foundations/cmd/gen/common"
	githubfoundations "gh_foundations/internal/pkg/types/github_foundations"

	zone "github.com/lrstanley/bubblezone"
	"github.com/spf13/cobra"
)

var GenTeamSetCmd = &cobra.Command{
	Use:   "team_set",
	Short: "Generates an hcl file that contains a team set input. Can only be run interactively.",
	Long:  `Generates an hcl file that contains a team set input. Can only be run interactively.`,
	Run: func(cmd *cobra.Command, args []string) {
		zone.NewGlobal()
		var teamSet *githubfoundations.TeamSetInput
		var err error

		teamSet, err = runInteractive()
		if err != nil {
			fmt.Println("Error running interactive mode:", err)
			os.Exit(1)
		}

		if err := common.OutputHCLToFile("team_set.inputs.hcl", teamSet); err != nil {
			fmt.Println("Error writing hcl file:", err)
			os.Exit(1)
		}
	},
}
