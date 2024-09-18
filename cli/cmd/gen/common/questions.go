package common

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
	zone "github.com/lrstanley/bubblezone"
	yaml "gopkg.in/yaml.v2"
)

var promptStyle = lipgloss.NewStyle().Margin(1).MarginLeft(0)
var inputStyle lipgloss.Style = lipgloss.NewStyle().Border(lipgloss.ThickBorder()).BorderTop(false).BorderLeft(false).BorderRight(false)
var selectedInputStyle lipgloss.Style = inputStyle.BorderForeground(lipgloss.Color("63"))

func renderPrompt(prompt string) string {
	return promptStyle.Render(fmt.Sprintf("%s:", prompt))
}

func removeListCursor(style lipgloss.Style) lipgloss.Style {
	return style.SetString(" ")
}

func addListCursor(style lipgloss.Style) lipgloss.Style {
	return style.SetString(">")
}

type IQuestion interface {
	GetAnswer() string
	View() string
	Update(msg tea.Msg) tea.Cmd
	SetDimensions(width, height int)
	Focus()
	Blur()
	Reset()
}

type TextQuestion struct {
	prompt       string
	inputModel   textinput.Model
	defaultValue string
}

func NewTextQuestion(prompt string, defaultValue string) *TextQuestion {
	m := textinput.New()
	m.Prompt = ""
	m.SetValue(defaultValue)
	return &TextQuestion{
		prompt:       prompt,
		inputModel:   m,
		defaultValue: defaultValue,
	}
}

func (t *TextQuestion) Reset() {
	t.inputModel.SetValue(t.defaultValue)
}

func (t *TextQuestion) GetAnswer() string {
	return t.inputModel.Value()
}

func (t *TextQuestion) View() string {
	inputRenderFn := inputStyle.Render
	if t.inputModel.Focused() {
		inputRenderFn = selectedInputStyle.Render
	}
	return lipgloss.JoinVertical(lipgloss.Left, renderPrompt(t.prompt), inputRenderFn(t.inputModel.View()))
}

func (t *TextQuestion) Update(msg tea.Msg) tea.Cmd {
	m, cmd := t.inputModel.Update(msg)
	t.inputModel = m
	return cmd
}

func (t *TextQuestion) SetDimensions(width, height int) {
	t.inputModel.Width = max(width/2, 45)
}

func (t *TextQuestion) Focus() {
	t.inputModel.Focus()
}

func (t *TextQuestion) Blur() {
	t.inputModel.Blur()
}

type item struct {
	strValue string
	value    any
}

func (i item) FilterValue() string { return i.strValue }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	fn := lipgloss.NewStyle().PaddingLeft(4).Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return m.Styles.FilterCursor.Foreground(lipgloss.Color("170")).PaddingLeft(2).Render("") + lipgloss.NewStyle().Foreground(lipgloss.Color("170")).Render(strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(i.strValue))
}

type SelectQuestion struct {
	prompt string
	model  list.Model
}

func NewSelectQuestion[T any](prompt string, selection []T) *SelectQuestion {
	items := make([]list.Item, len(selection))
	for i, s := range selection {
		items[i] = item{
			value:    s,
			strValue: fmt.Sprint(s),
		}
	}
	listModel := list.New(items, itemDelegate{}, 0, 0)
	listModel.SetShowTitle(false)
	listModel.SetShowStatusBar(false)
	listModel.Styles.FilterCursor = lipgloss.NewStyle().SetString(">")
	return &SelectQuestion{
		prompt: prompt,
		model:  listModel,
	}
}

func (s *SelectQuestion) Reset() {
	s.model.Select(0)
}

func (s *SelectQuestion) GetAnswer() string {
	switch rawValue := s.model.SelectedItem().(item).value.(type) {
	case string:
		return rawValue
	default:
		bytes, err := yaml.Marshal(rawValue)
		if err != nil {
			panic(err)
		}
		return string(bytes)
	}
}

func (s *SelectQuestion) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, renderPrompt(s.prompt), s.model.View())
}

func (s *SelectQuestion) Update(msg tea.Msg) tea.Cmd {
	m, cmd := s.model.Update(msg)
	s.model = m
	return cmd
}

func (s *SelectQuestion) SetDimensions(width, height int) {
	s.model.SetWidth(width)
	s.model.SetHeight(height - 1)
}

func (s *SelectQuestion) Focus() {
}

func (s *SelectQuestion) Blur() {
}

type listQuestionMode int

const (
	adding listQuestionMode = iota
	editing
	unfocused
)

type ListQuestion struct {
	prompt string
	values list.Model
	input  textinput.Model
	state  listQuestionMode
}

