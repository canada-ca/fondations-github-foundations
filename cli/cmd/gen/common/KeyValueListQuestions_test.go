package common

import (
	"fmt"
	testing "testing"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
)

const expectedList = "animal: Dog\nanimal1: Cat\nanimal2: Fish\nfruit: Apple\nfruit1: Banana\nfruit2: Pear\n"

func getKeyValueListQuestion() *KeyValueListQuestion {
	question := NewKeyValueListQuestion("Fruits or Animals?")
	fruits := []string{"Apple", "Banana", "Pear"}

	fruitItems := []list.Item{
		item{
			strValue: fmt.Sprintf(`fruit = %s`, fruits[0]),
			value:    keyValuePair {
				key: "fruit",
				value: fruits[0],
			},
		},
		item{
			strValue: fmt.Sprintf(`fruit1 = %s`, fruits[1]),
			value:    keyValuePair {
				key: "fruit1",
				value: fruits[1],
			},
		},
		item{
			strValue: fmt.Sprintf(`fruit2 = %s`, fruits[2]),
			value:    keyValuePair {
				key: "fruit2",
				value: fruits[2],
			},
		},
	}

	fruitList := list.New(fruitItems, itemDelegate{}, 0, 0)

	animals := []string{"Dog", "Cat", "Fish"}

	animalItems := []list.Item{
		item{
			strValue: fmt.Sprintf(`animal = %s`, animals[0]),
			value:    keyValuePair {
				key: "animal",
				value: animals[0],
			},
		},
		item{
			strValue: fmt.Sprintf(`animal1 = %s`, animals[1]),
			value:    keyValuePair {
				key: "animal1",
				value: animals[1],
			},
		},
		item{
			strValue: fmt.Sprintf(`animal2 = %s`, animals[2]),
			value:    keyValuePair {
				key: "animal2",
				value: animals[2],
			},
		},
	}

	animalList := list.New(animalItems, itemDelegate{}, 0, 0)

	listModel := list.New(append(fruitList.Items(), animalList.Items()...), itemDelegate{}, 0, 0)

	question.listModel.SetItems(listModel.Items())
	question.keyValueMap = map[string]string{
		"fruit": fruits[0],
		"fruit1": fruits[1],
		"fruit2": fruits[2],
		"animal": animals[0],
		"animal1": animals[1],
		"animal2": animals[2],
	}
	return question
}

func TestNewKeyValueList(t *testing.T) {

	question := getKeyValueListQuestion()

	// Assert that the question has been created
	assert.NotNil(t, question)
	assert.Equal(t, "Fruits or Animals?", question.prompt)
}

func TestKeyValueListQuestion_Reset(t *testing.T) {
	question := getKeyValueListQuestion()

	// Assert that the structure is correct
	answer := question.prompt
	assert.Equal(t, "Fruits or Animals?", answer)
	assert.NotEmpty(t, question.listModel.Items())
	answer = question.GetAnswer()
	assert.Equal(t, expectedList, answer)


	// Reset the question
	question.Reset()

	// Assert that the structure is correct
	answer = question.prompt
	assert.Equal(t, "Fruits or Animals?", answer)
	answer = question.GetAnswer()
	assert.Equal(t, "{}\n", answer)
}

func TestKeyValueListQuestion_GetAnswer(t *testing.T) {
	question := getKeyValueListQuestion()

	// Assert that the view is correct
	answer := question.GetAnswer()
	assert.Equal(t, expectedList, answer)
}

func TestKeyValueListQuestion_UpdateDelete(t *testing.T) {
	question := getKeyValueListQuestion()

	// Add some values for the List Model
	testItems := []list.Item{
		item{
			strValue: "String = Value",
		},
		item{
			strValue: "String = Value 2",
		},
	}
	question.listModel = list.New(testItems, itemDelegate{}, 0, 0)

	// the state must be 'editing'
	question.state = editing

	messageDelete := tea.KeyMsg {
		Type:  tea.KeyDelete,
	}
	// Update the question
	question.Update(messageDelete)

	// Assert that the view is correct
	listModel := question.listModel
	answer := listModel.Items()[0].FilterValue()
	assert.Equal(t, "String = Value 2", answer)

}

// Tests replacing a value in the list. The value is deleted,
// and the replacement is added to the end of the list.
func TestKeyValueListQuestion_UpdateReplace(t *testing.T) {
	question := getKeyValueListQuestion()

	// Add some values for the List Model
	testItems := []list.Item{
		item{
			strValue: "String = Value",
		},
		item{
			strValue: "String = Value 2",
		},
	}
	question.listModel = list.New(testItems, itemDelegate{}, 0, 0)


	question.keyInputModel.SetValue("String")
	question.valueInputModel.SetValue("Value 3")

	// the state must be 'editing'
	question.state = editing

	messageReplace := tea.KeyMsg {
		Type:  tea.KeyEnter,
	}
	// Update the question
	question.Update(messageReplace)

	// Assert that the view is correct
	listModel := question.listModel
	answer := listModel.Items()[1].FilterValue()
	assert.Equal(t, "String = Value 3", answer)
}

func TestKeyValueListQuestion_UpdateSelect(t *testing.T) {
	question := getKeyValueListQuestion()

	// the state must be 'editing'
	question.state = editing

	messageSelectUp := tea.KeyMsg {
		Type:  tea.KeyUp,
	}
	messageSelectDown := tea.KeyMsg {
		Type:  tea.KeyDown,
	}

	// Select the first Item
	question.listModel.Select(0)
	assert.Equal(t, "fruit = Apple", question.listModel.SelectedItem().FilterValue())

	// Select the second Item
	question.Update(messageSelectDown)
	assert.Equal(t, "fruit1 = Banana", question.listModel.SelectedItem().FilterValue())

	// Select the third Item
	question.Update(messageSelectDown)
	assert.Equal(t, "fruit2 = Pear", question.listModel.SelectedItem().FilterValue())

	// Select the first item
	question.Update(messageSelectUp)
	question.Update(messageSelectUp)
	assert.Equal(t, "fruit = Apple", question.listModel.SelectedItem().FilterValue())
}

