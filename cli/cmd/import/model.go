package import_cmd

import (
	"fmt"
	types "gh_foundations/internal/pkg/types/terragrunt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	errorStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#f00020"))
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprint(i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	textInput  textinput.Model
	ModulePath string
	Archive    types.IPlanFile
	spinner    spinner.Model
	list       list.Model
	importing  string
	loading    bool
	err        error
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	ti := textinput.New()
	return model{
		spinner:   s,
		list:      list.New(make([]list.Item, 0), itemDelegate{}, 0, 0),
		loading:   true,
		textInput: ti,
	}
}

func (m *model) showLoadingSpinner() tea.Cmd {
	m.loading = true
	return m.spinner.Tick
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		generatePlanFile(m.ModulePath),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case planAndArchiveMsg:
		for _, address := range msg.resourceAddresses {
			m.list.InsertItem(len(m.list.Items()), item(address))
		}
		m.Archive = msg.archive
		m.loading = false

	case terragruntImportMsg:
		m.list.RemoveItem(m.list.Index())
		m.importing = ""
		m.textInput.SetValue("")
		m.loading = false

	case resolveResourceIdMsg:
		m.loading = false
		m.textInput.Focus()
		m.textInput.SetValue(string(msg))

	case errMsg:
		m.loading = false
		m.importing = ""
		m.err = msg.err

	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
		m.textInput.Width = msg.Width
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.Archive.Cleanup()
			return m, tea.Quit

		case "enter":
			if m.err != nil {
				m.err = nil
			} else if m.importing == "" {
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.importing = string(i)
					return m, tea.Sequence(m.showLoadingSpinner(), resolveResourceId(string(i), m.Archive))
				}
			} else {
				return m, tea.Sequence(m.showLoadingSpinner(), runTerragruntImport(m.ModulePath, m.importing, m.textInput.Value()))
			}
		}
	}

	var cmd tea.Cmd

	if m.loading {
		m.spinner, cmd = m.spinner.Update(msg)
	} else if m.importing != "" {
		m.textInput, cmd = m.textInput.Update(msg)
	} else {
		m.list, cmd = m.list.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	if m.err != nil {
		return errorStyle.Render(m.err.Error())
	} else if m.loading {
		return fmt.Sprintf("\n\n %s Loading...", m.spinner.View())
	} else if m.importing != "" {
		return fmt.Sprintf("Enter Import Id:\n\n%s", m.textInput.View())
	}
	return "\n" + m.list.View()
}