func NewListQuestion(prompt string) *ListQuestion {
	listModel := list.New(make([]list.Item, 0), itemDelegate{}, 0, 0)
	listModel.SetShowTitle(false)
	listModel.SetShowStatusBar(false)
	listModel.Styles.FilterCursor = removeListCursor(lipgloss.NewStyle())

	inputModel := textinput.New()
	inputModel.Prompt = ""
	return &ListQuestion{
		prompt: prompt,
		values: listModel,
		input:  inputModel,
	}
}

func (l *ListQuestion) Reset() {
	l.input.SetValue("")
	l.values.SetItems(make([]list.Item, 0))
}

func (l *ListQuestion) GetAnswer() string {
	items := l.values.Items()
	values := make([]any, len(items))
	for i, it := range items {
		values[i] = it.(item).value
	}

	bytes, err := yaml.Marshal(values)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (l *ListQuestion) View() string {
	inputRenderFn := inputStyle.Render
	if l.input.Focused() {
		inputRenderFn = selectedInputStyle.Render
	}
	return lipgloss.JoinVertical(lipgloss.Left, renderPrompt(l.prompt), inputRenderFn(l.input.View()), l.values.View())
}

func (l *ListQuestion) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyDelete:
			if l.state == editing {
				l.values.RemoveItem(l.values.Index())
			}
		case tea.KeyEnter:
			inputValue := item{
				strValue: l.input.Value(),
				value:    l.input.Value(),
			}

			if l.state == adding {
				insertCmd := l.values.InsertItem(len(l.values.Items()), inputValue)
				l.input.Reset()
				return insertCmd
			} else if l.state == editing {
				idx := l.values.Index()
				l.values.RemoveItem(idx)
				l.values.InsertItem(idx, inputValue)
			}
		case tea.KeyUp, tea.KeyDown:
			if l.state == editing {
				l.selectItem(msg.Type)
			}
			return nil
		case tea.KeyTab:
			l.switchState(msg)
		default:
			inputModel, inputUpdateCmd := l.input.Update(msg)
			l.input = inputModel
			return inputUpdateCmd
		}
	}

	return nil
}

func (l *ListQuestion) SetDimensions(width, height int) {
	l.input.Width = max(width/2, 45)
	l.values.SetWidth(width)
	l.values.SetHeight(height - 5)
}

func (l *ListQuestion) Focus() {
	l.input.Focus()
	l.state = adding
}

func (l *ListQuestion) Blur() {
	l.input.Blur()
	l.state = unfocused
}

func (l *ListQuestion) selectItem(keyPressed tea.KeyType) {
	index := l.values.Index()
	switch keyPressed {
	case tea.KeyUp:
		index = max(0, index-1)
	case tea.KeyDown:
		index = min(len(l.values.Items())-1, index+1)
	case tea.KeyTab:
		index = 0
	}
	l.values.Select(index)
	l.input.SetValue(l.values.SelectedItem().(item).strValue)
}

func (l *ListQuestion) switchState(msg tea.KeyMsg) {
	if l.state == editing {
		l.state = adding
		l.values.Styles.FilterCursor = removeListCursor(l.values.Styles.FilterCursor)
		l.input.SetValue("")
	} else if l.state == adding && len(l.values.Items()) > 0 {
		l.state = editing
		l.values.Styles.FilterCursor = addListCursor(l.values.Styles.FilterCursor)
		l.selectItem(msg.Type)
	}
}

const keyValueSeparator string = " = "

type keyValuePair struct {
	key   string
	value string
}

type KeyValueListQuestion struct {
	prompt          string
	listModel       list.Model
	keyInputModel   textinput.Model
	valueInputModel textinput.Model
	keyValueMap     map[string]string
	state           listQuestionMode
}

func NewKeyValueListQuestion(prompt string) *KeyValueListQuestion {
	listModel := list.New(make([]list.Item, 0), itemDelegate{}, 0, 0)
	listModel.SetShowTitle(false)
	listModel.SetShowStatusBar(false)
	listModel.Styles.FilterCursor = removeListCursor(lipgloss.NewStyle())

	keyInputModel := textinput.New()
	keyInputModel.Prompt = ""

	valueInputModel := textinput.New()
	valueInputModel.Prompt = ""

	return &KeyValueListQuestion{
		prompt:          prompt,
		listModel:       listModel,
		keyInputModel:   keyInputModel,
		valueInputModel: valueInputModel,
		keyValueMap:     make(map[string]string),
	}
}

