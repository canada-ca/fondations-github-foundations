package import_cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var ImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Starts an interactive import process for resources in a Terraform plan.",
	Long:  `This command will start an interactive process to import resources into Terraform state. It uses the results of a terraform plan to determine which resources are available for import.`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		m := initialModel()
		m.ModulePath = args[0]
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func init() {
}
