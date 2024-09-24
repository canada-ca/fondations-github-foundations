package teamset

import (
	"fmt"
	"gh_foundations/cmd/gen/common"
	githubfoundations "gh_foundations/internal/pkg/types/github_foundations"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	yaml "gopkg.in/yaml.v2"
)

var questions []common.IQuestion = []common.IQuestion{
	common.NewTextQuestion(
		"Enter the name of the team",
		"",
	),
	common.NewTextQuestion(
		"Enter the description for the team",
		"",
	),
	common.NewSelectQuestion(
		"Select the level of privacy for the team",
		[]string{
			"secret",
			"closed",
		},
	),
	common.NewListQuestion(
		"Enter the team maintainers",
	),
	common.NewListQuestion(
		"Enter the team members",
	),
	common.NewTextQuestion(
		"Enter the id of the parent team if any",
		"",
	),
}

func submitFunc(answers []string, teamSet *githubfoundations.TeamSetInput) {
	team := new(githubfoundations.TeamInput)
	team.Name = answers[0]
	team.Description = answers[1]
	team.Privacy = answers[2]
	maintainers := make([]string, 0)
	err := yaml.Unmarshal([]byte(answers[3]), &maintainers)
	if err != nil {
		fmt.Println("Error converting maintainers input to array of strings:", err)
		os.Exit(1)
	}
	team.Maintainers = maintainers

	members := make([]string, 0)
	err = yaml.Unmarshal([]byte(answers[4]), &members)
	if err != nil {
		fmt.Println("Error converting members input to array of strings:", err)
		os.Exit(1)
	}
	team.Members = members

	team.ParentId = answers[5]
	teamSet.Teams = append(teamSet.Teams, team)
}

func runInteractive() (*githubfoundations.TeamSetInput, error) {
	m := common.NewModel(questions, new(githubfoundations.TeamSetInput), submitFunc)
	if _, err := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion()).Run(); err != nil {
		return nil, err
	}
	return m.Result, nil
}