func (k *KeyValueListQuestion) Reset() {
	k.keyInputModel.SetValue("")
	k.valueInputModel.SetValue("")
	k.listModel.SetItems(make([]list.Item, 0))
	k.keyValueMap = make(map[string]string)
}

func (k *KeyValueListQuestion) GetAnswer() string {
	bytes, err := yaml.Marshal(k.keyValueMap)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (k *KeyValueListQuestion) View() string {
	keyInput := inputStyle.Render(k.keyInputModel.View())
	valueInput := inputStyle.Render(k.valueInputModel.View())
	if k.keyInputModel.Focused() {
		keyInput = selectedInputStyle.Render(k.keyInputModel.View())
	} else if k.valueInputModel.Focused() {
		valueInput = selectedInputStyle.Render(k.valueInputModel.View())
	}
	inputTopBar := lipgloss.JoinHorizontal(lipgloss.Left, keyInput, "  ", valueInput)
	return lipgloss.JoinVertical(lipgloss.Left, renderPrompt(k.prompt), inputTopBar, k.listModel.View())
}

func (k *KeyValueListQuestion) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyDelete:
			if k.state == editing {
				k.listModel.RemoveItem(k.listModel.Index())
			}
		case tea.KeyEnter:
			if k.state == editing {
				k.listModel.RemoveItem(k.listModel.Index())
			}
			k.putEntry(k.keyInputModel.Value(), k.valueInputModel.Value())
		case tea.KeyUp, tea.KeyDown:
			if k.state == editing {
				k.selectItem(msg.Type)
			}
			return nil
		case tea.KeyShiftTab:
			if k.state != unfocused {
				if k.keyInputModel.Focused() {
					k.keyInputModel.Blur()
					k.valueInputModel.Focus()
				} else {
					k.keyInputModel.Focus()
					k.valueInputModel.Blur()
				}
			}
		case tea.KeyTab:
			k.switchState(msg)
		default:
			keyInputModel, keyInputUpdateCmd := k.keyInputModel.Update(msg)
			k.keyInputModel = keyInputModel
			valueInputModel, valueInputUpdateCmd := k.valueInputModel.Update(msg)
			k.valueInputModel = valueInputModel
			return tea.Batch(keyInputUpdateCmd, valueInputUpdateCmd)
		}
	}

	return nil
}

func (k *KeyValueListQuestion) SetDimensions(width, height int) {
	k.keyInputModel.Width = width/2 - 5
	k.valueInputModel.Width = width/2 - 5
	k.listModel.SetWidth(width)
	k.listModel.SetHeight(height - 5)
}

func (k *KeyValueListQuestion) Focus() {
	k.keyInputModel.Focus()
	k.valueInputModel.Blur()
	k.state = adding
}

func (k *KeyValueListQuestion) Blur() {
	k.keyInputModel.Blur()
	k.valueInputModel.Blur()
	k.state = unfocused
}

func (l *KeyValueListQuestion) createItem(key string, value string) item {
	return item{
		strValue: strings.Join([]string{key, value}, keyValueSeparator),
		value: keyValuePair{
			key:   key,
			value: value,
		},
	}
}

func (k *KeyValueListQuestion) getKeyValuePair(idx int) keyValuePair {
	item := k.listModel.Items()[idx].(item)
	return item.value.(keyValuePair)
}

func (k *KeyValueListQuestion) selectItem(keyPressed tea.KeyType) {
	index := k.listModel.Index()
	switch keyPressed {
	case tea.KeyUp:
		index = max(0, index-1)
	case tea.KeyDown:
		index = min(len(k.listModel.Items())-1, index+1)
	case tea.KeyTab:
		index = 0
	}
	k.listModel.Select(index)
	kvp := k.getKeyValuePair(index)
	k.keyInputModel.SetValue(kvp.key)
	k.valueInputModel.SetValue(kvp.value)
}

func (k *KeyValueListQuestion) switchState(msg tea.KeyMsg) {
	if k.state == editing {
		k.state = adding
		k.listModel.Styles.FilterCursor = removeListCursor(k.listModel.Styles.FilterCursor)
		k.keyInputModel.SetValue("")
		k.valueInputModel.SetValue("")
	} else if k.state == adding && len(k.listModel.Items()) > 0 {
		k.state = editing
		k.listModel.Styles.FilterCursor = addListCursor(k.listModel.Styles.FilterCursor)
		k.selectItem(msg.Type)
	}
}

