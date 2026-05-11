package banner

import (
	"fmt"
	"testing"
)

func TestSpaceMapsToEightBlankLines(t *testing.T) {
	charMap, err := LoadBanner("../banners/standard.txt")
	if err != nil {
		t.Errorf("LoadBanner returned unexpected error: %v", err)
	}

	lines, ok := charMap[' ']
	if !ok {
		t.Error("charMap has no entry for ' ' (space, ASCII 32)")
	}

	expected := 8
	if len(lines) != expected {
		t.Errorf("Space expected %d lines, got %d", expected, len(lines))
	} else {
		fmt.Println("Pass")
	}
}

func TestAMapsCorrectly(t *testing.T) {
	charMap, err := LoadBanner("../banners/standard.txt")
	if err != nil {
		t.Errorf("LoadBanner returned unexpected error: %v", err)
	}

	lines, ok := charMap['A']
	if !ok {
		t.Error("charMap has no entry for 'A'")
	}

	expected := 8
	if len(lines) != expected {
		t.Errorf("'A' expected %d lines, got %d", expected, len(lines))
	} else if len(lines[0]) == 0 {
		t.Error("'A' line[0] is empty — wrong character mapped or index off by one")
	} else {
		fmt.Println("Pass")
	}
}

func TestExclamationMarkMapsCorrectly(t *testing.T) {
	charMap, err := LoadBanner("../banners/standard.txt")
	if err != nil {
		t.Errorf("LoadBanner returned unexpected error: %v", err)
	}

	lines, ok := charMap['!']
	if !ok {
		t.Error("charMap has no entry for '!'")
	}

	expected := 8
	if len(lines) != expected {
		t.Errorf("'!' expected %d lines, got %d", expected, len(lines))
	} else {
		spaceLines := charMap[' ']
		if len(lines) > 0 && len(spaceLines) > 0 && lines[0] == spaceLines[0] {
			t.Error("'!' line[0] is identical to ' ' line[0] — index mapping is broken")
		} else {
			fmt.Println("Pass")
		}
	}
}

func TestNonExistentFileReturnsError(t *testing.T) {
	charMap, err := LoadBanner("../banners/does_not_exist.txt")

	if err == nil {
		t.Error("Expected an error for a nonexistent file, got nil")
	}

	if charMap != nil {
		t.Errorf("Expected nil map on error, got map with %d entries", len(charMap))
	} else {
		fmt.Println("Pass")
	}
}

func TestAllCharactersHaveEightLines(t *testing.T) {
	charMap, err := LoadBanner("../banners/standard.txt")
	if err != nil {
		t.Errorf("LoadBanner returned unexpected error: %v", err)
	}

	for r, lines := range charMap {
		expected := 8
		if len(lines) != expected {
			t.Errorf("Rune %q (ASCII %d): expected %d lines, got %d",
				string(r), r, expected, len(lines))
		}
	}
	fmt.Println("Pass")
}