package common

import (
	"strings"
	testing "testing"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
)


func getListQuestion() *ListQuestion {
	question := NewListQuestion("Animals?")
	// Use a BubbleTea List Model to set the values
	items := []list.Item{
		item{
			strValue: "Dog",
			value:    "Dog",
		},
		item{
			strValue: "Cat",
			value:    "Cat",
		},
		item{
			strValue: "Fish",
			value:    "Fish",
		},
	}

	list := list.New(items, itemDelegate{}, 0, 0)
	question.values = list
	return question
}

func TestNewListQuestion(t *testing.T) {

	question := NewListQuestion("Animals?")

	// Assert that the question has been created
	assert.NotNil(t, question)
	assert.Equal(t, "Animals?", question.prompt)
}

func TestListQuestion_GetAnswer(t *testing.T) {

	question := getListQuestion()

	// Assert that the view is correct
    answer := question.GetAnswer()
	parsedCmd := strings.ReplaceAll(answer, "\n", "")
	parsedCmd = strings.ReplaceAll(parsedCmd, "- ", " ")
	assert.Equal(t, "Dog Cat Fish", strings.Trim(parsedCmd, " "))
}

func TestListQuestion_Reset(t *testing.T) {

	question := getListQuestion()

	// Assert that the view is correct
    answer := question.GetAnswer()
	parsedCmd := strings.ReplaceAll(answer, "\n", "")
	parsedCmd = strings.ReplaceAll(parsedCmd, "- ", " ")
	assert.Equal(t, "Dog Cat Fish", strings.Trim(parsedCmd, " "))

	// Reset the question
	question.Reset()

	// Assert that the view is correct
	answer = question.GetAnswer()
	assert.Equal(t, "[]\n", answer)
}

func TestListQuestion_UpdateDelete(t *testing.T) {
	question := getListQuestion()

	// The state must be set to 'editing'
	question.state = editing

	// Test updating with a KeyMsg
	msg := tea.KeyMsg{
		Type:  tea.KeyDelete,
	}
	question.Update(msg)

	// Assert that the model has been updated with the input value
    answer := question.GetAnswer()
	parsedCmd := strings.ReplaceAll(answer, "\n", "")
	parsedCmd = strings.ReplaceAll(parsedCmd, "- ", " ")
	assert.Equal(t, "Cat Fish", strings.Trim(parsedCmd, " "))
}

func TestListQuestion_UpdateAdd(t *testing.T) {
	question := getListQuestion()
	// The state must be set to 'adding'
	question.state = adding
	question.input.SetValue("Bear")

	// Test updating with a KeyMsg
	msg := tea.KeyMsg{
		Type:  tea.KeyEnter,
	}
	question.Update(msg)

	// Assert that the model has been updated with the input value
    answer := question.GetAnswer()
	parsedCmd := strings.ReplaceAll(answer, "\n", "")
	parsedCmd = strings.ReplaceAll(parsedCmd, "- ", " ")
	assert.Equal(t, "Dog Cat Fish Bear", strings.Trim(parsedCmd, " "))
}

// Test Selecting, using the Down KeyMsg
func TestListQuestion_UpdateSelectDown(t *testing.T) {
	question := getListQuestion()

	// The state must be set to 'editing'
	question.state = editing

	// Test updating with a 'Down' KeyMsg
	msg := tea.KeyMsg{
		Type:  tea.KeyDown,
	}
	question.Update(msg)

	// Assert that the model has been updated with the input value
	answer := question.input.Value()
	assert.Equal(t, "Cat", answer)
}

// Test Selecting, using the Up KeyMsg
func TestListQuestion_UpdateSelectUp(t *testing.T) {
	question := getListQuestion()

	// The state must be set to 'editing'
	question.state = editing

	// Move with a 'Down' KeyMsg
	msg := tea.KeyMsg{
		Type:  tea.KeyDown,
	}
	question.Update(msg)

	// Assert that the model has been updated with the input value
	answer := question.input.Value()
	assert.Equal(t, "Cat", answer)

	// Test updating with an 'Up' KeyMsg
	msg = tea.KeyMsg{
		Type:  tea.KeyUp,
	}
	question.Update(msg)

	// Assert that the model has been updated with the input value
	answer = question.input.Value()
	assert.Equal(t, "Dog", answer)
}

