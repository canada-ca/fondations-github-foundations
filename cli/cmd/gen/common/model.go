package common

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type model[T any] struct {
	Result *T

	width           int
	height          int
	questions       []IQuestion
	currentQuestion int
	loadingSpinner  spinner.Model
	viewport        viewport.Model
	submitFn        func([]string, *T)
	submitted       bool
	showHelp        bool
}

var (
	submitButtonZoneKey string = "submitButton"
	resetButtonZoneKey  string = "resetButton"
	quitButtonZoneKey   string = "quitButton"
)

var buttonStyling = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).MarginLeft(2).Background(lipgloss.Color("63"))

var viewportKeyBindings viewport.KeyMap = viewport.KeyMap{
	PageDown: key.NewBinding(
		key.WithKeys("pgdown"),
		key.WithHelp("pgdn", "page down"),
	),
	PageUp: key.NewBinding(
		key.WithKeys("pgup"),
		key.WithHelp("pgup", "page up"),
	),
	HalfPageUp: key.NewBinding(
		key.WithKeys("ctrl+u"),
		key.WithHelp("ctrl+u", "½ page up"),
	),
	HalfPageDown: key.NewBinding(
		key.WithKeys("ctrl+d"),
		key.WithHelp("ctrl+d", "½ page down"),
	),
	Up: key.NewBinding(
		key.WithKeys("ctrl+up"),
		key.WithHelp("ctrl+↑", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("ctrl+down"),
		key.WithHelp("ctrl+↓", "down"),
	),
}

func NewModel[T any](questions []IQuestion, defaultValue *T, submitFn func([]string, *T)) model[T] {
	return model[T]{
		loadingSpinner:  spinner.New(),
		questions:       questions,
		currentQuestion: 0,
		submitFn:        submitFn,
		Result:          defaultValue,
		submitted:       false,
		showHelp:        false,
	}
}

func (m model[T]) Init() tea.Cmd {
	return m.loadingSpinner.Tick
}

func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.viewport = viewport.New(m.width, m.height-2)
		m.viewport.KeyMap = viewportKeyBindings
		resizeCmds := make([]tea.Cmd, 0)
		for _, q := range m.questions {
			q.SetDimensions(msg.Width, msg.Height)
			resizeCmds = append(resizeCmds, q.Update(msg))
		}
		m.questions[0].Focus()
		m.viewport.SetContent(m.questions[0].View())
		return m, tea.Batch(resizeCmds...)
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String():
			return m, tea.Quit
		case tea.KeyShiftLeft.String(), tea.KeyShiftRight.String():
			m.questions[m.currentQuestion].Blur()
			if msg.Type == tea.KeyShiftLeft {
				m.currentQuestion = max(0, m.currentQuestion-1)
			} else if msg.Type == tea.KeyShiftRight {
				m.currentQuestion = min(len(m.questions)-1, m.currentQuestion+1)
			}
			m.questions[m.currentQuestion].Focus()
		case "?":
			m.showHelp = !m.showHelp
		case tea.KeyEscape.String():
			m.showHelp = false
		}
	case tea.MouseMsg:
		if msg.Action == tea.MouseActionRelease && msg.Button == tea.MouseButtonLeft {
			if zone.Get(submitButtonZoneKey).InBounds(msg) {
				answers := make([]string, len(m.questions))
				for i, q := range m.questions {
					answers[i] = q.GetAnswer()
				}
				m.submitFn(answers, m.Result)
				m.submitted = true
			} else if zone.Get(resetButtonZoneKey).InBounds(msg) {
				m.reset()
			} else if zone.Get(quitButtonZoneKey).InBounds(msg) {
				return m, tea.Quit
			}
		}
	}

	questionUpdateCmd := m.questions[m.currentQuestion].Update(msg)
	m.viewport.SetContent(m.questions[m.currentQuestion].View())

	viewportModel, viewportUpdateCmd := m.viewport.Update(msg)
	m.viewport = viewportModel
	return m, tea.Batch(viewportUpdateCmd, questionUpdateCmd)
}

func (m model[T]) View() string {
	if m.width == 0 {
		return m.loadingSpinner.View()
	}

	if !m.submitted {
		bottomBar := lipgloss.JoinHorizontal(lipgloss.Center, zone.Mark(submitButtonZoneKey, buttonStyling.Render("Submit")))
		return zone.Scan(
			lipgloss.JoinVertical(lipgloss.Left, m.viewport.View(), bottomBar),
		)
	} else {
		// TODO the static styling can/should be it's own variable since the render function never changes. Any static styling that also has a static input string can also
		// be turned into a variable because the string contents will never change
		return zone.Scan(
			lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, lipgloss.NewStyle().Padding(2).Border(lipgloss.NormalBorder()).Render(
				lipgloss.JoinVertical(lipgloss.Center, lipgloss.NewStyle().MarginBottom(2).Render("Do you want to create another?"), lipgloss.JoinHorizontal(
					lipgloss.Left, zone.Mark(resetButtonZoneKey, buttonStyling.Render("Yes")), "    ", zone.Mark(quitButtonZoneKey, buttonStyling.Render("No")),
				)),
			)),
		)
	}
}

func (m *model[T]) reset() {
	for i := range m.questions {
		m.questions[i].Reset()
		m.questions[i].Blur()
	}
	m.currentQuestion = 0
	m.questions[0].Focus()
	m.submitted = false
}
