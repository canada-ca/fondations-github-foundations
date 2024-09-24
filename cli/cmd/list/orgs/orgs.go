/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"errors"
	"fmt"
	"gh_foundations/internal/pkg/functions"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var OrgsCmd = &cobra.Command{
	Use:   "orgs",
	Short: "List managed organizations's slugs.",
	Long: `This command reads the "providers.hcl" files in the "providers" directory and lists the organization slugs that are managed by the tool.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires the path of the \"providers\" directory")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		orgsDir := args[0]

		orgs, err := functions.FindManagedOrgSlugs(orgsDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var orgOut string = "[]"

		if len(orgs) > 0 {
			orgOut = fmt.Sprintf("['%s']", strings.Join(orgs, "', '"))
		}

		fmt.Println(orgOut)

	},
}

func init() {
}
