package common

import (
	"regexp"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
)

func TestNewTextQuestion(t *testing.T) {
	question := NewTextQuestion("Enter your name", "")

	// Assert that the question has been created
	assert.NotNil(t, question)
}

func TestTextQuestion_GetAnswer(t *testing.T) {
	question := NewTextQuestion("Enter your name", "My Name is GHF")

	// Assert that the answer is correct
	assert.Equal(t, "My Name is GHF", question.GetAnswer())
}

// Test for updating the text question when the input is not in focus
func TestTextQuestion_UpdateNoFocus(t *testing.T) {

	question := NewTextQuestion("Enter your name", "")
	message := tea.KeyMsg {
		Type: tea.KeyRunes,
		Runes: []rune("ghf"),
	}

	// This should fail, since the input is not in focus
	cmd := question.Update(message)
	// Assert that the model has not been updated with the input value
	assert.Equal(t, "", question.GetAnswer())
	// Assert that the command returned is nil
	assert.Nil(t, cmd)
}

// Test for updating the text question when the input is not in focus
func TestTextQuestion_Update(t *testing.T) {

	question := NewTextQuestion("Enter your name", "")
	// Set the question in focus
	question.Focus()

	message := tea.KeyMsg {
		Type: tea.KeyRunes,
		Runes: []rune("ghf"),
	}

	// This should succeed, since the input is in focus
	cmd := question.Update(message)
	// Assert that the model has been updated with the input value
	assert.Equal(t, "ghf", question.GetAnswer())
	// Assert that the command returned a TeaCmd
	assert.NotNil(t, cmd)
}

// Test the reset function of the text question
func TestTextQuestion_Reset(t *testing.T) {
	question := NewTextQuestion("Enter your name", "My Name is GHF")
	// Set the question in focus
	question.Focus()

	message := tea.KeyMsg {
		Type: tea.KeyRunes,
		Runes: []rune(", and I'm happy to meet you."),
	}
	question.Update(message)
	assert.Equal(t, "My Name is GHF, and I'm happy to meet you.", question.GetAnswer())

	// Reset the question)
	question.Reset()

	// Assert that the answer has been reset
	assert.Equal(t, "My Name is GHF", question.GetAnswer())
}

func TestTextQuestion_View(t *testing.T) {
	question := NewTextQuestion("Enter your name", "My Name is GHF")

	cmd := question.View()
	// Assert that the view is correct
	parsedCmd := strings.ReplaceAll(cmd, "\n", "")
	regex, _ := regexp.Compile(`\s+Enter your name:\s+My Name is GHF\s+━+`)
	assert.True(t, regex.MatchString(parsedCmd))

}

func TestTextQuestion_SetDimensions(t *testing.T) {
	question := NewTextQuestion("Enter your name", "")

	question.SetDimensions(20, 12345)

	// Assert that the width and height have been set correctly
	// Widths less than 45 are set to 45
	assert.Equal(t, 45, question.inputModel.Width)

	// Widths larger than 45 are set to w/2
	question.SetDimensions(100, 12345)
	assert.Equal(t, 100/2, question.inputModel.Width)
	question.SetDimensions(999, 12345)
	assert.Equal(t, 999/2, question.inputModel.Width)
	question.SetDimensions(5554, 12345)
	assert.Equal(t, 5554/2, question.inputModel.Width)
}

func TestTextQuestion_Focus(t *testing.T) {
	question := NewTextQuestion("Enter your name", "")

	// Assert that the question is not in focus
	assert.False(t, question.inputModel.Focused())

	question.Focus()

	// Assert that the question is now in focus
	assert.True(t, question.inputModel.Focused())
}

func TestTextQuestion_Blur(t *testing.T) {
	question := NewTextQuestion("Enter your name", "")

	question.Focus()

	// Assert that the question is in focus
	assert.True(t, question.inputModel.Focused())

	question.Blur()

	// Assert that the question is no longer in focus
	assert.False(t, question.inputModel.Focused())
}
