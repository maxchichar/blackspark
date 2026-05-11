package parser

import (
	"reflect"
	"testing"
)

func TestSingleWordNoNewlines(t *testing.T) {
	input := "hello"
	expected := []string{"hello"}
	result := Parser(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result: %q Expected: %q", result, expected)
	}
}

func TestEmpty(t *testing.T){
	input := ""
	expected := []string{""}
	result := Parser(input)

	if !reflect.DeepEqual(result, expected){
		t.Errorf("Result: %q Expected: %q", result, expected)
	}
}

func TestTwoLinesSeperatedByOneNewline(t *testing.T) {
	input := "Hello\nThere"
	expected := []string{"Hello", "There"}
	result := Parser(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result: %q Expected: %q", result, expected)
	}
}

func TestDoubleNewlineProducesEmptyStringInMiddle(t *testing.T) {
	input := "Hello\n\nThere"
	expected := []string{"Hello", "", "There"}
	result := Parser(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result: %q Expected: %q", result, expected)
	}
}

func TestJustNewLine(t *testing.T) {
	input := "\n"
	expected := []string{"", ""}
	result := Parser(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result: %q Expected: %q", result, expected)
	}
}

func TestEmptyString(t *testing.T) {
	input := ""
	expected := []string{""}
	result := Parser(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result: %q Expected %q", result, expected)
	}
}

func TestTrailingNewLine(t *testing.T) {
	input := "Hello\n"
	expected := []string{"Hello", ""}
	result := Parser(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result: %q Expected: %q", result, expected)
	}
}
