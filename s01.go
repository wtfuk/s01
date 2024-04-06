package s01

import (
	"fmt"
	"io"
	"os"
	"sync"
	"unicode/utf8"
)

var runeBufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, utf8.UTFMax)
	},
}

// PrintRune prints a single rune (Unicode code point) to the provided writer.
// If no writer is provided, it defaults to os.Stdout.
// It returns an error if the rune cannot be encoded in UTF-8 or if writing fails.
func PrintRune(r rune, writers ...io.Writer) error {
	var w io.Writer = os.Stdout
	if len(writers) > 0 && writers[0] != nil {
		w = writers[0]
	}

	if !utf8.ValidRune(r) {
		return fmt.Errorf("invalid UTF-8 rune: %U", r)
	}

	buf := runeBufferPool.Get().([]byte)
	defer runeBufferPool.Put(buf)

	n := utf8.EncodeRune(buf[:utf8.UTFMax], r)
	if _, err := w.Write(buf[:n]); err != nil {
		return fmt.Errorf("failed to write rune: %w", err)
	}

	return nil
}