// Test Selecting, using the Tab KeyMsg
func TestListQuestion_UpdateSelectTab(t *testing.T) {
	question := getListQuestion()

	question.state = editing
	// Move to the last item
	msg := tea.KeyMsg{
		Type:  tea.KeyDown,
	}
	question.Update(msg)
	question.Update(msg)

	// Assert that the model has been updated with the input value
	answer := question.input.Value()
	assert.Equal(t, "Fish", answer)

	msg = tea.KeyMsg{
		Type:  tea.KeyTab,
	}
	question.Update(msg)

	// Assert that the model has been updated with the input value
	answer = question.input.Value()
	assert.Equal(t, "", answer)
}

// TODO - Test that nothing happens on keypresses, when the state is not 'editing' or 'adding'
func TestListQuestion_SetDimensions(t *testing.T) {
	question := getListQuestion()

	question.SetDimensions(20, 12345)

	// Assert that the width and height have been set correctly
	// Min width is 45, Height is set to height - 5
	assert.Equal(t, 45, question.input.Width)
	assert.Equal(t, 12345 - 5, question.values.Height())

	// Set a width larger than 45
	question.SetDimensions(100, 12345)
	assert.Equal(t, 100/2, question.input.Width)
	assert.Equal(t, 12345 - 5, question.values.Height())
}

func TestListQuestion_Focus(t *testing.T) {
	question := getListQuestion()

	question.state = editing
	question.Focus()

	// Assert that the question is now in focus
	assert.True(t, question.input.Focused())
	// Assert that the state is now 'adding'
	assert.Equal(t, adding, question.state)
}

func TestListQuestion_Blur(t *testing.T) {
	question := getListQuestion()

	question.state = editing
	question.Focus()

	// Assert that the question is now in focus
	assert.True(t, question.input.Focused())
	// Assert that the state is now 'adding'
	assert.Equal(t, adding, question.state)

	question.Blur()

	// Assert that the question is now in focus
	assert.False(t, question.input.Focused())
	// Assert that the state is now 'unfocused'
	assert.Equal(t, unfocused, question.state)
}

func TestListQuestion_selectItem(t *testing.T) {
	question := getListQuestion()

	messageDown := tea.KeyMsg {
		Type: tea.KeyDown,
	}

	messageUp := tea.KeyMsg {
		Type: tea.KeyUp,
	}

	messageTab := tea.KeyMsg {
		Type: tea.KeyTab,
	}

	// Test selecting the first item
	question.selectItem(messageDown.Type)

	// Assert that the model has been updated with the input value
	answer := question.input.Value()
	assert.Equal(t, "Cat", answer)

	// Test selecting the second item
	question.selectItem(messageUp.Type)

	// Assert that the model has been updated with the input value
	answer = question.input.Value()
	assert.Equal(t, "Dog", answer)

	// Test selecting the third item
	question.selectItem(messageDown.Type)
	question.selectItem(messageDown.Type)

	// Assert that the model has been updated with the input value
	answer = question.input.Value()
	assert.Equal(t, "Fish", answer)

	// Test selecting the first item using the Tab key
	question.selectItem(messageTab.Type)
	answer = question.input.Value()
	assert.Equal(t, "Dog", answer)
}

func TestListQuestion_switchState(t *testing.T) {
	question := getListQuestion()

	// The state must be set to 'editing'
	question.state = editing

	message := tea.KeyMsg {
		Type:  tea.KeyTab,
	}

	// Test switching to 'adding'
	question.switchState(message)

	// Assert that the state has been switched
	assert.Equal(t, adding, question.state)

	// Test switching to 'editing'
	question.switchState(message)

	// Assert that the state has been switched
	assert.Equal(t, editing, question.state)
}
