package common

import (
	"fmt"
	"strings"
	testing "testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
)

const expectedAnswer = "SelectQuestion: Red\nTextQuestion: Hi my name is Question\n"

func getCompositeQuestion() *CompositeQuestion{

	// Create some composite questions
	textQuestion := NewTextQuestion("What is your name?", "Hi my name is Question")
	selectQuestion := NewSelectQuestion("What is your favorite color?", []string{"Red", "Green", "Blue"})

	textQuestionEntry := CompositeQuestionEntry{
		Key: "TextQuestion",
		Question: textQuestion,
	}
	selectQuestionEntry := CompositeQuestionEntry{
		Key: "SelectQuestion",
		Question: selectQuestion,
	}

	return NewCompositeQuestion("CompositeQuestion", []CompositeQuestionEntry{textQuestionEntry, selectQuestionEntry})
}

func TestNewCompositeQuestion(t *testing.T) {
	question := getCompositeQuestion()

	assert.NotNil(t, question)
}

func TestCompositeQuestion_Reset(t *testing.T) {
	question := getCompositeQuestion()

	textQuestion := question.questions[0].(*TextQuestion)
	selectQuestion := question.questions[1].(*SelectQuestion)

	// Change the answers
	textQuestion.inputModel.SetValue("Wrong Name!")
	selectQuestion.model.Select(1)
	assert.Equal(t, "Wrong Name!", textQuestion.GetAnswer())
	assert.Equal(t, "Green", selectQuestion.GetAnswer())

	// Reset the answers
	question.Reset()
	assert.Equal(t, "Hi my name is Question", textQuestion.GetAnswer())
	assert.Equal(t, "Red", selectQuestion.GetAnswer())
}

func TestCompositeQuestion_GetAnswer(t *testing.T) {
	question := getCompositeQuestion()

	// Change the answers
	assert.Equal(t, expectedAnswer, question.GetAnswer())
}

func TestCompositeQuestion_SetDimensions(t *testing.T) {
	question := getCompositeQuestion()

	// By setting the dimensions, we expect:
	// The questions height will be the max of max(height/3, q.questionHeight).
	// The questions width will be the value passed in.
	// For TextQuestion, there is not height property, and the width should
	// be half of the value passed in, when larger than 90.
	// For SelectQuestion, the height should be (height/3 - 1),
	// and the width should be the value passed in.
	question.SetDimensions(120, 200)

	textQuestion := question.questions[0].(*TextQuestion)
	selectQuestion := question.questions[1].(*SelectQuestion)

	assert.Equal(t, 60, textQuestion.inputModel.Width)
	assert.Equal(t, 120, selectQuestion.model.Width())
	// There's no height property for TextQuestions
	assert.Equal(t, 65, selectQuestion.model.Height())
}

func TestCompositeQuestion_Focus(t *testing.T) {
	question := getCompositeQuestion()

	textQuestion := question.questions[0].(*TextQuestion)

	// SelectQuestion don't have a focused property

	// Focus on the second question
	question.focusedQuestion = 1
	question.Focus()
	assert.False(t, textQuestion.inputModel.Focused())

	// Focus on the first question
	question.focusedQuestion = 0
	question.Focus()
	assert.True(t, textQuestion.inputModel.Focused())
}

func TestCompositeQuestion_Blur(t *testing.T) {
	question := getCompositeQuestion()

	textQuestion := question.questions[0].(*TextQuestion)

	// SelectQuestion don't have a focused property

	// Focus on the first question
	question.focusedQuestion = 0
	question.Focus()
	assert.True(t, textQuestion.inputModel.Focused())

	// Blur the first question
	question.Blur()
	assert.False(t, textQuestion.inputModel.Focused())
}

func TestCompositeQuestion_UpdateValue(t *testing.T) {
	question := getCompositeQuestion()

	textQuestion := question.questions[0].(*TextQuestion)

	// Set the question in focus
	textQuestion.Focus()
	question.focusedQuestion = 0

	messageRunes := tea.KeyMsg {
		Type: tea.KeyRunes,
		Runes: []rune("mark"),
	}

	// This should succeed, since the input is in focus
	question.Update(messageRunes)
	// Assert that the model has been updated with the input value
	trimmed, _ := strings.CutSuffix(expectedAnswer, "\n")
	expectedString := fmt.Sprintf("%smark\n", trimmed)
	assert.Equal(t, expectedString, question.GetAnswer())
}

func TestCompositeQuestion_UpdateKeyDown(t *testing.T) {
	question := getCompositeQuestion()

	// Update the select question
	messageSelect := tea.KeyMsg {
		Type: tea.KeyDown,
	}

	// Choose the SelectQuestion
	question.focusedQuestion = 1
	// This should select the second item in the select question (Green)
	question.Update(messageSelect)
	replaceRed := strings.ReplaceAll(expectedAnswer, "Red", "Green")
	assert.Equal(t, replaceRed, question.GetAnswer())

}

// TODO - Write tests for the following "Update" triggers:
// - tea.KeyShiftUp
// - tea.KeyShiftDown
// - tea.MouseActionRelease && tea.MouseButtonLeft

func TestCompositeQuestion_SelectQuestion(t *testing.T) {
	question := getCompositeQuestion()

	// Choose the SelectQuestion
	question.focusedQuestion = 1
	question.questions[question.focusedQuestion].Focus()
	assert.Equal(t, "Red", question.questions[question.focusedQuestion].GetAnswer())

	messageShiftUp := tea.KeyMsg {
		Type: tea.KeyShiftUp,
	}

	// This should select the first question
	question.Update(messageShiftUp)
	assert.Equal(t, 0, question.focusedQuestion)
	assert.Equal(t, "Hi my name is Question", question.questions[question.focusedQuestion].GetAnswer())

	messageShiftDown := tea.KeyMsg {
		Type: tea.KeyShiftDown,
	}

	// This should select the second question
	question.Update(messageShiftDown)
	assert.Equal(t, 1, question.focusedQuestion)

}
