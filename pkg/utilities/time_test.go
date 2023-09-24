package utilities

import (
	"testing"
	"time"
)

// Generate Test function for FormatDefaultTime
// Path: pkg/utilities/time_test.go
func TestFormatDefaultTime(t *testing.T) {
	// Test case 1
	// Input:
	now := time.UnixMilli(1210902972121)
	// Expected output:
	expect := "2008-05-16T08:56:12+07:00"
	// Actual output:
	actual := FormatDefaultTime(now)
	if actual != expect {
		t.Errorf("TestFormatDefaultTime() failed, actual: %s, expect: %s", actual, expect)
	}

	// Test case 3
	// Input:
	now3 := time.UnixMilli(1888282922121)
	// Expected output:
	expect = "2029-11-02T10:02:02+07:00"
	// Actual output:
	actual = FormatDefaultTime(now3)
	if actual != expect {
		t.Errorf("TestFormatDefaultTime() failed, actual: %s, expect: %s", actual, expect)
	}

}
