package match_finder

import (
	"testing"
)

func TestCanMatchByKey(t *testing.T) {
	slang := make(map[string]string)
	slang["arms"] = "strong"

	matches := CheckForMatches([]string{"arms", "blah blah"}, slang)

	if len(matches) != 1 {
		t.Error("Not of desired length")
	}

	if matches["arms"] != "strong" {
		t.Error("Invalid match")
	}
}

func TestCanMatchByValue(t *testing.T) {
	slang := make(map[string]string)
	slang["arms"] = "strong"

	matches := CheckForMatches([]string{"strong", "blah blah"}, slang)

	if len(matches) != 1 {
		t.Error("Not of desired length")
	}

	if matches["strong"] != "arms" {
		t.Error("Invalid match")
	}
}