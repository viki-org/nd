package nd

import (
	// "bytes"
	"regexp"
	"testing"
)

func TestGuidv4ReturnsRandomGuids(t *testing.T) {
	assertGuidv4IsRandom(t)
}

func TestGuidv4StringLooksOk(t *testing.T) {
	guid := Guidv4String()
	matched, _ := regexp.Match("^[0-9]{19}$", []byte(guid))
	// matched, _ := regexp.Match("[0-9]{19}[\\da-f]{8}\\-[\\da-f]{4}\\-[\\da-f]{4}\\-[\\da-f]{4}\\-[\\da-f]{8}", []byte(guid))
	if matched == false {
		t.Error("Invalid format for guid string")
	}
}

func TestCanForceAGuidv4(t *testing.T) {
	expected := "aaaaaaaa-bbbb-cccc-dddd-bbbbbbbb"
	ForceGuid(expected)

	guidBytes := Guidv4()
	if string(guidBytes) != expected {
		// if bytes.Compare(guidBytes, expectedBytes) != 0 {
		t.Errorf("Guid should be %q, got %q", []byte(expected), guidBytes)
	}

	guidString := Guidv4String()
	if guidString != "aaaaaaaa-bbbb-cccc-dddd-bbbbbbbb" {
		t.Errorf("Guid string should be %q, got %q", expected, guidString)
	}
}

func TestCanResetGuidv4(t *testing.T) {
	ForceGuid("aaaaaaaa-bbbb-cccc-dddd-bbbbbbbb")
	ResetGuidv4()
	assertGuidv4IsRandom(t)
}

func assertGuidv4IsRandom(t *testing.T) {
	seen := make(map[string]bool, 500)
	for i := 0; i < 500; i++ {
		seen[string(Guidv4())] = true
	}
	if len(seen) != 500 {
		t.Error("Should have seen 500 unique guids")
	}
}
