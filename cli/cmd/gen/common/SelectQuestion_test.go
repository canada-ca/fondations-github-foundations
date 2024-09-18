package common

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
)

func getSelectQuestion() *SelectQuestion {
	question := NewSelectQuestion("Fruits?", []string{"apple", "banana", "pear"})
	return question
}

func TestNewSelectQuestion(t *testing.T) {
	question := getSelectQuestion()

	// Assert that the question has been created
	assert.NotNil(t, question)
}

func TestGetAnswer(t *testing.T) {
	question := getSelectQuestion()

	// Should be the first value by default
	val := question.GetAnswer()
	assert.Equal(t, "apple", val)

	question.model.Select(2)
	val = question.GetAnswer()
	assert.Equal(t, "pear", val)
}

func TestReset(t *testing.T) {
	question := getSelectQuestion()

	question.model.Select(2)
	val := question.GetAnswer()
	assert.Equal(t, "pear", val)
	question.Reset()

	val = question.GetAnswer()
	assert.Equal(t, "apple", val)
}

func TestSelectQuestion_Update(t *testing.T) {
	question := getSelectQuestion()

	// Create a tea message for selecting the second option
	messageDown := tea.KeyMsg {
		Type: tea.KeyDown,
	}
	messageEnter := tea.KeyMsg {
		Type: tea.KeyEnter,
	}

	// Select the second option
	question.Update(messageDown)
	question.Update(messageEnter)

	// Assert that the answer is correct
	assert.Equal(t, "banana", question.GetAnswer())

	// Select the third option
	question.Update(messageDown)
	question.Update(messageEnter)

	// Assert that the answer is correct
	assert.Equal(t, "pear", question.GetAnswer())
}

func TestSelectQuestion_SetDimensions(t *testing.T) {
	question := getSelectQuestion()

	// Set the dimensions
	question.SetDimensions(123, 456)

	// Assert that the dimensions have been set
	// Width should be the same
	// Height should be the same - 1
	assert.Equal(t, 123, question.model.Width())
	assert.Equal(t, 456 - 1, question.model.Height())
}

// Focus and Blur are not implemented for SelectQuestion