func (k *KeyValueListQuestion) putEntry(key string, value string) tea.Cmd {
	_, existing := k.keyValueMap[key]
	k.keyValueMap[key] = value
	idx := len(k.listModel.Items())
	newItem := k.createItem(key, value)
	if existing {
		for i := 0; i < idx; i++ {
			kvp := k.getKeyValuePair(i)
			if kvp.key == key {
				k.listModel.RemoveItem(i)
				return k.listModel.InsertItem(i, newItem)
			}
		}
	}
	k.keyInputModel.Reset()
	k.valueInputModel.Reset()
	return k.listModel.InsertItem(idx, newItem)
}

type CompositeQuestion struct {
	title                 string
	questions             []IQuestion
	keys                  []string
	questionZonePrefix    string
	focusedQuestion       int
	focusedQuestionOffset int
	questionHeight        int
	questionWidth         int
}

type CompositeQuestionEntry struct {
	Key      string
	Question IQuestion
}

var compositeQuestionStyle lipgloss.Style = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).MarginLeft(1).MarginRight(1).PaddingLeft(1).PaddingRight(1)
var selectedCompositeQuestionStyle lipgloss.Style = compositeQuestionStyle.BorderForeground(lipgloss.Color("63"))

func NewCompositeQuestion(title string, questions []CompositeQuestionEntry) *CompositeQuestion {
	q := make([]IQuestion, len(questions))
	k := make([]string, len(questions))
	for i := range questions {
		q[i] = questions[i].Question
		k[i] = questions[i].Key
	}
	return &CompositeQuestion{
		title:                 title,
		questions:             q,
		keys:                  k,
		questionZonePrefix:    uuid.New().String(),
		focusedQuestion:       0,
		focusedQuestionOffset: 0,
		questionHeight:        10,
	}
}

func (q *CompositeQuestion) Reset() {
	for _, question := range q.questions {
		question.Reset()
	}
}

func (q *CompositeQuestion) GetAnswer() string {
	answers := make(map[string]interface{})
	for i, k := range q.keys {
		var a interface{}
		err := yaml.Unmarshal([]byte(q.questions[i].GetAnswer()), &a)
		if err != nil {
			panic(err)
		}
		answers[k] = a
	}
	bytes, err := yaml.Marshal(answers)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (q *CompositeQuestion) SetDimensions(width, height int) {
	q.questionHeight = max(height/3, q.questionHeight)
	q.questionWidth = width
	for _, question := range q.questions {
		question.SetDimensions(q.questionWidth, q.questionHeight)
	}
}

func (q *CompositeQuestion) Focus() {
	q.questions[q.focusedQuestion].Focus()
}

func (q *CompositeQuestion) Blur() {
	for _, question := range q.questions {
		question.Blur()
	}
}

func (q *CompositeQuestion) View() string {
	views := make([]string, len(q.questions))
	for i, question := range q.questions {
		renderFn := compositeQuestionStyle.Width(q.questionWidth - 5).Render
		if i == q.focusedQuestion {
			renderFn = selectedCompositeQuestionStyle.Width(q.questionWidth - 5).Render
		}

		views[i] = zone.Mark(q.getQuestionZoneMarkKey(i), renderFn(question.View()))
	}

	if len(q.title) > 0 {
		views = append([]string{q.title}, views...)
	}

	return lipgloss.JoinVertical(lipgloss.Left, views...)
}

func (q *CompositeQuestion) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyShiftDown, tea.KeyShiftUp:
			qCmd := q.questions[q.focusedQuestion].Update(msg)
			if qCmd == nil {
				return q.selectQuestion(msg.Type)
			}
		default:
			return q.questions[q.focusedQuestion].Update(msg)
		}
	case tea.MouseMsg:
		if msg.Action != tea.MouseActionRelease || msg.Button != tea.MouseButtonLeft {
			return nil
		}

		for i := range q.questions {
			qZone := zone.Get(q.getQuestionZoneMarkKey(i))
			if qZone.InBounds(msg) {
				if i == q.focusedQuestion {
					return nil
				}
				q.questions[q.focusedQuestion].Blur()
				q.questions[i].Focus()
				q.focusedQuestion = i
				return nil
			}
		}
	}

	return nil
}

func (q *CompositeQuestion) selectQuestion(keyPressed tea.KeyType) tea.Cmd {
	index := q.focusedQuestion
	q.questions[index].Blur()
	switch keyPressed {
	case tea.KeyShiftUp:
		if index == 0 {
			return nil
		}
		index -= 1
	case tea.KeyShiftDown:
		if index == len(q.questions)-1 {
			return nil
		}
		index += 1
	}
	q.focusedQuestion = index
	q.questions[index].Focus()
	return (func() tea.Msg {
		return ""
	})
}

func (q *CompositeQuestion) getQuestionZoneMarkKey(index int) string {
	return fmt.Sprintf("%s-%d", q.questionZonePrefix, index)
}
