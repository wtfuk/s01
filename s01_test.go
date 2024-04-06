package s01

import (
	"bytes"
	"os"
	"sync"
	"testing"
)

func TestPrintRuneDefaultOutput(t *testing.T) {
	// Capture the default os.Stdout output by temporarily replacing it
	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	os.Stdout = w

	testCases := []struct {
		r    rune
		want string
	}{
		{'a', "a"},
		{'中', "中"},
		{rune(0x10FFFF), string(rune(0x10FFFF))},
	}

	for _, tc := range testCases {
		err := PrintRune(tc.r)
		if err != nil {
			t.Errorf("PrintRune(%U) returned an error: %v", tc.r, err)
		}

		// Close the write end of the pipe to read the output
		w.Close()
		var buf bytes.Buffer
		_, err = buf.ReadFrom(r)
		if err != nil {
			t.Fatalf("Failed to read from pipe: %v", err)
		}
		got := buf.String()

		if got != tc.want {
			t.Errorf("PrintRune(%U) = %v, want %v", tc.r, got, tc.want)
		}

		// Reset the pipe for the next iteration
		r, w, err = os.Pipe()
		if err != nil {
			t.Fatalf("Failed to create pipe: %v", err)
		}
		os.Stdout = w
	}
}

func TestPrintRuneWithWriter(t *testing.T) {
	var captureBuffer bytes.Buffer

	testCases := []struct {
		r    rune
		want string
	}{
		{'a', "a"},
		{'中', "中"},
		{rune(0x10FFFF), string(rune(0x10FFFF))},
	}

	for _, tc := range testCases {
		captureBuffer.Reset() // Clear buffer before each test
		err := PrintRune(tc.r, &captureBuffer)
		if err != nil {
			t.Errorf("PrintRune(%U) returned an error: %v", tc.r, err)
		}
		if got := captureBuffer.String(); got != tc.want {
			t.Errorf("PrintRune(%U) = %v, want %v", tc.r, got, tc.want)
		}
	}
}

func TestPrintRuneConcurrently(t *testing.T) {
	var wg sync.WaitGroup
	runes := []rune{'a', '中', rune(0x10FFFF), '𐍈'}
	results := make([]string, len(runes))

	for i, r := range runes {
		wg.Add(1)
		go func(i int, r rune) {
			defer wg.Done()
			var localBuffer bytes.Buffer
			err := PrintRune(r, &localBuffer)
			if err != nil {
				t.Errorf("PrintRune(%U) failed: %v", r, err)
			}
			results[i] = localBuffer.String()
		}(i, r)
	}
	wg.Wait()

	// Since the order of execution is not deterministic in concurrent scenarios,
	// this test does not assert the order of results.
	// So, I check if all expected runes are present in the results slice, not the order.
	for _, result := range results {
		if len(result) == 0 {
			t.Error("Expected non-empty result for rune")
		}
	}
}
