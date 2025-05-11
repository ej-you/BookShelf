package excel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_filepath = "./export_test.xlsx"
)

type Human struct {
	Name   string
	gender string
	Age    int
	About  string
}

type Humans []Human

type OneFieldStruct struct {
	Greeting string
}

func TestFillAndSave_HumanSlice(t *testing.T) {
	t.Log("Export humans to excel file")

	humans := []Human{
		{
			Name:   "Ivan",
			gender: "M",
			Age:    18,
			About:  "I am clever",
		},
		{
			Name:   "Vasya",
			gender: "M",
			Age:    25,
			About:  "I am stupid",
		},
	}

	err := FillAndSave(_filepath, humans)
	require.NoError(t, err)

	t.Logf("Humans was saved to %s successfully!", _filepath)
}

func TestFillAndSave_Humans(t *testing.T) {
	t.Log("Export humans to excel file")

	humans := Humans{
		{
			Name:   "Ivan",
			gender: "M",
			Age:    18,
			About:  "I am clever",
		},
		{
			Name:   "Vasya",
			gender: "M",
			Age:    25,
			About:  "I am stupid",
		},
	}

	err := FillAndSave(_filepath, humans)
	require.NoError(t, err)

	t.Logf("Humans was saved to %s successfully!", _filepath)
}

func TestFillAndSave_OneFieldStruct(t *testing.T) {
	t.Log("Export OneFieldStruct to excel file")

	oneField := []OneFieldStruct{
		{Greeting: "Hello"},
		{Greeting: "Hi"},
	}

	err := FillAndSave(_filepath, oneField)
	require.NoError(t, err)

	t.Logf("OneFieldStruct was saved to %s successfully!", _filepath)
}

func TestFillAndSave_AllErrors(t *testing.T) {
	var err error

	humans := []*Human{
		{
			Name:   "Ivan",
			gender: "M",
			Age:    18,
			About:  "I am clever",
		},
	}

	t.Log("Empty data slice")
	err = FillAndSave(_filepath, []Human{})
	require.Error(t, err)
	t.Log("Gotten error (expected):", err)

	t.Log("Not xlsx file")
	err = FillAndSave("./sample.txt", humans)
	require.Error(t, err)
	t.Log("Gotten error (expected):", err)

	t.Log("Slice of pointers to structs")
	err = FillAndSave(_filepath, humans)
	require.Error(t, err)
	t.Log("Gotten error (expected):", err)

	t.Log("Slice of strings")
	err = FillAndSave(_filepath, []string{"first", "second"})
	require.Error(t, err)
	t.Log("Gotten error (expected):", err)
}