func TestKeyValueListQuestion_UpdateFocus(t *testing.T) {
	question := getKeyValueListQuestion()

	// the state must be not be 'unfocused'
	question.state = adding

	messageShiftTab := tea.KeyMsg {
		Type: tea.KeyShiftTab,
	}
	// Give the keyInputModel an initial state
	question.keyInputModel.Focus()
	// Give the valueInputModel an initial state
	question.valueInputModel.Blur()

	// Update the question
	question.Update(messageShiftTab)
	assert.False(t, question.keyInputModel.Focused())
	assert.True(t, question.valueInputModel.Focused())

	// Switch back
	question.Update(messageShiftTab)
	assert.True(t, question.keyInputModel.Focused())
	assert.False(t, question.valueInputModel.Focused())
}

func TestKeyValueListQuestion_UpdateSwitchState(t *testing.T) {
	question := getKeyValueListQuestion()

	// the state must be 'editing'
	question.state = editing

	messageTab := tea.KeyMsg {
		Type:  tea.KeyTab,
	}
	// Update the question
	question.Update(messageTab)
	// Assert that the state is correct
	assert.Equal(t, adding, question.state)

	// Switch back
	question.Update(messageTab)
	// Assert that the state is correct
	assert.Equal(t, editing, question.state)
}

func TestKeyValueListQuestion_SetDimensions(t *testing.T) {
	question := getKeyValueListQuestion()

	// Set the dimensions. We assume that the dimensions will be set as follows:
	// question.keyInputModel.Width = widthValue/2 - 5
	// question.valueInputModel.Width = widthValue/2 - 5
	// question.listModel.Width = widthValue
	// question.listModel.Height = heightValue - 5
	question.SetDimensions(444, 666)

	// Assert that the dimensions are correct
	assert.Equal(t, 217, question.keyInputModel.Width)
	assert.Equal(t, 217, question.valueInputModel.Width)
	assert.Equal(t, 444, question.listModel.Width())
	assert.Equal(t, 661, question.listModel.Height())
}

func TestKeyValueListQuestion_Focus(t *testing.T) {
	question := getKeyValueListQuestion()

	// Setup conditions for the test
	question.keyInputModel.Blur()
	question.valueInputModel.Focus()
	question.state = editing

	// Focus the question
	question.Focus()

	// Assert that the question is focused
	// We expect that:
	// question.keyInputModel.Focused() = true
	// question.valueInputModel.Focused() = false
	// question.state = adding
	assert.True(t, question.keyInputModel.Focused())
	assert.False(t, question.valueInputModel.Focused())
	assert.Equal(t, adding, question.state)
}

func TestKeyValueListQuestion_Blur(t *testing.T) {
	question := getKeyValueListQuestion()

	// Setup conditions for the test
	question.keyInputModel.Focus()
	question.valueInputModel.Blur()
	question.state = adding

	// Blur the question
	question.Blur()

	// Assert that the question is blurred
	// We expect that:
	// question.keyInputModel.Focused() = false
	// question.valueInputModel.Focused() = false
	// question.state = unfocused
	assert.False(t, question.keyInputModel.Focused())
	assert.False(t, question.valueInputModel.Focused())
	assert.Equal(t, unfocused, question.state)
}

func TestCreateItem(t *testing.T) {
	question := getKeyValueListQuestion()

	// Create an item
	item := question.createItem("key", "value")

	// Assert that the item has been created
	assert.NotNil(t, item)
	assert.Equal(t, "key = value", item.FilterValue())
}

func TestGetKeyValuePair(t *testing.T) {
	question := getKeyValueListQuestion()

	// Get the keyValuePair
	keyValue := question.getKeyValuePair(3)

	// Assert that the keyValuePair is correct
	assert.Equal(t, "animal", keyValue.key)
	assert.Equal(t, "Dog", keyValue.value)

	// Get the keyValuePair
	keyValue = question.getKeyValuePair(5)

	// Assert that the keyValuePair is correct
	assert.Equal(t, "animal2", keyValue.key)
	assert.Equal(t, "Fish", keyValue.value)
}

func TestSwitchState(t *testing.T) {
	question := getKeyValueListQuestion()

	// Set an initial state
	question.state = editing

	messageKey := tea.KeyMsg {
		Type: tea.KeyEnter,
	}

	// Switch the state
	question.switchState(messageKey)

	// Assert that the state has been switched
	assert.Equal(t, adding, question.state)
	assert.Empty(t, question.keyInputModel.Value())
	assert.Empty(t, question.valueInputModel.Value())

	// Switch the state again
	question.switchState(messageKey)

	// Assert that the state has been switched
	assert.Equal(t, editing, question.state)
	assert.Equal(t, "fruit", question.keyInputModel.Value())
	assert.Equal(t, "Apple", question.valueInputModel.Value())
}

func TestPutEntry(t *testing.T) {
	question := getKeyValueListQuestion()

	// Add an entry
	question.putEntry("fruit3", "Orange")

	// Assert that the entry has been added
	assert.Equal(t, "Orange", question.keyValueMap["fruit3"])
	assert.Equal(t, "fruit3 = Orange", question.listModel.Items()[6].FilterValue())

	// Try adding an existing entry
	question.putEntry("animal1", "Mouse")
	assert.Equal(t, "Mouse", question.keyValueMap["animal1"])
	assert.Equal(t, "animal1 = Mouse", question.listModel.Items()[4].FilterValue())
}
